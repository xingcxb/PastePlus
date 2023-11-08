<template>
  <a-form
      style="margin-top: 30px"
      :label-col="{ span: 10 }"
      :wrapper-col="{ span: 8 }"
      layout="horizontal"
      justify="end"
      labelAlign="right">
    <a-form-item label="启动">
      <a-checkbox v-model:checked="bootUp" class="formStyle">开机后启动 PastePlus</a-checkbox>
    </a-form-item>
    <a-form-item label="集成">
      <a-checkbox v-model:checked="pasteText" class="formStyle">粘贴为纯文本</a-checkbox>
    </a-form-item>
    <a-form-item label="其它">
      <a-checkbox style="cursor:default" v-model:checked="sound" disabled class="formStyle">启用音效</a-checkbox>
      <a-checkbox style="cursor: default" v-model:checked="menuIcon" disabled class="formStyle">在菜单栏显示图标</a-checkbox>
    </a-form-item>
    <a-form-item label="存储时长">
      <a-slider
          class="sliderStyle"
          :marks="marks"
          :step="null"
          :value="sliderValue"
          @change="handleChange"
          :tooltipOpen=false
          v-model:value="historyCapacity">
      </a-slider>
    </a-form-item>
  </a-form>
  <a-row>
    <a-col :offset="10">
      <a-button>清除所有记录</a-button>
    </a-col>
  </a-row>
  <a-row style="margin-top: 30px">
    <a-col :offset="4">赞助商：</a-col>
    <a-col>
      <a href="https://xingcxb.com?from=PastePlus">不器小窝</a>
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
// 滑块值
let sliderValue = ref(35);
// 滑块刻度
const marks = {
  0: '天',
  35: '周',
  70: '月',
  100: '无限'
}

// 限制滑块刻度选择
function handleChange(value) {
  sliderValue.value = parseInt(value);
}
</script>
<style scoped>
.formStyle {
  width: 100%
}
</style>