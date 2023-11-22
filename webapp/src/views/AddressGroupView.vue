<script setup>
import {Layout, LayoutPanel} from "v3-easyui";
import AddressGroup from "../components/AddressGroup.vue";
import AddressList from "../components/AddressList.vue";
</script>

<template>
    <Layout style="width:100%;height:100%;">
      <LayoutPanel region="west" style="width:50%;height:100%">
        <AddressGroup group-type="address"></AddressGroup>
      </LayoutPanel>
      <LayoutPanel region="center" style="height:100%;width: 100%">
        <AddressList group-type="address"></AddressList>
      </LayoutPanel>
    </Layout>
</template>