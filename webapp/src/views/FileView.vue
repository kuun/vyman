<script setup>
import {DataGrid, Layout, LayoutPanel, Tree} from "v3-easyui";
</script>

<template>
  <Layout style="width:100%;height:100%;">
    <LayoutPanel region="west" style="width:240px;">
      <div class="title">
        <Tree :data="dirList" onSelectionChange="onSelectionChange"/>
      </div>
    </LayoutPanel>
    <LayoutPanel region="center" style="height:100%">
      <DataGrid :data="fileList" selection-mode="multiple" style="height: 100%">
        <GridColumn field="id" title="Item ID" :frozen="true"></GridColumn>
        <GridColumn field="name" title="名称" :frozen="true"></GridColumn>
        <GridColumn field="modifyTime" title="修改时间" align="right" width="100"></GridColumn>
        <GridColumn field="type" title="类型" width="200"></GridColumn>
        <GridColumn field="size" title="大小" align="right" width="100"></GridColumn>
        <GridColumn field="status" title="状态" width="100"></GridColumn>
      </DataGrid>
    </LayoutPanel>
  </Layout>

</template>

<style scoped>
</style>