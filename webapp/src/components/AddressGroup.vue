<script setup>
import {onMounted, ref} from "vue";
import {GridColumn, LinkButton, Panel} from "v3-easyui";
import CheckGrid from "./CheckGrid.vue";
import config from "../utils/config";
import _ from "lodash";
import {useGroupStore} from "../stores/group";


const groupStore = useGroupStore();
const groupEditTitle = ref("");
const groupEditDlg = ref(null);
const groupFormRules = ref({
  id: 'required',
  description: 'required'
});
const editingGroup = ref({});
const groupList = ref([]);


onMounted(() => {
  refreshGroup();
  groupStore.setSelectedGroup(null);
})

const getError = (key) => {

}

const saveGroup = () => {
  const value = editingGroup.value;
  const cmds = [
    config.command('set', 'firewall group address-group', value.id),
    config.command('set', `firewall group address-group ${value.id} description`, value.description)
  ];
  config.configure(cmds)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('保存地址组失败.')
          return;
        }
        closeGroupDlg();
        refreshGroup();
      })
      .catch((resp) => {
        alert('保存地址组失败.')
      });
}

const closeGroupDlg = () => {
  editingGroup.value = {};
  groupEditDlg.value.close();
}

const addGroup = () => {
  groupEditTitle.value = "添加地址组";
  editingGroup.value = {};
  groupEditDlg.value.open();
}

const modifyGroup = () => {
  const selections = groupList.value.filter(row => row.selected);
  if (selections.length !== 1) {
    alert('请选择一个地址组');
    return;
  }
  groupEditTitle.value = "修改地址组";
  editingGroup.value = selections[0];
  groupEditDlg.value.open();
}

const deleteGroup = () => {
  const selections = groupList.value.filter(row => row.selected);
  if (selections.length < 1) {
    alert('至少选择一个地址组');
    return;
  }
  const cmds = _.map(selections, (value) => {
    return config.command('delete', `firewall group address-group ${value.id}`);
  });
  config.configure(cmds)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('删除地址组失败.')
          return;
        }
        refreshGroup();
      })
      .catch((resp) => {
        alert('删除地址组失败.')
      });
}

const refreshGroup = () => {
  config.showConfig('firewall group address-group')
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('获取地址组列表失败.')
          return;
        }
        if (resp.data.data) {
          groupList.value = _.map(resp.data.data['address-group'], (value, key) => {
            console.log(key, value)
            return {
              id: key,
              description: value.description
            };
          });
        }
      })
      .catch((resp) => {
        alert('获取地址组列表失败.')
      });
};

const onSelectionChange = (selection) => {
  console.log(selection)
  groupStore.setSelectedGroup(selection);
}

</script>

<template>
  <Panel style="width:100%;height:100%" :border="false" title="地址组列表">
    <div class="dialog-toolbar">
      <LinkButton iconCls="icon-add" style="width:80px" :plain="true" @click="addGroup">添加</LinkButton>
      <LinkButton iconCls="icon-edit" style="width:80px" :plain="true" @click="modifyGroup">修改</LinkButton>
      <LinkButton iconCls="icon-remove" style="width:80px" :plain="true" @click="deleteGroup">删除</LinkButton>
      <LinkButton iconCls="icon-reload" style="width:80px" :plain="true" @click="refreshGroup">刷新</LinkButton>
    </div>
    <Dialog ref="groupEditDlg" bodyCls="f-column" :title="groupEditTitle" :draggable="true" :modal="true" closed :dialogStyle="{height:'250px', width:'350px'}">
      <div class="f-full" style="overflow:auto">
        <Form ref="groupForm" :model="editingGroup" :rules="groupFormRules" @validate="errors=$event" style="padding:20px 40px">
          <div>
            <Label for="id">ID:</Label>
            <TextBox name="id" v-model="editingGroup.id"></TextBox>
            <div class="error">{{getError('id')}}</div>
          </div>
          <div>
            <Label for="description">备注:</Label>
            <TextBox name="description" v-model="editingGroup.description"></TextBox>
            <div class="error">{{getError('description')}}</div>
          </div>
        </Form>
      </div>
      <div class="dialog-button f-noshrink">
        <LinkButton @click="saveGroup" style="width: 80px">保存</LinkButton>
        <LinkButton @click="closeGroupDlg" style="width: 80px">取消</LinkButton>
      </div>
    </Dialog>
    <CheckGrid :data="groupList" selectionMode="single" @selection-change="onSelectionChange" >
      <GridColumn field="id" title="ID" width="100"></GridColumn>
      <GridColumn field="description" title="备注"></GridColumn>
    </CheckGrid>
  </Panel>
</template>