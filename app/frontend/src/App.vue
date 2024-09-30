<template>
  <n-config-provider :theme="Theme">
    <n-message-provider>
      <n-layout has-sider>
        <n-layout>
          <n-layout-header :class="headerClass" bordered
                           style="height: 8dvh; padding: 10px 20px; --wails-draggable:drag">
            <Header @update_theme="themeChange" @select="handleMenuSelect"/>
          </n-layout-header>

          <n-layout has-sider style="height: 92dvh;">
            <n-layout-sider
                bordered
                collapse-mode="width"
                :collapsed-width="64"
                :width="140"
                :collapsed="collapsed"
                show-trigger
                @collapse="collapsed = true"
                @expand="collapsed = false"
            >
              <SideMenu @select="handleMenuSelect" :options="menuOptions"/>
            </n-layout-sider>
            <n-layout-content>
              <n-tabs
                  type="card"
                  animated
                  :value="activeTab"
                  @update:value="handleTabChange"
              >
                <n-tab-pane v-for="tab in activeMenu.tabs" :key="tab.key" :name="tab.key" display-directive="show"
                            :tab="tab.name">
                  <template #tab>
                    <n-space align="center">
                      <n-icon :component="tab.icon"/>
                      {{ tab.name }}
                    </n-space>
                  </template>

                  <div class="tab-content">
                      <component :is="tab.component" :name="tab.name" :key="tab.key" @update_theme="themeChange"/>
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
  lightTheme,
  NConfigProvider,
  NLayout,
  NLayoutContent,
  NLayoutHeader,
  NLayoutSider,
  NMessageProvider,
  NTabPane,
  NTabs
} from 'naive-ui'
import {HomeOutline, SettingsOutline} from '@vicons/ionicons5'
import SideMenu from './components/SideMenu.vue'
import Header from './components/Header.vue'
import Settings from './components/Settings.vue'
import HelloWorld from './components/HelloWorld.vue'

const collapsed = ref(true)
let Theme = ref(lightTheme)
let headerClass = ref('lightTheme')

// 左侧竖排菜单
const menuOptions = [

  {
    label: '首页',
    key: 'HelloWorld',
    icon: () => h(HomeOutline),
    tabs: [
      {name: 'HelloWorld', key: 'HelloWorld', icon: HomeOutline, component: HelloWorld},
    ]
  },
  {
    label: '设置',
    key: 'Settings',
    icon: () => h(SettingsOutline),
    tabs: [
      {name: 'Settings', key: 'Settings', icon: SettingsOutline, component: Settings},
    ]
  },
]
// 选中的菜单，是json对象；activeTab是key
const activeMenu = shallowRef(menuOptions[0])
const activeTab = shallowRef(menuOptions[0].tabs[0].key)

// 点击左侧菜单，切换activeTab的key
function handleMenuSelect(key) {
  const newActiveMenu = menuOptions.find(item => item.key === key)
  if (newActiveMenu) {
    activeMenu.value = newActiveMenu
    if (newActiveMenu.tabs && newActiveMenu.tabs.length > 0) {
      activeTab.value = newActiveMenu.tabs[0].key
    } else {
      activeTab.value = null
    }
  }
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
</style>