package models

type OpType string

const (
	OpShow       OpType = "show"
	OpShowConfig OpType = "showConfig"
	OpSet        OpType = "set"
	OpDelete     OpType = "delete"
	OpAdd        OpType = "add"
	OpGenerate   OpType = "generate"
	OpComment    OpType = "comment"
)

