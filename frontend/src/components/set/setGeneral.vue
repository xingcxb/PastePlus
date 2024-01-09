<template>
  <a-form
      style="margin-top: 30px"
      :label-col="{ span: 10 }"
      :wrapper-col="{ span: 8 }"
      layout="horizontal"
      justify="end"
      labelAlign="right">
    <a-form-item label="启动">
      <a-checkbox v-model="bootUp" @change="handleBootUp" class="formStyle">开机后启动 PastePlus</a-checkbox>
    </a-form-item>
    <a-form-item label="集成">
      <a-checkbox v-model="pasteText" class="formStyle">粘贴为纯文本</a-checkbox>
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
          @change="handleSliderChange"
          :tooltipOpen=false
          v-model:value="historyCapacity">
      </a-slider>
    </a-form-item>
  </a-form>
  <a-row>
    <a-col :offset="10">
      <a-button @click="cleanAllPasteHistory">清除所有记录</a-button>
    </a-col>
  </a-row>
  <a-row style="margin-top: 30px">
    <a-col :offset="4">友联：</a-col>
    <a-col>
      <a href="https://xingcxb.com?from=PastePlus">不器小窝</a>
    </a-col>
  </a-row>
</template>
<script setup>
import {onMounted, ref} from 'vue';

// 是否开机启动
let bootUp = ref(false);
// 是否粘贴为纯文本
let pasteText = ref(false);
// 是否启动音效
let sound = ref(false);
// 是否在菜单栏显示图标
let menuIcon = ref(false);
// 历史记录容量
let historyCapacity = ref("周");
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
function handleSliderChange(value) {
  sliderValue.value = parseInt(value);
  switch (value){
    case 0:
      historyCapacity.value = "天";
      break;
    case 35:
      historyCapacity.value = "周";
      break;
    case 70:
      historyCapacity.value = "月";
      break;
    case 100:
      historyCapacity.value = "无限";
      break;
  }
}

// 操作开机启动
function handleBootUp(value) {
  wails.Events.Emit({name: "handleBootUpToCore", Data: value})
  wails.Events.On("handleBootUpToFrontend", (data) => {
    bootUp.value = data.data
  })
}

// 清除所有记录
function cleanAllPasteHistory(){
  wails.Events.Emit({name: "cleanAllPasteHistoryToCore"})
}

// 加载配置
function loadPasteConfig(){
  wails.Events.Emit({name: "loadPasteConfigToCore"})
  wails.Events.On("loadPasteConfigToFrontend", (data) => {
    bootUp.value = data.data.bootUp
    sound.value = data.data.sound
    historyCapacity.value = data.data.historyCapacity
    switch (data.data.historyCapacity){
      case "天":
        sliderValue.value = 0;
        break;
      case "周":
        sliderValue.value = 35;
        break;
      case "月":
        sliderValue.value = 70;
        break;
      case "无限":
        sliderValue.value = 100;
        break;
    }
  })
}

// 页面打开时监听数据
onMounted(()=>{
  loadPasteConfig()
})

</script>
<style scoped>
.formStyle {
  width: 100%
}
</style>