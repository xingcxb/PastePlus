<template>
  <a-segmented class="segmentedStyle" :selectedIndex="setType"
               size="large" v-model:value="setType" block
               :options="setTopData" @change="onChange">
    <template #label="{ value: val, payload = {} }">
      <div style="padding: 4px 4px;text-align:center">
        <template v-if="payload.icon">
          <Icon :icon="payload.icon" class="iconStyle">
          </Icon>
          <div style="line-height: 10px">{{ payload.label }}</div>
        </template>
      </div>
    </template>
  </a-segmented>
  <div v-if="setType === 'SettingsGeneral'">
    <SetGeneral/>
  </div>
  <div v-else-if="setType === 'SettingsShortcutKey'">
    <SetShortcutKey/>
  </div>
  <div v-else-if="setType === 'SettingsUpdate'">
    <SetUpdate/>
  </div>
  <div v-else>
    <SetAbout/>
  </div>
</template>

<script setup>
import {ref} from 'vue';
import SetGeneral from "@/components/set/setGeneral.vue";
import {Icon} from "@iconify/vue";
import SetShortcutKey from "@/components/set/setShortcutKey.vue";
import SetUpdate from "@/components/set/setUpdate.vue";
import SetAbout from "@/components/set/setAbout.vue";

// 顶部导航栏数据
const setTopData = ref([
  {
    value: 'SettingsGeneral',
    payload: {
      label: '通用',
      icon: "material-symbols:settings-outline",
    },
  },
  {
    value: 'SettingsShortcutKey',
    disabled: true,
    payload: {
      label: '快捷键',
      icon: "material-symbols:keyboard-alt-outline",
    },
  },
  {
    value: 'SettingsUpdate',
    disabled: true,
    payload: {
      label: '更新',
      icon: "material-symbols:update",
    },
  },
  {
    value: 'SettingsAbout',
    payload: {
      label: '关于',
      icon: "mdi:information-outline",
    },
  },
]);
// 选中的导航栏
const setType = ref('SettingsGeneral');

// 改变分段控制器操作
function onChange(key) {
  setType.value = key;
}

</script>

<style scoped>
.segmentedStyle {
  margin: 5px;
  height: 70px;
}

.iconStyle {
  font-size: 36px;
  color: #00c042;
  margin-bottom: -5px;
}
</style>