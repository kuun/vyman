<script setup>
import {Layout, LayoutPanel} from "v3-easyui";
import ObjectGroup from "../components/ObjectGroup.vue";
import GroupItemList from "../components/GroupItemList.vue";
</script>

<template>
  <Layout style="width:100%;height:100%;">
    <LayoutPanel region="west" style="width:50%;height:100%">
      <ObjectGroup group-type="ipv6-address"></ObjectGroup>
    </LayoutPanel>
    <LayoutPanel region="center" style="height:100%;width: 100%">
      <GroupItemList group-type="ipv6-address" item-type="address"></GroupItemList>
    </LayoutPanel>
  </Layout>
</template>