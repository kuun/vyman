<script setup>
import {ref, watch} from "vue";
import {GridColumn, LinkButton, Panel} from "v3-easyui";
import CheckGrid from "./CheckGrid.vue";
import config from "../utils/config";
import _ from "lodash";
import {useGroupStore} from "../stores/group";


const groupStore = useGroupStore();
const editTitle = ref("");
const editDlg = ref(null);
const formRules = ref({
  address: 'required',
});
const editing = ref({});
const itemList = ref([]);

const getError = (key) => {

}

watch(() => groupStore.selectedGroup, (newVal, oldVal) => {
  if (newVal !== null) {
    refreshAddress();
  }
})

const saveAddress = () => {
  const value = editing.value;
  const selectedGroup = groupStore.selectedGroup;
  const cmds = [
    config.command('set', `firewall group address-group ${selectedGroup.id} address`, value.address),
  ];
  config.configure(cmds)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('保存地址失败.')
          return;
        }
        closeDlg();
        refreshAddress();
      })
      .catch((resp) => {
        alert('保存地址失败.')
      });
}

const closeDlg = () => {
  editing.value = {};
  editDlg.value.close();
}

const addAddress = () => {
  editTitle.value = "添加地址";
  editing.value = {};
  editDlg.value.open();
}

const modifyAddress = () => {
  if (itemList.value.length === 0) {
    alert('请选择地址');
    return;
  }
  editTitle.value = "修改地址";
  editing.value = _.cloneDeep(itemList.value[0]);
  oldIp.value = _.cloneDeep(itemList.value[0]);
  editDlg.value.open();
}

const deleteAddress = () => {
  const selections = itemList.value.filter(row => row.selected);
  if (selections.length < 1) {
    alert('至少选择一个地址组');
    return;
  }
  const selectedGroup = groupStore.selectedGroup;
  const cmds = selections.map(row => {
    return config.command('delete', `firewall group address-group ${selectedGroup.id} address`,  row.address);
  });
  config.configure(cmds)
      .then((resp) => {
        if (!resp.data.success) {
          console.log(resp.data)
          alert('删除地址失败.')
          return;
        }
        refreshAddress();
      })
      .catch((resp) => {
        alert('删除地址失败.')
      });
}

const refreshAddress = () => {
  // refresh address in selected group
  const selectedGroup = groupStore.selectedGroup;
  if (selectedGroup && selectedGroup.id) {
    config.returnValues(`firewall group address-group ${selectedGroup.id} address`)
        .then((resp) => {
          if (!resp.data.success) {
            console.log(resp.data)
            alert('获取地址失败.')
            return;
          }
          const data = resp.data.data;
          if (!data) {
            return;
          }
          itemList.value = _.map(data, (value) => {
            return {
              address: value,
            }
          });
        })
        .catch((resp) => {
          alert('获取地址失败.')
        });
  } else {
    alert('请选择地址组');
    itemList.value = [];
  }
}

</script>


<template>
  <Panel style="width:100%;height:100%" :border="false" title="地址列表">
    <div class="dialog-toolbar">
      <LinkButton iconCls="icon-add" style="width:80px" :plain="true" @click="addAddress">添加</LinkButton>
      <LinkButton iconCls="icon-edit" style="width:80px" :plain="true" @click="modifyAddress">修改</LinkButton>
      <LinkButton iconCls="icon-remove" style="width:80px" :plain="true" @click="deleteAddress">删除</LinkButton>
      <LinkButton iconCls="icon-reload" style="width:80px" :plain="true" @click="refreshAddress">刷新</LinkButton>
    </div>
    <Dialog ref="editDlg" bodyCls="f-column" :title="editTitle" :draggable="true" :modal="true" closed :dialogStyle="{height:'180px', width:'350px'}">
      <div class="f-full" style="overflow:auto">
        <Form ref="groupForm" :model="editing" :rules="formRules" @validate="errors=$event" style="padding:20px 40px">
          <div>
            <Label for="address">地址:</Label>
            <TextBox name="address" v-model="editing.address"></TextBox>
            <div class="error">{{getError('address')}}</div>
          </div>
        </Form>
      </div>
      <div class="dialog-button f-noshrink">
        <LinkButton @click="saveAddress" style="width: 80px">保存</LinkButton>
        <LinkButton @click="closeDlg" style="width: 80px">取消</LinkButton>
      </div>
    </Dialog>
    <CheckGrid :data="itemList">
      <GridColumn field="address" title="地址"></GridColumn>
    </CheckGrid>
  </Panel>
</template>