<template>
  <n-config-provider :theme="Theme">
    <n-message-provider>
      <n-layout has-sider>
        <n-layout>
          <n-layout-header bordered style="height: 8vh; padding: 10px 20px; ">
            <TopMenu @update_theme="themeChange" @select="handleMenuSelect"/>
          </n-layout-header>


          <n-layout has-sider style="height: 82vh;">
            <n-layout-sider
                bordered
                collapse-mode="width"
                :collapsed-width="64"
                :width="138"
                :collapsed="collapsed"
                show-trigger
                @collapse="collapsed = true"
                @expand="collapsed = false"
            >
              <SideMenu @select="handleMenuSelect" :options="menuOptions"/>
            </n-layout-sider>

            <n-layout-content>
              <n-tabs
                  v-if="activeMenu && activeMenu.tabs && activeMenu.tabs.length > 0"
                  type="card"
                  animated
                  :value="activeTab"
                  @update:value="handleTabChange"
                  pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;"
              >
                <n-tab-pane v-for="tab in activeMenu.tabs" :key="tab.key" :name="tab.key" :tab="tab.name">
                  <template #tab>
                    <n-space align="center">
                      <n-icon :component="tab.icon"/>
                      {{ tab.name }}
                    </n-space>
                  </template>

                  <div class="tab-content">
                    <component :is="tab.component"/>
                  </div>
                </n-tab-pane>
              </n-tabs>

            </n-layout-content>
          </n-layout>

          <Footer style="height: 10vh;"></Footer>
        </n-layout>
      </n-layout>
    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
import {computed, h, ref} from 'vue'
import {
  NConfigProvider,
  NLayout,
  NLayoutContent,
  NLayoutHeader,
  NLayoutSider,
  NMessageProvider,
  NTabPane,
  NTabs,
  darkTheme,
  lightTheme
} from 'naive-ui'
import {HomeOutline, SettingsOutline} from '@vicons/ionicons5'
import SideMenu from './components/SideMenu.vue'
import TopMenu from './components/TopMenu.vue'
import Footer from './components/Footer.vue'
import Settings from './components/Settings.vue'
import HelloWorld from './components/HelloWorld.vue'

const activeTab = ref('settings')
const activeMenu = ref(null)
const collapsed = ref(false)
const Theme = ref(lightTheme)

// 左侧竖排菜单
const menuOptions = [

  {
    label: '首页',
    key: 'shouye',
    icon: () => h(HomeOutline),
    tabs: [
      {name: 'hello world',key: 'hello', icon: HomeOutline, component: HelloWorld},
    ]
  },
  {
    label: '设置',
    key: 'settings',
    icon: () => h(SettingsOutline),
    tabs: [
      {name: '界面设置',key: 'settings1', icon: SettingsOutline, component: Settings},
      {name: '界面设置2', key: 'settings2', icon: SettingsOutline, component: Settings},
    ]
  },
]

// 点击左侧菜单，切换tab
function handleMenuSelect(key) {
  activeMenu.value = menuOptions.find(item => item.key === key)
  if (activeMenu.value && activeMenu.value.tabs && activeMenu.value.tabs.length > 0) {
    activeTab.value = activeMenu.value.tabs[0].key
  } else {
    activeTab.value = null
  }
}
// 点击tab切换
function handleTabChange(key) {
  activeTab.value = key
}
function themeChange(newTheme) {
  console.log(newTheme.value)
  Theme.value = newTheme
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
</style>