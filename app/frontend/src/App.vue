<template>
  <n-config-provider
      :theme="Theme"
  >
    <n-message-provider>
      <n-notification-provider placement="bottom-right">

        <n-layout has-sider position="absolute" style="height: 100vh;" :class="headerClass">

            <n-layout-header bordered style="height: 42px; bottom: 0; padding: 0;  --wails-draggable:drag">
              <Header
                  :value="activeTabLabel"
                  :options="menuOptions"
                  @update:value="handleMenuSelect"
                  @update_theme="themeChange"
              />
            </n-layout-header>

            <n-layout has-sider position="absolute" style="top: 42px; bottom: 0;">
              <n-layout-sider
                  bordered
                  collapsed
                  collapse-mode="width"
              >
              <Aside
                  :value="activeTabLabel"
                  @update:value="handleMenuSelect"
                  :options="sideMenuOptions"
              />

              </n-layout-sider>

              <n-layout-content>
                <n-tabs
                    type="card"
                    animated
                    size="small"
                    :value="activeTabLabel"
                    @update:value="handleTabChange"
                    @close="handleTabClose"
                >

                  <n-tab-pane v-for="tab in openTabs" :key="tab.key" :name="tab.label" display-directive="show"
                              :closable="openTabs.length > 1">
                    <!--                  <template #tab>-->
                    <!--                    <n-flex align="center">-->
                    <!--                      <n-icon :component="tab.tab_icon"/>-->
                    <!--                      {{ tab.label }}-->
                    <!--                    </n-flex>-->
                    <!--                  </template>-->

                    <div class="tab-content">
                      <component :is="tab.component" :name="tab.label" :key="tab.key"
                                 @update_theme="themeChange"
                                 @selectTab="handleTabChangeByIndex"
                      />
                    </div>
                  </n-tab-pane>
                </n-tabs>

              </n-layout-content>
            </n-layout>
          </n-layout>
      </n-notification-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
import {onMounted, ref, shallowRef} from 'vue'
import {
  darkTheme,
  lightTheme,
  NConfigProvider,
  NLayout,
  NLayoutContent,
  NLayoutHeader,
  NMessageProvider,
  NTabPane,
  NTabs
} from 'naive-ui'
import {HomeOutline, SettingsOutline,} from '@vicons/ionicons5'
import Header from './components/Header.vue'
import Settings from './components/Settings.vue'
import HelloWorld from './components/HelloWorld.vue'
import {GetConfig, SaveTheme} from "../wailsjs/go/config/AppConfig";
import {WindowSetSize} from "../wailsjs/runtime";
import {renderIcon} from "./utils/common";
import Aside from "./components/Aside.vue";

let headerClass = shallowRef('lightTheme')

onMounted(async () => {
  // 从后端加载配置
  const loadedConfig = await GetConfig()
  if (loadedConfig) {
    await WindowSetSize(loadedConfig.width, loadedConfig.height)
    if (loadedConfig.theme === 'light') {
      Theme.value = lightTheme
      headerClass = "lightTheme"
    } else {
      Theme.value = darkTheme
      headerClass = "darkTheme"
    }
  }
})
// 左侧菜单
const sideMenuOptions = [
  {
    label: '主页',
    key: '主页',
    icon: renderIcon(HomeOutline),
    component: HelloWorld,
    children: [
      {
        label: 'topic1',
        key: 'topic1',
        component: HelloWorld
      },
      {
        label: 'topic2',
        key: 'topic2',
        component: HelloWorld
      },
      {
        label: 'topic3',
        key: 'topic3',
        component: HelloWorld
      },
    ]
  },
  {
    label: '设置',
    key: '设置',
    icon: renderIcon(SettingsOutline),
    component: Settings
  },

]


// 顶部菜单
const menuOptions = []


const openTabs = shallowRef([sideMenuOptions[0].children[0]])
const activeTabLabel = ref(sideMenuOptions[0].children[0].label)

// 切换菜单
// item是从Header里handleSelect emits传过来的菜单对象
function handleMenuSelect(key, item) {
  // 检查 item 是否已存在于 openTabs 中
  const existingTab = findTabByItem(openTabs.value, item);

  if (!existingTab) {
    // 如果不存在，则添加到 openTabs 并设置为当前活动标签
    openTabs.value.push(item);
  }
  activeTabLabel.value = item.label;
  console.log(activeTabLabel.value)
}

// 递归查找子菜单中的项
function findTabByItem(tabs, item) {
  for (let tab of tabs) {
    if (tab === item) {
      return tab;
    }
    if (tab.children && tab.children.length > 0) {
      const found = findTabByItem(tab.children, item);
      if (found) {
        return found;
      }
    }
  }
  return null;
}

// 关闭tab
function handleTabClose(label) {
  const index = openTabs.value.findIndex(tab => tab.label === label);

  if (index !== -1) {
    // 创建一个新数组来触发响应式更新
    const newOpenTabs = [...openTabs.value]
    newOpenTabs.splice(index, 1)
    openTabs.value = newOpenTabs

    // 更新 activeTabLabel 为剩余标签中的最后一个
    if (activeTabLabel.value === label && newOpenTabs.length > 0) {
      activeTabLabel.value = newOpenTabs[newOpenTabs.length - 1].label
    }
  }
}

// 根据key切换
function handleTabChange(label) {
  console.log(label)
  activeTabLabel.value = label
}

// 根据index切换
function handleTabChangeByIndex(index) {
  // index转为int
  activeTabLabel.value = openTabs.value[parseInt(index)].label
}


let Theme = shallowRef(lightTheme)

// 主题切换
function themeChange(newTheme) {
  console.log(newTheme.name)
  Theme.value = newTheme
  headerClass = newTheme === lightTheme ? "lightTheme" : "darkTheme"
  SaveTheme(newTheme.name)
}

// // 自定义主题
// /**
//  * @type import('naive-ui').GlobalThemeOverrides
//  */
// const themeOverrides = {
//   Button: {
//     color: '#f2f2f7'
//   }
// }
</script>

<style>
body {
  margin: 0;
  font-family: sans-serif;
}

.tab-content {
  padding: 8px 24px;
}

.lightTheme .n-layout-header {
  background-color: #f2f2f7;
}

.right-section .n-button {
  margin-left: 10px;
}

.lightTheme .n-layout-sider {
  background-color: #f2f2f7 !important;
}
</style>