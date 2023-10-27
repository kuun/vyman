<script setup>
import {DataGrid, Layout, LayoutPanel, DataList} from "v3-easyui";
import {onMounted, ref} from "vue";
import axios from "axios";
import _ from "lodash";


const ifaceList = ref([])
const ipAddrList = ref([])

const props = defineProps({
  ifaceType: {
    type: String,
    required: true
  }
})

onMounted(() => {
  axios.get('/api/interfaces/' + props.ifaceType)
      .then((resp) => {
        console.log('interfaces: ', resp.data)
        _.map(resp.data, (iface) => {
          ifaceList.value.push({
            text: iface.name,
            value: iface
          })
        })
      })
      .catch((resp) => {
        alert('获取网卡列表失败.')
      })
})
</script>

<template>
  <Layout style="width:100%;height:100%;">
    <LayoutPanel region="west" style="width:150px;">
      <div class="title">
        <DataList :data="ifaceList" onSelectionChange="onSelectionChange"/>
      </div>
    </LayoutPanel>
    <LayoutPanel region="center" style="height:100%">
      <DataGrid :data="ipAddrList" selection-mode="multiple" style="height: 100%">
        <GridColumn field="ip" title="IP"></GridColumn>
        <GridColumn field="mask" title="掩码" width="100"></GridColumn>
      </DataGrid>
    </LayoutPanel>
  </Layout>

</template>

<style scoped>
</style>