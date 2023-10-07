<template>
  <a-row>
    <a-col :offset="10" :span="8">
      <a-form justify="end" labelAlign="right">
        <a-form-item label="启动">
          <a-checkbox v-model:checked="bootUp">开机后启动 Paste+</a-checkbox>
        </a-form-item>
        <a-form-item label="集成">
          <a-checkbox v-model:checked="pasteText">粘贴为纯文本</a-checkbox>
        </a-form-item>
        <a-form-item label="其它">
          <a-checkbox v-model:checked="sound" disabled>启用音效</a-checkbox>
          <a-checkbox v-model:checked="menuIcon" disabled>在菜单栏显示图标</a-checkbox>
        </a-form-item>
        <a-form-item label="历史记录容量">
          <a-slider v-model:value="historyCapacity" :marks="marks" dots/>
        </a-form-item>
      </a-form>
    </a-col>
  </a-row>
</template>
<script setup>
import {ref} from 'vue';

// 是否开机启动
const bootUp = ref(false);
// 是否粘贴为纯文本
const pasteText = ref(false);
// 是否启动音效
const sound = ref(false);
// 是否在菜单栏显示图标
const menuIcon = ref(false);
// 历史记录容量
const historyCapacity = ref("周");
const marks = {
  0: '天',
  35: '周',
  70: '月',
  100: '无限',
}
</script>
<style scoped>

</style>