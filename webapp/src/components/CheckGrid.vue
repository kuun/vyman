<script setup>
import {computed, ref} from "vue";

const props = defineProps({
  data: {
    type: Array,
    required: true
  }
})

// 实现checkbox选中
const allChecked = ref(false);
const rowClicked = ref(false);
const checkedRows = computed(() => {
  return props.data.value.filter(row => row.selected);
});
const onAllCheckedChange = (checked) => {
  if (rowClicked.value) {
    return;
  }
  ipAddrList.value.map(row => {
    row.selected = checked;
  });
};

const onCheckedChange = (checked) => {
  allChecked.value = checkedRows.value.length === ipAddrList.value.length;
  rowClicked.value = true;
  nextTick(() => (rowClicked.value = false));
};
</script>


<template>
  <DataGrid :data="props.data">
    <GridColumn field="selected" :width="30" align="center">
      <template v-slot:header="scope">
        <CheckBox v-model="allChecked" @checkedChange="onAllCheckedChange($event)"></CheckBox>
      </template>
      <template v-slot:body="scope" @checkedChange="onCheckedChange($event)">
        <CheckBox v-model="scope.row.selected" @checkedChange="onCheckedChange($event)"></CheckBox>
      </template>
    </GridColumn>
    <slot></slot>
  </DataGrid>
</template>