<script setup>

import {ref} from "vue";

const ipAddrList = ref([])
const props = defineProps({
  ifaceName: {
    type: String,
    required: true
  }
})


onMounted(() => {
  axios.get('/api/interfaces/' + props.ifaceName + '/ip')
      .then((resp) => {
        console.log('ip list: ', resp.data)
        ipAddrList.value = resp.data
      })
      .catch((resp) => {
        alert('获取IP列表失败.')
      })
})
</script>

<template>
  <DataGrid :data="ipAddrList" selection-mode="multiple" style="height: 100%">
    <GridColumn field="ip" title="IP"></GridColumn>
    <GridColumn field="mask" title="掩码" width="100"></GridColumn>
  </DataGrid>
</template>

<style scoped>
.product-title {
  height: 43px;
}
</style>
