<script setup>
import {LinkButton, MenuItem, Panel, SplitButton} from "v3-easyui";
import axios from "axios";
import alert from "../utils/alert";
import {onMounted, ref} from "vue";

const user = ref("")

const logout = () => {
  alert.alertConfirm("退出登录", "确定退出登录吗?", r => {
      if (r) {
        axios.post("/api/logout").then(() => {
          window.location.href = "/ui/login";
        });
      }
    }
  )
}

const save = () => {
  alert.alertConfirm("保存配置", "确定要保存设备当前的配置吗?", r => {
      if (r) {
        //发送保存配置的请求
      }
    }
  )
}

onMounted(() => {
  axios.get("/api/login/user").then(resp => {
    if (resp.data.success) {
      user.value = resp.data.user
    } else {
      window.location.href = "/ui/login";
    }
  })
})

</script>

<template>
  <Panel :border="false">
    <template v-slot:header>
      <div>
        <img class="product-title" src="../assets/img/netguard.png">
      </div>
      <div class="right">
        <LinkButton :text="user" :plain="true" iconCls="icon-man"></LinkButton>
        <LinkButton iconCls="icon-save" v-Tooltip="{content:'保存'}" :plain="true" @click="save"></LinkButton>
        <LinkButton v-Tooltip="{content:'退出登录'}" :plain="true" @click="logout"><i class="fa fa-right-from-bracket"
                                                                                      style="color: #3b64ab;"></i>
        </LinkButton>
      </div>
    </template>
  </Panel>
</template>

<style scoped>
.product-title {
    height: 45px;
}

.right {
    position: absolute;
    right: 5px;
}
</style>
