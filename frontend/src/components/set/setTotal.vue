<template>
  <a-tabs v-model:activeKey="setType" centered>
    <a-tab-pane key="SettingsGeneral" tab="通用">
      <SetGeneral/>
    </a-tab-pane>
    <a-tab-pane key="SettingsShortcutKey" tab="快捷键">
      <SetShortcutKey/>
    </a-tab-pane>
    <a-tab-pane key="SettingsUpdate" tab="更新">
      <SetUpdate/>
    </a-tab-pane>
    <a-tab-pane key="SettingsAbout" tab="关于">
      <SetAbout/>
    </a-tab-pane>
  </a-tabs>
</template>

<script setup>
import {ref} from 'vue';
import router from "@/router/index.js";
import SetAbout from "@/components/set/setAbout.vue";
import SetUpdate from "@/components/set/setUpdate.vue";
import SetGeneral from "@/components/set/setGeneral.vue";
import SetShortcutKey from "@/components/set/setShortcutKey.vue";

const setType = ref('SettingsGeneral');
const activeRoute = ref('SettingsGeneral');

// 点击操作
function handleClick(routeName) {
  activeRoute.value = routeName;
  router.push({name: routeName});
}

</script>

<style scoped>
.headerIcon {
  font-size: 24px;
}

.headerFont {
  font-size: 14px;
  font-weight: bold;
}

.ant-checkbox-wrapper {
  text-align: left;
}
</style>