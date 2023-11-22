<script setup>
import {ref, watch} from "vue";
import {GridColumn, LinkButton, Panel} from "v3-easyui";
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
const editTitle = ref("");
const editDlg = ref(null);
const formRules = ref({
  content: 'required',
});
const editing = ref({});
const itemList = ref([]);

const getError = (key) => {

}

watch(() => groupStore.selectedGroup, (newVal, oldVal) => {
  if (newVal !== null) {
    refreshItem();
  } else {
    itemList.value = [];
  }
})

const saveItem = () => {
  const value = editing.value;
  const selectedGroup = groupStore.selectedGroup;
  const cmds = [
    config.command('set', `firewall group ${props.groupType}-group ${selectedGroup.id} ${props.groupType}`, value.content),
  ];
  config.configure(cmds)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('保存失败.')
          return;
        }
        closeDlg();
        refreshItem();
      })
      .catch((resp) => {
        alert('保存失败.')
      });
}

const closeDlg = () => {
  editing.value = {};
  editDlg.value.close();
}

const addItem = () => {
  editTitle.value = "添加";
  editing.value = {};
  editDlg.value.open();
}

const modifyItem = () => {
  if (itemList.value.length !== 1) {
    alert('请选择一条数据');
    return;
  }
  editTitle.value = "修改";
  editing.value = _.cloneDeep(itemList.value[0]);
  oldIp.value = _.cloneDeep(itemList.value[0]);
  editDlg.value.open();
}

const deleteItem = () => {
  const selections = itemList.value.filter(row => row.selected);
  if (selections.length < 1) {
    alert('至少选择一条数据');
    return;
  }
  const selectedGroup = groupStore.selectedGroup;
  const cmds = selections.map(row => {
    return config.command('delete', `firewall group ${props.groupType}-group ${selectedGroup.id} ${props.groupType}`,  row.content);
  });
  config.configure(cmds)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('删除失败.')
          return;
        }
        refreshItem();
      })
      .catch((resp) => {
        alert('删除失败.')
      });
}

const refreshItem = () => {
  // refresh items in selected group
  const selectedGroup = groupStore.selectedGroup;
  if (selectedGroup && selectedGroup.id) {
    config.returnValues(`firewall group ${props.groupType}-group ${selectedGroup.id} ${props.groupType}`)
        .then((resp) => {
          if (!resp.data.success) {
            console.log(resp.data)
            alert('获取失败.')
            return;
          }
          const data = resp.data.data;
          if (!data) {
            return;
          }
          itemList.value = _.map(data, (value) => {
            return {
              content: value,
            }
          });
        })
        .catch((resp) => {
          alert('获取失败.')
        });
  } else {
    alert('请选择组');
    itemList.value = [];
  }
}

</script>


<template>
  <Panel style="width:100%;height:100%" :border="false" title="地址列表">
    <div class="dialog-toolbar">
      <LinkButton iconCls="icon-add" style="width:80px" :plain="true" @click="addItem">添加</LinkButton>
      <LinkButton iconCls="icon-edit" style="width:80px" :plain="true" @click="modifyItem">修改</LinkButton>
      <LinkButton iconCls="icon-remove" style="width:80px" :plain="true" @click="deleteItem">删除</LinkButton>
      <LinkButton iconCls="icon-reload" style="width:80px" :plain="true" @click="refreshItem">刷新</LinkButton>
    </div>
    <Dialog ref="editDlg" bodyCls="f-column" :title="editTitle" :draggable="true" :modal="true" closed :dialogStyle="{height:'180px', width:'350px'}">
      <div class="f-full" style="overflow:auto">
        <Form ref="groupForm" :model="editing" :rules="formRules" @validate="errors=$event" style="padding:20px 40px" :labelWidth="80" errorType="tooltip">
          <FormField name="content" label="地址:">
            <TextBox v-model="editing.content"></TextBox>
          </FormField>
        </Form>
      </div>
      <div class="dialog-button f-noshrink">
        <LinkButton @click="saveItem" style="width: 80px">保存</LinkButton>
        <LinkButton @click="closeDlg" style="width: 80px">取消</LinkButton>
      </div>
    </Dialog>
    <CheckGrid :data="itemList">
      <GridColumn field="content" title="内容"></GridColumn>
    </CheckGrid>
  </Panel>
</template>