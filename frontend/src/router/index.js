import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  {
    // 设置页面
    path: "/settings",
    name: "Settings",
    component: () => import("../components/set/setTotal.vue"),
  },
  {
    // 设置通用页面
    path: "/settings/general",
    name: "SettingsGeneral",
    component: () => import("../components/set/setGeneral.vue"),
  },
  {
    // 设置快捷键页面
    path: "/settings/shortcut",
    name: "SettingsShortcut",
    component: () => import("../components/set/setShortcutKey.vue"),
  },
  {
    // 设置更新页面
    path: "/settings/update",
    name: "SettingsUpdate",
    component: () => import("../components/set/setUpdate.vue"),
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;