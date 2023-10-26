<template>
  <a-row style="padding-top: 5px;padding-bottom: 5px" type="flex" justify="center">
    <a-col :offset="10" :span="4" >
        <a-input-search
            placeholder="输入搜索关键字"
            v-model="searchKeyword"
            class="search-input"
            key="searchInput"
        >
        </a-input-search>
    </a-col>
    <a-col :offset="9" :span="1">
      <!--   下拉菜单   -->
      <a-dropdown :trigger="['click']">
        <a-button type="text">
          <template #icon>
            <Icon icon="material-symbols:more-horiz" style="font-size: 24px"/>
          </template>
        </a-button>
        <template #overlay>
          <a-menu style="background-color: #f3f5f2">
            <a-menu-item key="1">偏好设置...&nbsp;&nbsp;</a-menu-item>
            <a-menu-item key="2">帮助中心</a-menu-item>
            <a-menu-item key="3">退出</a-menu-item>
          </a-menu>
        </template>
      </a-dropdown>
    </a-col>
  </a-row>
  <a-space style="padding-left: 17px;width: 100%" size="middle">
    <div v-for="paste in pasteList.value">
      <a-card class="cardInfo" :title=paste.from_app :bordered="false">
        <a-typography>
          <a-typography-paragraph :ellipsis=true :content=paste.content>
          </a-typography-paragraph>
        </a-typography>
      </a-card>
    </div>
  </a-space>
</template>
<script setup>
import {Icon} from "@iconify/vue";
import {nextTick, onBeforeUnmount, onMounted, reactive, ref} from "vue";
// 输入框的值
const searchKeyword = ref("");

// 查询数据
const onSearch = val => {
  console.log('search:', val);
};
// 查询值
let searchKey = ref("");
// 粘贴历史列表
let pasteList = reactive({value: {}});

// 加载粘贴历史数据
function loadHistoryPasteData() {
  wails.Events.Emit({name: "findPasteHistoryToCore", Data: ""})
  wails.Events.Once("findPasteHistoryToFrontend", function (data) {
    let pasteJson = JSON.parse(data.data)
    console.log("======>",pasteJson)
    pasteList.value = pasteJson
  })
}

onMounted(() => {
  loadHistoryPasteData()
});
onBeforeUnmount(() => {
});
</script>
<style scoped>
.cardInfo {
  height: 200px;
  width: 256px;
}

.search-input {
  width: 200px;
}

.search-icon {
  cursor: pointer;
}

</style>
