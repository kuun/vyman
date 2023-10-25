<script setup>
import {Layout, LayoutPanel} from "v3-easyui";
import {RouterView} from "vue-router";

</script>

<template>
  <div>
    <Layout style="width:100%;height:100%;">
      <LayoutPanel region="north" style="height:50px;">
        <div class="title">North Region</div>
      </LayoutPanel>
      <LayoutPanel region="center" style="height:100%">
        <RouterView/>
      </LayoutPanel>
    </Layout>
  </div>
</template>
