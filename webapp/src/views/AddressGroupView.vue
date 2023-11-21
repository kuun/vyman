<script setup>
import {Layout, LayoutPanel, LinkButton, Form, TextBox, DataGrid, GridColumn} from "v3-easyui";
import {ref} from "vue";
import CheckGrid from "../components/CheckGrid.vue";
import AddressGroup from "../components/AddressGroup.vue";

const groupEditTitle = ref("");
const groupEditDlg = ref(null);
const groupFormRules = ref({
  id: 'required',
  description: 'required'
});
const editingGroup = ref({});
const groupList = ref([]);
const itemList = ref([]);
const errors = ref({});

const getError = (key) => {

}

const saveGroup = () => {

}

const closeGroupDlg = () => {

}

</script>

<template>
    <Layout style="width:100%;height:100%;">
      <LayoutPanel region="west" style="width:400px;height:100%">
        <AddressGroup></AddressGroup>
      </LayoutPanel>
      <LayoutPanel region="center" style="height:100%;width: 100%">
        <Panel style="width:100%;height:100%" :border="false" title="地址列表">
          <div class="dialog-toolbar">
            <LinkButton iconCls="icon-add" style="width:80px" :plain="true" @click="addIp">添加</LinkButton>
            <LinkButton iconCls="icon-edit" style="width:80px" :plain="true" @click="modifyIp">修改</LinkButton>
            <LinkButton iconCls="icon-remove" style="width:80px" :plain="true" @click="deleteIp">删除</LinkButton>
            <LinkButton iconCls="icon-reload" style="width:80px" :plain="true" @click="refreshIp">刷新</LinkButton>
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
          <CheckGrid :data="itemList">
            <GridColumn field="address" title="ID" width="200"></GridColumn>
            <GridColumn field="description" title="备注"></GridColumn>
          </CheckGrid>
        </Panel>
      </LayoutPanel>
    </Layout>
</template>