<template>
  <a-row>
    <a-col :offset="8" :span="8">
      <a-image style="border-radius: 20px;user-select: none;" :width="80" src="../logo.png" :preview="false"/>
      <a-typography-title
          class="AppInfo"
          :level="3"
      >Paste+
      </a-typography-title>
      <a-divider />
      <a-typography-paragraph class="appVersion">
        Paste+ 是一款跨平台的剪贴板管理工具，支持 Windows、MacOS、Linux。
      </a-typography-paragraph>
      <p class="appVersion">版本 0.0.1</p>
      <p class="appVersion">开发者：<span onclick="openBrowser('https://xingcxb.com')" class="link">Symbol</span></p>
      <p class="appVersion">Github：<span onclick="openBrowser('https://xingcxb.com')" class="link">Symbol</span></p>
    </a-col>
  </a-row>
</template>
<script setup>

// 浏览器打开链接
function openBrowser(site) {
  window.open(site);
}

</script>
<style scoped>
.AppInfo{
  font-weight: bold;
  margin-top: 10px;
  user-select: none;
}
.appVersion{
  font-size: 13px;
  user-select: none;
  padding-bottom: 10px;
}
.link{
  color: #5562f9;
  user-select: none;
  cursor: pointer;
}
</style>