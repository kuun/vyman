<script setup>
import {onMounted, ref, watch} from "vue";
import {CheckBox, ComboBox, Form, FormField, LinkButton, TextBox, Panel, Label} from "v3-easyui";
import {useIfaceStore} from "../stores/iface";
import config from "@/utils/config";
import alert from "@/utils/alert";

const ifaceStore = useIfaceStore();
const nicConfig = ref({});
const arp = ref({});

const duplexMode = [
  {text: '自动', value: 'auto'},
  {text: '全双工', value: 'full'},
  {text: '半双工', value: 'half'}
]

const saveDHCP = () => {
    const iface = ifaceStore.selectedIface;
    const op = nicConfig.value.dhcp ? 'set' : 'delete';
    const cmds = [
        config.command(op, `interfaces ${iface.type} ${iface.name} address`, 'dhcp')
    ];
    config.configure(cmds)
        .then((resp) => {
            if (resp.data.success) {
                alert.alertSuccess('保存成功');
            }
        })
}

watch(() => ifaceStore.selectedIface, (newVal, oldVal) => {
    if (newVal !== null) {
        // loadIp(newVal);
        nicConfig.value.linkStatus = newVal.linkStatus === 'UP' ? '已连接' : '未连接';
        loadDHCP(newVal);
    }
})

const loadDHCP = (iface) => {
    config.showConfig(`interfaces ${iface.type} ${iface.name} address`)
        .then((resp) => {
            if (resp.data.data == null) {
                nicConfig.value.dhcp = false;
                return;
            }
            nicConfig.value.dhcp = resp.data.data.address === 'dhcp';
        })
}

</script>

<template>
    <div>
        <Panel title="其他配置" style="height: 300px">
            <Form :model="nicConfig" labelAlign="right" :label-width="100" style="max-width:500px; padding-top: 10px">
                <FormField name="name" label="网卡状态: ">
                    <TextBox v-model="nicConfig.linkStatus" disabled></TextBox>
                </FormField>
                <FormField label="DHCP: ">
                    <CheckBox v-model="nicConfig.dhcp" @click="saveDHCP"></CheckBox>
                    <Label class="label-dhcp" v-if="nicConfig.dhcp">关闭</Label>
                    <Label class="label-dhcp" v-else>开启</Label>
                </FormField>
<!--                <FormField name="name" label="MTU: ">
                    <TextBox v-model="nicConfig.mtu"></TextBox>
                </FormField>
                <FormField label="双工模式: ">
                    <ComboBox :data="duplexMode" v-model="nicConfig.duplex"></ComboBox>
                </FormField>-->
            </Form>
        </Panel>
<!--        <Panel title="ARP配置" style="height:200px">
            <Form :model="arp" labelAlign="right" :label-width="100" style="max-width:500px; padding-top: 10px">
                <FormField name="name" label="缓存时间: ">
                    <TextBox v-model="arp.cacheTimeout"></TextBox>
                </FormField>
                <FormField name="name" label="禁用过滤: ">
                    <CheckBox v-model="arp.diableFilter"></CheckBox>
                </FormField>
                <FormField name="name" label="ARP代理: ">
                    <CheckBox v-model="arp.proxy"></CheckBox>
                </FormField>

                <FormField>
                    <LinkButton :disabled="false" @click="saveArp" style="width:100px">提交</LinkButton>
                </FormField>
            </Form>
        </Panel>-->

    </div>
</template>

<style scoped>
.form-field {
     margin-bottom: 5px;
}

.label-dhcp {
    margin-left: 5px;
}
</style>