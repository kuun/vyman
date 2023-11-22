<script setup>
import {onMounted, ref} from "vue";
import {Form, FormField, GridColumn, LinkButton, Panel, TagBox, TextBox} from "v3-easyui";
import CheckGrid from "./CheckGrid.vue";
import config from "../utils/config";
import _ from "lodash";
import _alert from "../utils/alert";
import axios from "axios";

const zoneList = ref([]);
const zoneEditTitle = ref("");
const zoneEditDlg = ref(null);
const zoneFormRules = ref({
  name: 'required',
  description: 'required'
});
const editingZone = ref({});
const ifaceList = ref([]);
let useIfaceList = ref([]);


const addZone = async () => {
  zoneEditTitle.value = "添加安全域";
  await loadIfaceList();
  editingZone.value = {};
  zoneEditDlg.value.open();
}

const saveZone = () => {
  const value = editingZone.value;
  const cmds = [];
  if (zoneEditTitle.value === "修改安全域") {
    cmds.push(config.command('delete', `firewall zone ${value.name} interface`));
  }
  cmds.push(config.command('set', 'firewall zone', value.name));
  cmds.push(config.command('set', `firewall zone ${value.name} description`, value.description));
  _.each(value.iface, (iface) => {
    cmds.push(config.command('set', `firewall zone ${value.name} interface`, iface));
  });
  config.configure(cmds)
    .then((resp) => {
      if (!resp.data.success) {
        console.log(resp.data)
        _alert.alertError('保存安全域失败.')
        return;
      }
      closeZoneDlg();
      refreshZone();
    })
    .catch((resp) => {
      _alert.alertError('保存安全域失败.')
    });
}

const modifyZone = async () => {
  const selected = zoneList.value.filter((v) => v.selected);
  if (selected.length === 0) {
    _alert.alertWarning('请选择要修改的安全域.')
    return;
  }
  // editingZone.value = _.cloneDeep(selected[0]);
  await loadIfaceList();
  const curIface = selected[0].iface.split(",");
  _.each(curIface, (iface) => {
    ifaceList.value.push({value: iface, text: iface});
  });
  editingZone.value = {
    name: selected[0].name,
    description: selected[0].description,
    iface: curIface
  };
  zoneEditTitle.value = "修改安全域";
  zoneEditDlg.value.open();
}

const deleteZone = () => {
  const selected = zoneList.value.filter((v) => v.selected);
  if (selected.length === 0) {
    _alert.alertWarning('请选择要删除的安全域.')
    return;
  }
  const cmds = [
    config.command('delete', 'firewall zone', selected[0].name)
  ];
  config.configure(cmds)
    .then((resp) => {
      if (!resp.data.success) {
        console.log(resp.data)
        _alert.alertError('删除安全域失败.')
        return;
      }
      refreshZone();
    })
    .catch((resp) => {
      _alert.alertError('删除安全域失败.')
    });
}

const refreshZone = () => {
  config.showConfig('firewall zone')
    .then((resp) => {
      useIfaceList.value = [];
      if (resp.data.data !== null) {
        zoneList.value = _.map(resp.data.data.zone, (v, k) => {
          if (v.interface instanceof Array) {
            _.each(v.interface, (iface) => {
              useIfaceList.value.push({value: iface, text: iface});
            })
          } else {
            useIfaceList.value.push({value: v.interface, text: v.interface});
          }
          return {
            name: k,
            iface: (v.interface instanceof Array) ? _.join(v.interface, ",") : v.interface,
            description: v.description
          }
        });
        return;
      }
      zoneList.value = [];
    })
    .catch((resp) => {
      _alert.alertError('获取安全域列表失败.')
    })
}

const loadIfaceList = async () => {
  await axios.get('/api/interfaces/ethernet')
    .then((resp) => {
      ifaceList.value = _.map(resp.data, (v) => {
        return {
          value: v.name,
          text: v.name
        }
      });
      _.pullAllWith(ifaceList.value, useIfaceList.value, _.isEqual);
    }).catch((resp) => {
      alert('获取网卡列表失败.')
    })
}

const closeZoneDlg = () => {
  editingZone.value = {};
  zoneEditDlg.value.close();
}

onMounted(() => {
  refreshZone();
})

</script>

<template>
  <Panel style="width:100%;height:100%" :border="false" title="安全域列表">
    <div class="dialog-toolbar">
      <LinkButton iconCls="icon-add" style="width:80px" :plain="true" @click="addZone">添加</LinkButton>
      <LinkButton iconCls="icon-edit" style="width:80px" :plain="true" @click="modifyZone">修改</LinkButton>
      <LinkButton iconCls="icon-remove" style="width:80px" :plain="true" @click="deleteZone">删除</LinkButton>
      <LinkButton iconCls="icon-reload" style="width:80px" :plain="true" @click="refreshZone">刷新</LinkButton>
    </div>
    <Dialog ref="zoneEditDlg" bodyCls="f-column" :title="zoneEditTitle" :draggable="true" :modal="true" closed
            :dialogStyle="{height:'300px', width:'450px'}">
      <div class="f-full" style="overflow:auto">
        <Form ref="zoneForm" :model="editingZone" :rules="zoneFormRules" @validate="errors=$event"
              style="padding:20px 40px" :labelWidth="80" errorType="tooltip">
          <FormField name="name" label="名称:">
            <TextBox v-model="editingZone.name"></TextBox>
          </FormField>
          <FormField name="iface" label="网卡列表:">
            <TagBox v-model="editingZone.iface" valueField="value" textField="text" placeholder="选择网卡"
                    :limitToList="true" :hasDownArrow="true" :data="ifaceList">
            </TagBox>
          </FormField>
          <FormField name="description" label="备注:">
            <TextBox v-model="editingZone.description"></TextBox>
          </FormField>
        </Form>
      </div>
      <div class="dialog-button f-noshrink">
        <LinkButton @click="saveZone" style="width: 80px">保存</LinkButton>
        <LinkButton @click="closeZoneDlg" style="width: 80px">取消</LinkButton>
      </div>
    </Dialog>
    <CheckGrid :data="zoneList" selectionMode="single" @selection-change="onSelectionChange">
      <GridColumn field="name" title="名称" width="100"></GridColumn>
      <GridColumn field="iface" title="网卡列表"></GridColumn>
      <GridColumn field="description" title="备注"></GridColumn>
    </CheckGrid>
  </Panel>
</template>