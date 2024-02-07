<template>
  <a-row>
    <a-card style="padding-top: 30px;width: 620px;height: 300px" :body-style="{padding:0,overflow:'hidden'}">
      <a-flex justify="space-between">
        <a-image
            alt="avatar"
            :preview="false"
            src="../logo.png"
            style="display: block;width:273px"
        />
        <a-flex vertical align="flex-end" justify="space-between" :style="{ padding: '32px' }">
          <a-typography>
            <a-typography-title :level="3">
              开发中...
            </a-typography-title>
          </a-typography>
          <a-button type="primary" onclick="openBrowser('https://xingcxb.com')" target="_blank">去催催</a-button>
        </a-flex>
      </a-flex>
    </a-card>
  </a-row>
</template>
<script setup>
// 浏览器打开链接
function openBrowser(site) {
  window.open(site);
}
</script>
<style scoped>

</style>