<script setup>
import {Panel} from "v3-easyui";
</script>

<template>
  <Panel :border="false">
    <template v-slot:header>
      <div>
        <img class="product-title" src="../assets/img/netguard.png">
      </div>
    </template>
  </Panel>
</template>

<style scoped>
.product-title {
  height: 45px;
}
</style>
