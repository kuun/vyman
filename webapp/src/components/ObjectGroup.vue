<script setup>
import {onMounted, ref} from "vue";
import {Form, FormField, GridColumn, LinkButton, Panel} from "v3-easyui";
import CheckGrid from "./CheckGrid.vue";
import config from "../utils/config";
import _ from "lodash";
import {useGroupStore} from "../stores/group";

const props = defineProps({
  groupType: {
    type: String,
    required: true
  }
});
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
    config.command('set', `firewall group ${props.groupType}-group`, value.id),
    config.command('set', `firewall group ${props.groupType}-group ${value.id} description`, value.description)
  ];
  config.configure(cmds)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('保存失败.')
          return;
        }
        closeGroupDlg();
        refreshGroup();
      })
      .catch((resp) => {
        alert('保存失败.')
      });
}

const closeGroupDlg = () => {
  editingGroup.value = {};
  groupEditDlg.value.close();
}

const addGroup = () => {
  groupEditTitle.value = "添加";
  editingGroup.value = {};
  groupEditDlg.value.open();
}

const modifyGroup = () => {
  const selections = groupList.value.filter(row => row.selected);
  if (selections.length !== 1) {
    alert('请选择一个组');
    return;
  }
  groupEditTitle.value = "修改";
  editingGroup.value = selections[0];
  groupEditDlg.value.open();
}

const deleteGroup = () => {
  const selections = groupList.value.filter(row => row.selected);
  if (selections.length < 1) {
    alert('至少选择一个组');
    return;
  }
  const cmds = _.map(selections, (value) => {
    return config.command('delete', `firewall group ${props.groupType}-group ${value.id}`);
  });
  config.configure(cmds)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('删除失败.')
          return;
        }
        refreshGroup();
      })
      .catch((resp) => {
        alert('删除失败.')
      });
}

const refreshGroup = () => {
  config.showConfig(`firewall group ${props.groupType}-group`)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('获取组列表失败.')
          return;
        }
        if (resp.data.data) {
          groupList.value = _.map(resp.data.data[`${props.groupType}-group`], (value, key) => {
            console.log(key, value)
            return {
              id: key,
              description: value.description
            };
          });
          groupStore.setSelectedGroup(null);
        } else {
          groupList.value = [];
        }
      })
      .catch((resp) => {
        alert('获取组列表失败.')
      });
};

const onSelectionChange = (selection) => {
  console.log(selection)
  groupStore.setSelectedGroup(selection);
}

</script>

<template>
  <Panel style="width:100%;height:100%" :border="false" title="组列表">
    <div class="dialog-toolbar">
      <LinkButton iconCls="icon-add" style="width:80px" :plain="true" @click="addGroup">添加</LinkButton>
      <LinkButton iconCls="icon-edit" style="width:80px" :plain="true" @click="modifyGroup">修改</LinkButton>
      <LinkButton iconCls="icon-remove" style="width:80px" :plain="true" @click="deleteGroup">删除</LinkButton>
      <LinkButton iconCls="icon-reload" style="width:80px" :plain="true" @click="refreshGroup">刷新</LinkButton>
    </div>
    <Dialog ref="groupEditDlg" bodyCls="f-column" :title="groupEditTitle" :draggable="true" :modal="true" closed :dialogStyle="{height:'250px', width:'350px'}">
      <div class="f-full" style="overflow:auto">
        <Form ref="groupForm" :model="editingGroup" :rules="groupFormRules" @validate="errors=$event" style="padding:20px 40px" :labelWidth="80" errorType="tooltip">
          <FormField name="id" label="ID:">
            <TextBox v-model="editingGroup.id"></TextBox>
          </FormField>
          <FormField name="description" label="备注:">
            <TextBox v-model="editingGroup.description"></TextBox>
          </FormField>
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