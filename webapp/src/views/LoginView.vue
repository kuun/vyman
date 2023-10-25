<script setup>
import { ref } from 'vue'
import axios  from 'axios'

const username = ref('')
const password = ref('')

function login() {
  console.log('username: ', username.value, ', password: ', password.value)
  axios.post('/api/login', {
    username: username.value,
    password: password.value
  }).then(function (response) {
    console.log(response)
    if (response.data.code == 0) {
      console.log('登录成功')
      window.location.href = '/'
    } else {
      console.log('登录失败')
      alert("登录失败: 用户名或密码错误！")
    }
  }).catch(function (error) {
    console.log(error);
    alert("登录失败: 用户名或密码错误！")
  });
}
</script>

<template>
  <div id="backgroud">
    <form id="login" @keyup.enter="login">
      <h1>登录</h1>
      <fieldset id="inputs">
        <input v-model="username" id="username" type="text" placeholder="用户名" autofocus required>
        <input v-model="password" id="password" type="password" placeholder="密码" required>
      </fieldset>
      <fieldset id="actions">
        <input id="submit" @click="login" type="button" value="登录">
      </fieldset>
    </form>
  </div>

</template>

<style scoped>
html, body
{
  height: 100%;
}

#backgroud
{
  height: 100%;
  width: 100%;
  font: 12px 'Lucida Sans Unicode', 'Trebuchet MS', Arial, Helvetica;
  margin: 0;
  background-color: #d9dee2;
  background-image: -webkit-gradient(linear, left top, left bottom, from(#ebeef2), to(#d9dee2));
  background-image: -webkit-linear-gradient(top, #ebeef2, #d9dee2);
  background-image: -moz-linear-gradient(top, #ebeef2, #d9dee2);
  background-image: -ms-linear-gradient(top, #ebeef2, #d9dee2);
  background-image: -o-linear-gradient(top, #ebeef2, #d9dee2);
  background-image: linear-gradient(top, #ebeef2, #d9dee2);
}

#login
{
  background-color: #fff;
  background-image: -webkit-gradient(linear, left top, left bottom, from(#fff), to(#eee));
  background-image: -webkit-linear-gradient(top, #fff, #eee);
  background-image: -moz-linear-gradient(top, #fff, #eee);
  background-image: -ms-linear-gradient(top, #fff, #eee);
  background-image: -o-linear-gradient(top, #fff, #eee);
  background-image: linear-gradient(top, #fff, #eee);
  height: 240px;
  width: 400px;
  margin: -150px 0 0 -230px;
  padding: 30px;
  position: absolute;
  top: 50%;
  left: 50%;
  z-index: 0;
  -moz-border-radius: 3px;
  -webkit-border-radius: 3px;
  border-radius: 3px;
  -webkit-box-shadow:
      0 0 2px rgba(0, 0, 0, 0.2),
      0 1px 1px rgba(0, 0, 0, .2),
      0 3px 0 #fff,
      0 4px 0 rgba(0, 0, 0, .2),
      0 6px 0 #fff,
      0 7px 0 rgba(0, 0, 0, .2);
  -moz-box-shadow:
      0 0 2px rgba(0, 0, 0, 0.2),
      1px 1px   0 rgba(0,   0,   0,   .1),
      3px 3px   0 rgba(255, 255, 255, 1),
      4px 4px   0 rgba(0,   0,   0,   .1),
      6px 6px   0 rgba(255, 255, 255, 1),
      7px 7px   0 rgba(0,   0,   0,   .1);
  box-shadow:
      0 0 2px rgba(0, 0, 0, 0.2),
      0 1px 1px rgba(0, 0, 0, .2),
      0 3px 0 #fff,
      0 4px 0 rgba(0, 0, 0, .2),
      0 6px 0 #fff,
      0 7px 0 rgba(0, 0, 0, .2);
}

#login:before
{
  content: '';
  position: absolute;
  z-index: -1;
  border: 1px dashed #ccc;
  top: 5px;
  bottom: 5px;
  left: 5px;
  right: 5px;
  -moz-box-shadow: 0 0 0 1px #fff;
  -webkit-box-shadow: 0 0 0 1px #fff;
  box-shadow: 0 0 0 1px #fff;
}

/*--------------------*/

h1
{
  text-shadow: 0 1px 0 rgba(255, 255, 255, .7), 0px 2px 0 rgba(0, 0, 0, .5);
  text-transform: uppercase;
  text-align: center;
  color: #666;
  margin: 0 0 30px 0;
  letter-spacing: 4px;
  font: normal 26px/1 Verdana, Helvetica;
  position: relative;
}

h1:after, h1:before
{
  background-color: #777;
  content: "";
  height: 1px;
  position: absolute;
  top: 15px;
  width: 120px;
}

h1:after
{
  background-image: -webkit-gradient(linear, left top, right top, from(#777), to(#fff));
  background-image: -webkit-linear-gradient(left, #777, #fff);
  background-image: -moz-linear-gradient(left, #777, #fff);
  background-image: -ms-linear-gradient(left, #777, #fff);
  background-image: -o-linear-gradient(left, #777, #fff);
  background-image: linear-gradient(left, #777, #fff);
  right: 0;
}

h1:before
{
  background-image: -webkit-gradient(linear, right top, left top, from(#777), to(#fff));
  background-image: -webkit-linear-gradient(right, #777, #fff);
  background-image: -moz-linear-gradient(right, #777, #fff);
  background-image: -ms-linear-gradient(right, #777, #fff);
  background-image: -o-linear-gradient(right, #777, #fff);
  background-image: linear-gradient(right, #777, #fff);
  left: 0;
}

/*--------------------*/

fieldset
{
  border: 0;
  padding: 0;
  margin: 0;
}

/*--------------------*/

#inputs input
{
  background: #f1f1f1;
  padding: 15px 15px 15px 30px;
  margin: 0 0 10px 0;
  width: 353px; /* 353 + 2 + 45 = 400 */
  border: 1px solid #ccc;
  -moz-border-radius: 5px;
  -webkit-border-radius: 5px;
  border-radius: 5px;
  -moz-box-shadow: 0 1px 1px #ccc inset, 0 1px 0 #fff;
  -webkit-box-shadow: 0 1px 1px #ccc inset, 0 1px 0 #fff;
  box-shadow: 0 1px 1px #ccc inset, 0 1px 0 #fff;
}

#username
{
  background-position: 5px -2px !important;
}

#password
{
  background-position: 5px -52px !important;
}

#inputs input:focus
{
  background-color: #fff;
  border-color: #e8c291;
  outline: none;
  -moz-box-shadow: 0 0 0 1px #e8c291 inset;
  -webkit-box-shadow: 0 0 0 1px #e8c291 inset;
  box-shadow: 0 0 0 1px #e8c291 inset;
}

/*--------------------*/
#actions
{
  margin: 0 auto 0 115px;
}

#submit
{
  background-color: #4e4bff;
  background-image: -webkit-gradient(linear, left top, left bottom, from(#4e4bff), to(#0f0e46));
  background-image: -webkit-linear-gradient(top, #4e4bff, #0f0e46);
  background-image: -moz-linear-gradient(top, #4e4bff, #0f0e46);
  background-image: -ms-linear-gradient(top, #4e4bff, #0f0e46);
  background-image: -o-linear-gradient(top, #4e4bff, #0f0e46);
  background-image: linear-gradient(top, #4e4bff, #0f0e46);

  -moz-border-radius: 3px;
  -webkit-border-radius: 3px;
  border-radius: 3px;

  /*text-shadow: 0 1px 0 rgba(255,255,255,0.5);*/

  -moz-box-shadow: 0 0 1px rgba(0, 0, 0, 0.3), 0 1px 0 rgba(255, 255, 255, 0.3) inset;
  -webkit-box-shadow: 0 0 1px rgba(0, 0, 0, 0.3), 0 1px 0 rgba(255, 255, 255, 0.3) inset;
  box-shadow: 0 0 1px rgba(0, 0, 0, 0.3), 0 1px 0 rgba(255, 255, 255, 0.3) inset;

  border-width: 1px;
  border-style: solid;
  border-color: #4533eb #4533eb #4533eb #4533eb;

  float: left;
  height: 35px;
  padding: 0;
  width: 120px;
  cursor: pointer;
  font: bold 15px Arial, Helvetica;
  color: #e0ddd7;
}

#submit:hover,#submit:focus
{
  background-color: #4e4bff;
  background-image: -webkit-gradient(linear, left top, left bottom, from(#0f0e46), to(#4e4bff));
  background-image: -webkit-linear-gradient(top, #0f0e46, #4e4bff);
  background-image: -moz-linear-gradient(top, #0f0e46, #4e4bff);
  background-image: -ms-linear-gradient(top, #0f0e46, #4e4bff);
  background-image: -o-linear-gradient(top, #0f0e46, #4e4bff);
  background-image: linear-gradient(top, #0f0e46, #4e4bff);
}

#submit:active
{
  outline: none;

  -moz-box-shadow: 0 1px 4px rgba(0, 0, 0, 0.5) inset;
  -webkit-box-shadow: 0 1px 4px rgba(0, 0, 0, 0.5) inset;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.5) inset;
}

#submit::-moz-focus-inner
{
  border: none;
}

#actions a
{
  color: #3151A2;
  float: right;
  line-height: 35px;
  margin-left: 10px;
}

/*--------------------*/

#back
{
  display: block;
  text-align: center;
  position: relative;
  top: 60px;
  color: #999;
}
</style>
