<template>
  <a-row style="padding-top: 10px;position: fixed;top: 0;width: 100%" type="flex" justify="center">
    <a-col :offset="10" :span="4">
      <a-input-search
          size="small"
          placeholder="输入搜索关键字"
          v-model="searchKeyword"
          class="search-input"
          key="searchInput"
      >
      </a-input-search>
    </a-col>
    <a-col :offset="9" :span="1">
      <!--   下拉菜单   -->
      <a-dropdown :trigger="['click']" placement="bottomRight">
        <a-button type="text">
          <template #icon>
            <Icon icon="material-symbols:more-horiz" style="font-size: 24px"/>
          </template>
        </a-button>
        <template #overlay>
          <!--   下拉菜单的菜单   -->
          <a-menu style="background-color: #f3f5f2">
            <a-menu-item key="1">偏好设置...&nbsp;&nbsp;</a-menu-item>
            <a-menu-item key="2">帮助中心</a-menu-item>
            <a-menu-item key="3">退出</a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </a-col>
  </a-row>
  <a-space style="padding-left: 17px;margin-top: 43px;width: 100%;height: 100%;overflow-y: hidden;" size="middle">
    <div v-for="paste in pasteList.value">
      <a-card hoverable
              class="cardInfo"
              @click="handleCardClick(paste.id)"
              @dblclick.native="handleCardDoubleClick(paste.id)"
              :title=paste.spacing_time
              :bordered="false">
        <!--    图片类型渲染   -->
        <a-image
            v-if="paste.type === 'image'"
            :src="paste.content"
            :preview="false"
            :previewMask="false"
        />
        <!--    文本类型渲染   -->
        <a-typography-paragraph
            v-else
            style="text-align: left"
            :ellipsis="{rows:5,expandable:false}"
            :content=paste.content>
        </a-typography-paragraph>
      </a-card>
    </div>
  </a-space>
</template>
<script setup>
import {Icon} from "@iconify/vue";
import {nextTick, onBeforeUnmount, onMounted, reactive, ref} from "vue";
// 粘贴历史列表
let pasteList = reactive({value: {}});
// 输入框的值
const searchKeyword = ref("");
// 用于处理单击双击事件的区分
let timeRecords = null;

// 查询数据
const onSearch = val => {
  console.log('search:', val);
};

let isDoubleClick = false;
// 卡片单击操作，将数据写入粘贴板中
function handleCardClick(pasteDataId) {
  clearTimeout(timeRecords);
  timeRecords = setTimeout(() => {
    if (isDoubleClick) {
      isDoubleClick = false;
      return;
    }
    console.log("单击来咯")
    wails.Events.Emit({name: "handleCardClickToCore", Data: pasteDataId})
  },300);
}

// 卡片双击操作，将数据写入粘贴板，并执行粘贴动作
function handleCardDoubleClick(pasteDataId) {
  clearTimeout(timeRecords);  //清除
  console.log("双击来咯")
  wails.Events.Emit({name: "handleCardDoubleClickToCore", Data: pasteDataId})
}


// 加载粘贴历史数据
function loadHistoryPasteData() {
  wails.Events.Emit({name: "findPasteHistoryToCore", Data: ""})
  wails.Events.Once("findPasteHistoryToFrontend", function (data) {
    let pasteJson = JSON.parse(data.data)
    pasteList.value = pasteJson
  })
//   let pasteJson = [
//   {
//     "id": 3,
//     "from_app": "",
//     "content": "HandleCardClick",
//     "type": "text",
//     "spacing_time": "现在",
//     "created_at": "2023-11-13 17:36:49"
//   },
//   {
//     "id": 2,
//     "from_app": "",
//     "content": "单击卡片绑定事件",
//     "type": "text",
//     "spacing_time": "现在",
//     "created_at": "2023-11-13 17:36:46"
//   }
// ];
//   pasteList.value = pasteJson
}

// 启动加载数据
onMounted(() => {
  loadHistoryPasteData()
});
</script>
<style scoped>
.cardInfo {
  height: 200px;
  width: 256px;
  user-select: none;
  cursor: default;
}

.search-input {
  width: 100%;
}

.search-icon {
  cursor: pointer;
}

::-webkit-scrollbar {
  width: 0 !important;
}

::-webkit-scrollbar {
  width: 0 !important;
  height: 0;
}

</style>
