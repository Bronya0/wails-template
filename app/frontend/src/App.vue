<template>
  <n-config-provider :theme="Theme">
    <n-message-provider>
      <n-layout has-sider>
        <n-layout>

            <n-notification-provider placement="bottom-right">
              <n-layout-header :class="headerClass" bordered
                               style="height: 6dvh; padding: 0 20px 0 0; --wails-draggable:drag">

              <Header :value="activeTab" :options="menuOptions" @select="handleMenuSelect"/>
              </n-layout-header>
            </n-notification-provider>



          <n-layout has-sider style="height: 94dvh;">
            <n-layout-sider
                bordered
                :collapsed="collapsed"
                collapse-mode="width"
                :collapsed-width="64"
                :width="120"
                show-trigger
                @collapse="collapsed = true"
                @expand="collapsed = false"
            >
              <n-menu
                  :collapsed-width="64"
                  :collapsed-icon-size="22"
                  :mode="'vertical'"
                  :value="activeTab"
                  @select="handleMenuSelect"
                  :options="sideMenuOptions"
              />
            </n-layout-sider>

            <n-layout-content>
              <n-tabs
                  type="card"
                  animated
                  size="small"
                  :value="activeTab"
                  @update:value="handleTabChange"
                  @close="handleTabClose"
              >

                <n-tab-pane v-for="tab in openTabs" :key="tab.key" :name="tab.key" display-directive="show"
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
      </n-layout>
    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
import {computed, h, onMounted, ref, shallowRef} from 'vue'
import {
  darkTheme,
  lightTheme,
  NConfigProvider,
  NIcon,
  NLayout,
  NLayoutContent,
  NLayoutHeader, NMenu,
  NMessageProvider,
  NTabPane,
  NTabs, useOsTheme
} from 'naive-ui'
import {HomeOutline, SettingsOutline,} from '@vicons/ionicons5'
import Header from './components/Header.vue'
import Settings from './components/Settings.vue'
import HelloWorld from './components/HelloWorld.vue'
import {GetConfig} from "../wailsjs/go/config/AppConfig";
import {WindowSetSize} from "../wailsjs/runtime";
import {renderIcon} from "./utils/common";

let headerClass = shallowRef('lightTheme')
const collapsed = shallowRef(true)

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

// 顶部菜单
const menuOptions = [
  {
    label: '首页',
    key: 'HelloWorld',
    icon: renderIcon(HomeOutline),
    tab_icon: HomeOutline,
    children: [
      {
        label: 'HelloWorld',
        key: 'HelloWorld',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld1',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld2',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld3',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld4',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld5',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld6',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld7',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld8',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld9',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
      {
        label: 'HelloWorld',
        key: 'HelloWorld10',
        icon: renderIcon(HomeOutline),
        component: HelloWorld
      },
    ]
  },
  {
    label: '设置',
    key: 'Settings',
    icon: renderIcon(SettingsOutline),
    children: [
      {
        label: 'Settings',
        key: 'Settings',
        icon: renderIcon(SettingsOutline),
        component: Settings
      },
    ]
  },
]

// 左侧菜单
const sideMenuOptions = [
  {
    label: '主页',
    key: 'HelloWorld3',
    icon: renderIcon(HomeOutline),
    component: HelloWorld
  },
  {
    label: '设置',
    key: 'Settings4',
    icon: renderIcon(SettingsOutline),
    component: Settings
  },
]

const openTabs = shallowRef([menuOptions[0].children[0]])
const activeTab = ref(menuOptions[0].children[0].key)

// 切换菜单
// item是从Header里handleSelect emits传过来的菜单对象
function handleMenuSelect(key, item) {
  // 检查 item 是否已存在于 openTabs 中
  const existingTab = findTabByKey(openTabs.value, item.key);

  if (!existingTab) {
    // 如果不存在，则添加到 openTabs 并设置为当前活动标签
    openTabs.value.push(item);
  }
  activeTab.value = key;
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
    // 创建一个新数组来触发响应式更新
    const newOpenTabs = [...openTabs.value]
    newOpenTabs.splice(index, 1)
    openTabs.value = newOpenTabs

    // 更新 activeTab 为剩余标签中的最后一个
    if (activeTab.value === key && newOpenTabs.length > 0) {
      activeTab.value = newOpenTabs[newOpenTabs.length - 1].key
    }
  }
}


// 根据key切换
function handleTabChange(key) {
  activeTab.value = key
}

// 根据index切换
function handleTabChangeByIndex(index) {
  // index转为int
  activeTab.value = openTabs.value[parseInt(index)].key
}


let Theme = shallowRef(lightTheme)

// 主题切换
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

.right-section .n-button {
  margin-left: 10px;
}
</style>