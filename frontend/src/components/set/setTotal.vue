<template>
  <!-- 设置 -->
  <!--  <a-row justify="center">-->
  <!--    <a-col :span="2">-->
  <!--      <a-button-->
  <!--        type="primary"-->
  <!--        @click="handleClick('SettingsGeneral')"-->
  <!--      >-->
  <!--      <Icon icon="mingcute:switch-line" class="headerIcon"/>-->
  <!--      <p class="headerFont">通用</p>-->
  <!--      </a-button>-->
  <!--    </a-col>-->
  <!--    <a-col :span="2">-->
  <!--      <Icon icon="material-symbols:keyboard-command-key" class="headerIcon"/>-->
  <!--      <p class="headerFont">快捷键</p>-->
  <!--    </a-col>-->
  <!--    <a-col :span="2">-->
  <!--      <Icon icon="iconamoon:synchronize-fill" class="headerIcon"/>-->
  <!--      <p class="headerFont">同步</p>-->
  <!--    </a-col>-->
  <!--    <a-col :span="2">-->
  <!--      <Icon icon="mdi:information-outline" class="headerIcon"/>-->
  <!--      <p class="headerFont">关于</p>-->
  <!--    </a-col>-->
  <!--  </a-row>-->
  <!--  <router-view />-->
  <a-tabs v-model:activeKey="setType" tab-bar-style="card" centered>
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