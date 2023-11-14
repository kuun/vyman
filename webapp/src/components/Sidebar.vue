<script setup>
import {ref} from "vue";
import {SideMenu, MenuItem, SubMenu} from "v3-easyui";
import router from "../router";

const menus = ref([
  {
    text: "网口管理",
    state: 'open',
    iconCls: 'icon-more',
    children: [{
      text: 'IP管理',
      path: '/ui/interface/ethernet'
    }]
  }, {
    text: '路由设置',
    iconCls: 'icon-more',
    children: [{
      text: '静态路由',
      path: '/ui/router/static'
    }]
  }, {
    text: "对象组管理",
    iconCls: 'icon-more',
    children: [{
      text: 'IPv4地址组',
      path: '/ui/group/ipv4-group'
    }, {
      text: 'IPv6地址组',
      path: '/ui/group/ipv6-group'
    }]
  }
])

function selectMenu(event) {
  if (event.path !== undefined && event.path !== '') {
    router.push(event.path)
  }
}

</script>

<template>

    <SideMenu style="width: 100%;"
              :data="menus"
              @selectionChange="selectMenu" :border="false">

    </SideMenu>

</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
