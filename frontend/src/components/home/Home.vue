<template>
  <a-row>
    <a-col :offset="22" :span="2" style="margin-bottom: 10px;background-color: aqua">
      11
    </a-col>
  </a-row>
  <a-space>
    <a-card class="cardInfo" title="app name" :bordered="false">
      <p>card content</p>
    </a-card>
    <a-card class="cardInfo" title="app name" :bordered="false">
      <p>card content</p>
    </a-card>
    <a-card class="cardInfo" title="app name" :bordered="false">
      <p>card content</p>
    </a-card>
  </a-space>
</template>
<script setup>

</script>
<style scoped>
.cardInfo{
  height: 200px;
  width: 256px;
}
</style>