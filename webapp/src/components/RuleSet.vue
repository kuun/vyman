<script setup>
import {onMounted, ref} from "vue";
import {Form, FormField, GridColumn, LinkButton, Panel, ComboBox} from "v3-easyui";
import CheckGrid from "./CheckGrid.vue";
import config from "../utils/config";
import _ from "lodash";
import {useRuleStore} from "../stores/rule";



const ruleStore = useRuleStore();
const editTitle = ref("");
const editDlg = ref(null);
const formRules = ref({
  id: 'required',
  description: 'required'
});
const editingRuleSet = ref({});
const rulesetList = ref([]);
const actions = ref([
  {value: 'accept', text: '允许'},
  {value: 'drop', text: '丢弃'},
  {value: 'reject', text: '拒绝'},
])

onMounted(() => {
  refresh();
  ruleStore.setSelectedRuleSet(null);
})

const getError = (key) => {

}

const saveGroup = () => {
  const value = editingRuleSet.value;
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
        closeDlg();
        refresh();
      })
      .catch((resp) => {
        alert('保存失败.')
      });
}

const closeDlg = () => {
  editingRuleSet.value = {};
  editDlg.value.close();
}

const add = () => {
  editTitle.value = "添加";
  editingRuleSet.value = {};
  editDlg.value.open();
}

const modify = () => {
  const selections = rulesetList.value.filter(row => row.selected);
  if (selections.length !== 1) {
    alert('请选择一个策略集');
    return;
  }
  editTitle.value = "修改";
  editingRuleSet.value = selections[0];
  editDlg.value.open();
}

const deleteRuleSet = () => {
  const selections = rulesetList.value.filter(row => row.selected);
  if (selections.length < 1) {
    alert('至少选择一个策略集');
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
        refresh();
      })
      .catch((resp) => {
        alert('删除失败.')
      });
}

const refresh = () => {
  config.showConfig(`firewall name`)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('获取策略集列表失败.')
          return;
        }
        if (resp.data.data) {
          rulesetList.value = _.map(resp.data.data['name'], (value, key) => {
            console.log(key, value)
            return {
              id: key,
              description: value.description
            };
          });
          groupStore.setSelectedGroup(null);
        }
      })
      .catch((resp) => {
        alert('获取策略集列表失败.')
      });
};

const onSelectionChange = (selection) => {
  console.log(selection)
  ruleStore.setSelectedRuleSet(selection);
}

</script>

<template>
  <Panel style="width:100%;height:100%" :border="false" title="策略集">
    <div class="dialog-toolbar">
      <LinkButton iconCls="icon-add" style="width:80px" :plain="true" @click="add">添加</LinkButton>
      <LinkButton iconCls="icon-edit" style="width:80px" :plain="true" @click="modify">修改</LinkButton>
      <LinkButton iconCls="icon-remove" style="width:80px" :plain="true" @click="deleteRuleSet">删除</LinkButton>
      <LinkButton iconCls="icon-reload" style="width:80px" :plain="true" @click="refresh">刷新</LinkButton>
    </div>
    <Dialog ref="editDlg" bodyCls="f-column" :title="editTitle" :draggable="true" :modal="true" closed :dialogStyle="{height:'310px', width:'350px'}">
      <div class="f-full" style="overflow:auto">
        <Form ref="groupForm" :model="editingRuleSet" :rules="formRules" @validate="errors=$event" style="padding:20px 40px" :labelWidth="80" errorType="tooltip">
          <FormField name="id" label="ID:">
            <TextBox v-model="editingRuleSet.id"></TextBox>
          </FormField>
          <FormField name="description" label="备注:">
            <TextBox v-model="editingRuleSet.description"></TextBox>
          </FormField>
          <FormField name="default-action" label="默认行为:">
            <ComboBox :data="actions" v-model="editingRuleSet['default-action']"></ComboBox>
          </FormField>
        </Form>
      </div>
      <div class="dialog-button f-noshrink">
        <LinkButton @click="saveGroup" style="width: 80px">保存</LinkButton>
        <LinkButton @click="closeDlg" style="width: 80px">取消</LinkButton>
      </div>
    </Dialog>
    <CheckGrid :data="rulesetList" selectionMode="single" @selection-change="onSelectionChange" >
      <GridColumn field="id" title="ID" width="100"></GridColumn>
      <GridColumn field="description" title="备注"></GridColumn>
      <GridColumn field="default-action" title="默认行为"></GridColumn>
    </CheckGrid>
  </Panel>
</template>