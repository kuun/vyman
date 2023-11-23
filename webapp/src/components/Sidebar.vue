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
      iconCls: 'fa fa-newspaper remove-tree-file-icon',
      text: 'IP管理',
      path: '/ui/interface/ethernet'
    }]
  }, {
    text: '路由设置',
    iconCls: 'fa fa-network-wired',
    children: [{
      iconCls: 'fa-solid fa-route remove-tree-file-icon',
      text: '静态路由',
      path: '/ui/router/static'
    }]
  }, {
    text: "对象组管理",
    iconCls: 'fa fa-object-group',
    children: [{
      iconCls: 'fa fa-group-arrows-rotate remove-tree-file-icon',
      text: 'IPv4地址组',
      path: '/ui/group/ipv4-address-group'
    }, {
      iconCls: 'fa fa-group-arrows-rotate remove-tree-file-icon',
      text: 'IPv6地址组',
      path: '/ui/group/ipv6-address-group'
    }, {
      iconCls: 'fa fa-layer-group remove-tree-file-icon',
      text: 'IPv4网络组',
      path: '/ui/group/ipv4-network-group'
    }, {
      iconCls: 'fa fa-layer-group remove-tree-file-icon',
      text: 'IPv6网络组',
      path: '/ui/group/ipv6-network-group'
    }, {
      iconCls: 'fa fa-object-group remove-tree-file-icon',
      text: '端口组',
      path: '/ui/group/port-group'
    }]
  }, {
    text: "防火墙",
    iconCls: 'fa fa-shield-alt',
    children: [{
      iconCls: 'fa-solid fa-link remove-tree-file-icon',
      text: '策略',
      path: '/ui/firewall/rule'
    }, {
      iconCls: 'fa-solid fa-money-bill remove-tree-file-icon',
      text: '安全域',
      path: '/ui/firewall/zone'
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

<style>
.remove-tree-file-icon {
  background: none !important;
}
</style>
