<script setup>

import {computed, nextTick, ref, watch} from "vue";
import {CheckBox, DataGrid, GridColumn, LinkButton, Panel} from "v3-easyui";
import axios from "axios";
import {useIfaceStore} from "../stores/iface";

const ipAddrList = ref([]);
const allChecked = ref(false);
const ifaceStore = useIfaceStore();

// 实现checkbox选中
const rowClicked = ref(false);
const checkedRows = computed(() => {
  return ipAddrList.value.filter(row => row.selected);
});
const onAllCheckedChange = (checked) => {
  if (rowClicked.value) {
    return;
  }
  ipAddrList.value.map(row => {
    row.selected = checked;
  });
};

const onCheckedChange = (checked) => {
  allChecked.value = checkedRows.value.length === ipAddrList.value.length;
  rowClicked.value = true;
  nextTick(() => (rowClicked.value = false));
}


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
    <DataGrid :data="ipAddrList" style="height: 100%">
      <GridColumn field="selected" :width="50" align="center">
        <template v-slot:header="scope">
          <CheckBox v-model="allChecked" @checkedChange="onAllCheckedChange($event)"></CheckBox>
        </template>
        <template v-slot:body="scope" @checkedChange="onCheckedChange($event)">
          <CheckBox v-model="scope.row.selected" @checkedChange="onCheckedChange($event)"></CheckBox>
        </template>
      </GridColumn>
      <GridColumn field="local" title="IP"></GridColumn>
      <GridColumn field="prefixlen" title="掩码" width="100"></GridColumn>
    </DataGrid>
  </Panel>
</template>

<style scoped>

</style>
