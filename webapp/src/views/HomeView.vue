<script setup>
import {Layout, LayoutPanel} from "v3-easyui";
import {RouterView} from "vue-router";
import Header from "../components/Header.vue";
import Sidebar from "../components/Sidebar.vue";

</script>

<template>
  <div>
    <Layout style="width:100%;height:100%;">
      <LayoutPanel region="north" style="height:60px;">
        <Header></Header>
      </LayoutPanel>
      <LayoutPanel region="west" style="height:100%;width: 250px">
        <Sidebar/>
      </LayoutPanel>
      <LayoutPanel region="south" style="height:25px">
      </LayoutPanel>
      <LayoutPanel region="center" style="height:100%;width: 100%">
        <RouterView/>
      </LayoutPanel>
    </Layout>
  </div>
</template>
