<script setup>
import {Layout, LayoutPanel, DataList} from "v3-easyui";
import {onMounted, ref} from "vue";
import axios from "axios";

const ifaceList = ref([])
const selectedIfaceName = ref('')

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
        ifaceList.value = resp.data
      })
      .catch((resp) => {
        alert('获取网卡列表失败.')
      })
})

const onSelectionChange = event => {
  console.log('selection change: ', event)
  selectedIfaceName.value = event.data.name
};

</script>

<template>
  <Layout style="width:100%;height:100%;">
    <LayoutPanel region="west" style="width:150px; height: 100%">
      <Panel title="网卡列表" style="height:100%; width: 100%">
        <DataList :data="ifaceList" selectionMode="single" @onSelectionChange="onSelectionChange">
          <template v-slot="scope">
            <div class="ifaceListItem">
              {{scope.row.name}}
            </div>
          </template>
        </DataList>
      </Panel>
    </LayoutPanel>
    <LayoutPanel region="center" style="height:100%">
      <Tabs style="height:100%">
        <TabPanel :title="'IP管理'">
          <IpManager :ifaceName="selectedIfaceName"></IpManager>
        </TabPanel>
        <TabPanel :title="'网卡设置'">
          <p>Tab Panel2</p>
        </TabPanel>
      </Tabs>
    </LayoutPanel>
  </Layout>

</template>

<style scoped>
.ifaceListItem {
  font-size: 20px;
  padding: 5px;
}
</style>