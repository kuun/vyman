<script setup>

import {computed, nextTick, ref, watch} from "vue";
import {CheckBox, DataGrid, Dialog, Form, GridColumn, LinkButton, Panel, TextBox} from "v3-easyui";
import axios from "axios";
import {useIfaceStore} from "../stores/iface";

const ifaceStore = useIfaceStore();
const ipAddrList = ref([]);
const allChecked = ref(false);
const editingIp = ref({});
const dlgTitle = ref("");
const formRules = ref({
  local: 'required',
  prefixlen: 'required'
});
const dlg = ref(null);
const errors = ref({});

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
  console.log('selected iface changed: ', newVal, oldVal);
  if (newVal !== null) {
    loadIp(newVal);
  }
})

const refreshIp = () => {
  loadIp(ifaceStore.selectedIface);
};

const loadIp = (iface) => {
  if (!iface.name) {
    alert('请选择网卡');
    return;
  }
  axios.get(`/api/interfaces/${iface.type}/${iface.name}/address`)
      .then((resp) => {
        console.log('ip list: ', resp.data)
        ipAddrList.value = resp.data
      })
      .catch((resp) => {
        alert('获取IP列表失败.')
      })
}

const addIp = () => {
  dlgTitle.value = '添加';
  editingIp.value = {
    local: '',
    prefixlen: 24,
  }
  dlg.value.open();
};

const modifyIp = () => {

};

const deleteIp = () => {

};

const saveIp = () => {

};

const close = () => {
  dlg.value.close();
}

const getError = (name) => {
  return errors.value[name] && errors.value[name].length
      ? errors.value[name][0]
      : null;
};

</script>

<template>
  <Panel :border="false">
    <div class="dialog-toolbar">
      <LinkButton iconCls="icon-add" style="width:80px" :plain="true" @click="addIp">添加</LinkButton>
      <LinkButton iconCls="icon-edit" style="width:80px" :plain="true" @click="modifyIp">修改</LinkButton>
      <LinkButton iconCls="icon-remove" style="width:80px" :plain="true" @click="deleteIp">删除</LinkButton>
      <LinkButton iconCls="icon-reload" style="width:80px" :plain="true" @click="refreshIp">刷新</LinkButton>
    </div>
    <Dialog ref="dlg" bodyCls="f-column" :title="dlgTitle" :modal="true" closed :dialogStyle="{height:'250px', width:'350px'}">
      <div class="f-full" style="overflow:auto">
        <Form ref="form" :model="deleteIp" :rules="formRules" @validate="errors=$event" style="padding:20px 40px">
          <div>
            <Label for="itemid">IP:</Label>
            <TextBox name="local" v-model="editingIp.local"></TextBox>
            <div class="error">{{getError('local')}}</div>
          </div>
          <div>
            <Label for="name">掩码:</Label>
            <TextBox name="prefixlen" v-model="editingIp.prefixlen"></TextBox>
            <div class="error">{{getError('prefixlen')}}</div>
          </div>
        </Form>
      </div>
      <div class="dialog-button f-noshrink">
        <LinkButton @click="saveIp" style="width: 80px">保存</LinkButton>
        <LinkButton @click="close" style="width: 80px">取消</LinkButton>
      </div>
    </Dialog>
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
      <GridColumn field="broadcast" title="广播地址" width="200"></GridColumn>
      <GridColumn field="dynamic" title="动态IP" width="100">
        <template v-slot:body="scope">
          <div>{{scope.row.dynamic ? '是' : '否'}}</div>
        </template>
      </GridColumn>
    </DataGrid>
  </Panel>
</template>

<style scoped>
.error {
  color: red;
  font-size: 12px;
  margin: 4px 0;
  margin-left: 80px;
}
</style>
