package vyosclient

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/kuun/slog"
	"github.com/pkg/errors"
)

var log = slog.GetLogger(&VyosError{})

type VyosError struct {
	msg string
}

func (e *VyosError) Error() string {
	return e.msg
}

type Command struct {
	Op    string   `json:"op"`
	Path  []string `json:"path,omitempty"`
	Value string   `json:"value,omitempty"`
	File  string   `json:"file,omitempty"`
	URL   string   `json:"url,omitempty"`
	Name  string   `json:"name,omitempty"`
}

func NewCommand(op string, path string, value string) *Command {
	cmd := &Command{
		Op:    op,
		Value: value,
	}
	if path != "" {
		cmd.Path = strings.Split(path, " ")
	}
	return cmd
}

type CmdResult struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
}

func NewVyosError(msg string) *VyosError {
	return &VyosError{
		msg: msg,
	}
}

type VyosClient interface {
	Configure(cmds []*Command) (err error)
	ShowConfig(path string) (data interface{}, err error)
	ReturnValue(path string) (value string, err error)
	ReturnValues(path string) (values []string, err error)
}

type vyosClient struct {
	key    string
	server string
	port   uint16
}

var client *vyosClient

func GetClient() VyosClient {
	return client
}

func InitVyosClient(server string, port uint16, key string) {
	client = &vyosClient{
		key:    key,
		server: server,
		port:   port,
	}
}

func (client *vyosClient) Configure(cmds []*Command) (err error) {
	var httpRes *http.Response
	var result *CmdResult
	url := client.buildUrl("configure")

	httpRes, err = client.post(url, cmds)
	if err != nil {
		goto ERROR
	}
	defer httpRes.Body.Close()
	result, err = client.parseJsonResponse(httpRes)
	if err != nil {
		goto ERROR
	}
	log.Debugf("configure result: %#v", result)
	if !result.Success {
		err = errors.New(result.Error)
		goto ERROR
	}
	return nil
ERROR:
	return err
}

func (client *vyosClient) ShowConfig(path string) (data interface{}, err error) {
	var httpRes *http.Response
	var result *CmdResult
	url := client.buildUrl("retrieve")

	httpRes, err = client.post(url, &Command{
		Op:   "showConfig",
		Path: strings.Split(path, " "),
	})
	if err != nil {
		goto ERROR
	}
	defer httpRes.Body.Close()
	result, err = client.parseJsonResponse(httpRes)
	if err != nil {
		goto ERROR
	}
	if !result.Success {
		err = errors.Wrap(NewVyosError(result.Error), "")
		goto ERROR
	}
	return result.Data, nil
ERROR:
	return nil, err
}

func (client *vyosClient) ReturnValue(path string) (value string, err error) {
	var httpRes *http.Response
	var result *CmdResult
	url := client.buildUrl("retrieve")
	var ok bool

	httpRes, err = client.post(url, &Command{
		Op:   "returnValue",
		Path: strings.Split(path, " "),
	})
	if err != nil {
		goto ERROR
	}
	defer httpRes.Body.Close()
	result, err = client.parseJsonResponse(httpRes)
	if err != nil {
		goto ERROR
	}
	if !result.Success {
		err = errors.Wrap(NewVyosError(result.Error), "")
		goto ERROR
	}
	if result.Data == nil {
		return "", nil
	} else {
		value, ok = result.Data.(string)
		if !ok {
			log.Errorf("server return invalid data, result: %#v", result)
		}
		return value, nil
	}
ERROR:
	return "", err
}

func (client *vyosClient) ReturnValues(path string) (values []string, err error) {
	var httpRes *http.Response
	var result *CmdResult
	url := client.buildUrl("retrieve")
	var ok bool

	httpRes, err = client.post(url, &Command{
		Op:   "returnValues",
		Path: strings.Split(path, " "),
	})
	if err != nil {
		goto ERROR
	}
	defer httpRes.Body.Close()
	result, err = client.parseJsonResponse(httpRes)
	if err != nil {
		goto ERROR
	}
	if !result.Success {
		err = errors.Wrap(NewVyosError(result.Error), "")
		goto ERROR
	}

	values, ok = result.Data.([]string)
	if !ok {
		log.Errorf("server return invalid data, result: %#v", result)
	}
	return values, nil
ERROR:
	return nil, err
}

func (client *vyosClient) post(url string, cmds interface{}) (res *http.Response, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	data, contentType, err := client.buildMultiparBuf(cmds)
	if err != nil {
		goto ERROR
	}
	res, err = httpClient.Post(url, contentType, data)
	if err != nil {
		err = errors.Wrap(err, "")
		goto ERROR
	}
	return
ERROR:
	return nil, err
}

func (client *vyosClient) parseJsonResponse(httpRes *http.Response) (result *CmdResult, err error) {
	var res CmdResult
	decoder := json.NewDecoder(httpRes.Body)
	if err = decoder.Decode(&res); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return &res, err
}

func (client *vyosClient) buildUrl(uri string) string {
	return fmt.Sprintf("https://%s:%d/%s", client.server, client.port, uri)
}

func (client *vyosClient) buildMultiparBuf(cmds interface{}) (*bytes.Buffer, string, error) {
	var buf bytes.Buffer
	var data []byte

	writer := multipart.NewWriter(&buf)
	writer.WriteField("key", client.key)
	dataWriter, err := writer.CreateFormField("data")
	if err != nil {
		err = errors.Wrap(err, "")
		goto ERROR
	}
	data, err = json.Marshal(cmds)
	if err != nil {
		err = errors.Wrap(err, "")
		goto ERROR
	}
	_, err = dataWriter.Write(data)
	if err != nil {
		err = errors.Wrap(err, "")
		goto ERROR
	}
	writer.Close()
	return &buf, writer.FormDataContentType(), err
ERROR:
	return nil, "", err
}
