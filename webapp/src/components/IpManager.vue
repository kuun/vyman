<script setup>

import {ref, watch} from "vue";
import {DataGrid, GridColumn, LinkButton, Panel} from "v3-easyui";
import axios from "axios";
import {useIfaceStore} from "../stores/iface";

const ipAddrList = ref([]);
const ifaceStore = useIfaceStore();

watch(() => ifaceStore.selectedIface, (newVal, oldVal) => {
  console.log('selected iface changed: ', newVal, oldVal)
  if (newVal !== null) {
    axios.get(`/api/interfaces/${newVal.type}/${newVal.name}/address`)
        .then((resp) => {
          console.log('ip list: ', resp.data)
          ipAddrList.value = resp.data
        })
        .catch((resp) => {
          alert('获取IP列表失败.')
        })
  }
})
</script>

<template>
  <Panel :border="false">
    <div class="dialog-toolbar">
      <LinkButton iconCls="icon-add" style="width:80px" :plain="true">添加</LinkButton>
      <LinkButton iconCls="icon-edit" style="width:80px" :plain="true">修改</LinkButton>
      <LinkButton iconCls="icon-remove" style="width:80px" :plain="true">删除</LinkButton>
      <LinkButton iconCls="icon-reload" style="width:80px" :plain="true">刷新</LinkButton>
    </div>
    <DataGrid :data="ipAddrList" selection-mode="multiple" style="height: 100%">
      <GridColumn field="local" title="IP"></GridColumn>
      <GridColumn field="prefixlen" title="掩码" width="100"></GridColumn>
    </DataGrid>
  </Panel>
</template>

<style scoped>

</style>
