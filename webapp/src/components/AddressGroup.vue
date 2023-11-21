<script setup>
import {onMounted, ref} from "vue";
import {GridColumn, LinkButton, Panel} from "v3-easyui";
import CheckGrid from "./CheckGrid.vue";
import config from "../utils/config";
import _ from "lodash";

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
})

const getError = (key) => {

}

const saveGroup = () => {

}

const closeGroupDlg = () => {

}

const addGroup = () => {

}

const modifyGroup = () => {

}

const deleteGroup = () => {

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
    <CheckGrid :data="groupList" selectionMode="single">
      <GridColumn field="id" title="ID" width="100"></GridColumn>
      <GridColumn field="description" title="备注"></GridColumn>
    </CheckGrid>
  </Panel>
</template>