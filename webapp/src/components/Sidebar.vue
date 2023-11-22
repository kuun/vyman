<script setup>
import {ref} from "vue";
import {SideMenu} from "v3-easyui";
import router from "../router";

const menus = ref([
  {
    text: "网口管理",
    state: 'open',
    iconCls: 'fa fa-ethernet',
    children: [{
      text: 'IP管理',
      path: '/ui/interface/ethernet'
    }]
  }, {
    text: '路由设置',
    iconCls: 'fa fa-network-wired',
    children: [{
      text: '静态路由',
      path: '/ui/router/static'
    }]
  }, {
    text: "对象组管理",
    iconCls: 'fa fa-object-group',
    children: [{
      text: 'IPv4地址组',
      path: '/ui/group/ipv4-group'
    }, {
      text: 'IPv6地址组',
      path: '/ui/group/ipv6-group'
    }]
  }, {
    text: "安全域设置",
    iconCls: 'fa fa-shield-alt',
    children: [{
      text: '安全域',
      path: '/ui/zone'
    }]
  }, {
    text: "VPN",
    iconCls: 'fa fa-globe',
    children: [
      {
        text: 'IPSec',
        path: '/ui/vpn/ipsec'
      }, {
        text: 'L2TP',
        path: '/ui/vpn/l2tp'
      }, {
        text: 'PPTP-Server',
        path: '/ui/vpn/pptp-server'
      }, {
        text: 'OpenConnect',
        path: '/ui/vpn/openconnect'
      }, {
        text: 'SSTP',
        path: '/ui/vpn/sstp'
    }]
  }
])

const selectMenu = (event) => {
  if (event.path !== undefined && event.path !== '') {
    router.push(event.path)
  }
};

</script>

<template>
  <SideMenu style="width: 100%;"
            :data="menus"
            @itemClick="selectMenu" :border="false">
  </SideMenu>
</template>

<style scoped>

</style>
