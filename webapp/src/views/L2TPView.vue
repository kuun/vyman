<script setup>

import {Layout, LayoutPanel, TabPanel, Tabs} from "v3-easyui";
import L2TPManager from "@/components/L2TPManager.vue";
</script>

<template>
  <Layout style="width:100%;height:100%;">
    <LayoutPanel region="center" style="height:100%">
      <Tabs style="height:100%" :border="false">
        <TabPanel :title="'L2TP配置'" :border="false">
           <L2TPManager/>
        </TabPanel>
          <TabPanel :title="'隧道信息'">
            <p>Tab Panel2</p>
          </TabPanel>
      </Tabs>
    </LayoutPanel>
  </Layout>
</template>

<style scoped>

</style>