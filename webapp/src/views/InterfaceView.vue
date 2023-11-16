<script setup>
import {DataList, Layout, LayoutPanel, Panel, TabPanel, Tabs} from "v3-easyui";
import {onMounted, ref} from "vue";
import axios from "axios";
import IpManager from "../components/IpManager.vue";
import {useIfaceStore} from "../stores/iface";

const ifaceList = ref([])
const ifaceStore = useIfaceStore()

const props = defineProps({
  ifaceType: {
    type: String,
    required: true
  }
})

onMounted(() => {
  axios.get('/api/interfaces/' + props.ifaceType)
      .then((resp) => {
        ifaceList.value = resp.data
      })
      .catch((resp) => {
        alert('获取网卡列表失败.')
      })
})

const onSelectionChange = event => {
  ifaceStore.setSelectedIface(event)
};

</script>

<template>
  <Layout style="width:100%;height:100%;">
    <LayoutPanel region="west" style="width:150px; height: 100%">
      <Panel title="网卡列表" style="height:100%; width: 100%" :border="false">
        <DataList :data="ifaceList" selectionMode="single" @selectionChange="onSelectionChange"
                  :border="false">
          <template v-slot="scope">
            <div class="ifaceListItem">
              {{scope.row.name}}
            </div>
          </template>
        </DataList>
      </Panel>
    </LayoutPanel>
    <LayoutPanel region="center" style="height:100%">
      <Tabs style="height:100%" :border="false">
        <TabPanel :title="'IP管理'" :border="false">
          <IpManager/>
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