<template>
  <n-config-provider :theme="Theme">
    <n-message-provider>
      <n-layout has-sider>
        <n-layout>
          <n-layout-header :class="headerClass" bordered
                           style="height: 8dvh; padding: 10px 20px; --wails-draggable:drag">

            <Header :options="menuOptions" @update_theme="themeChange" @select="handleMenuSelect"/>

          </n-layout-header>

          <n-layout has-sider style="height: 92dvh;">
            <n-layout-sider
                bordered
                collapse-mode="width"
                :collapsed-width="64"
                :width="140"
                show-trigger
            >
              <SideMenu @select="handleMenuSelect" :options="sideMenuOptions"/>

            </n-layout-sider>

            <n-layout-content>
              <n-tabs
                  type="card"
                  animated
                  :value="activeTab"
                  @update:value="handleTabChange"
                  @close="handleTabClose"
              >
                <n-tab-pane v-for="tab in openTabs" :key="tab.key" :name="tab.key" display-directive="show"
                            :closable="openTabs.length > 1">
                  <template #tab>
                    <n-space align="center">
                      <n-icon :component="tab.icon"/>
                      {{ tab.label }}
                    </n-space>
                  </template>

                  <div class="tab-content">
                    <component :is="tab.component" :name="tab.label" :key="tab.key" @update_theme="themeChange"/>
                  </div>
                </n-tab-pane>
              </n-tabs>

            </n-layout-content>
          </n-layout>
        </n-layout>
      </n-layout>
    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
import {h, ref, shallowRef} from 'vue'
import {
  darkTheme,
  lightTheme,
  NConfigProvider, NIcon,
  NLayout,
  NLayoutContent,
  NLayoutHeader, NMenu,
  NMessageProvider,
  NTabPane,
  NTabs
} from 'naive-ui'
import {HomeOutline, SettingsOutline,RefreshOutline, MoonOutline, SunnyOutline,} from '@vicons/ionicons5'
import Header from './components/Header.vue'
import Settings from './components/Settings.vue'
import HelloWorld from './components/HelloWorld.vue'
import SideMenu from "./components/SideMenu.vue";

let Theme = shallowRef(lightTheme)
let headerClass = shallowRef('lightTheme')


// 菜单
const menuOptions = [
  {
    label: '首页',
    key: 'HelloWorld',
    icon: renderIcon(HomeOutline),
    children: [
      {label: 'HelloWorld', key: 'HelloWorld',  icon: HomeOutline,component: HelloWorld},
    ]
  },
  {
    label: '设置',
    key: 'Settings',
    icon: renderIcon(SettingsOutline),
    children: [
      {label: 'Settings', key: 'Settings', icon: SettingsOutline, component: Settings},
    ]
  },
]
const sideMenuOptions = [
  {
    label: '主页',
    key: 'HelloWorld3',
    icon: renderIcon(HomeOutline),
    component: Settings
  },
  {
    label: '设置',
    key: 'Settings4',
    icon: renderIcon(SettingsOutline),
    component: Settings
  },
]
const openTabs = shallowRef([menuOptions[0].children[0]])
const activeTab = shallowRef(menuOptions[0].children[0].key)

// 切换菜单
// item是从Header里handleSelect emits传过来的菜单对象
function handleMenuSelect(item) {
  // 检查 item 是否已存在于 openTabs 中
  const existingTab = findTabByKey(openTabs.value, item.key);

  if (existingTab) {
    // 如果存在，则更新 activeTab
    activeTab.value = item.key;
  } else {
    // 如果不存在，则添加到 openTabs 并设置为当前活动标签
    openTabs.value.push(item);
    activeTab.value = item.key;
  }
}

// 递归查找子菜单中的项
function findTabByKey(tabs, key) {
  for (let tab of tabs) {
    if (tab.key === key) {
      return tab;
    }
    if (tab.children && tab.children.length > 0) {
      const found = findTabByKey(tab.children, key);
      if (found) {
        return found;
      }
    }
  }
  return null;
}

// 关闭tab
function handleTabClose(key) {
  const index = openTabs.value.findIndex(tab => tab.key === key);

  if (index !== -1) {
    // 从 openTabs 中删除该项
    openTabs.value.splice(index, 1);

    // 更新 activeTab 为剩余标签中的最后一个
    if (activeTab.value === key && openTabs.value.length > 0) {
      activeTab.value = openTabs.value[openTabs.value.length - 1].key;
    }
  }
}

function renderIcon(icon) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

// 点击tab切换
function handleTabChange(key) {
  activeTab.value = key
}

function themeChange(newTheme) {
  console.log(newTheme.name)
  Theme.value = newTheme
  headerClass = newTheme === lightTheme ? "lightTheme" : "darkTheme"
}


</script>

<style>
body {
  margin: 0;
  font-family: sans-serif;
}

.tab-content {
  padding: 8px 24px;
}

.lightTheme {
  background-color: #f2f2f2;
}

.top-menu {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 0 0 20px;
  height: 100%;
}

.left-section {
  display: flex;
  align-items: center;
}

.title {
  margin-left: 10px;
  font-size: 1.2em;
}

.right-section {
  display: flex;
  align-items: center;
}

.right-section .n-button {
  margin-left: 10px;
}
</style>