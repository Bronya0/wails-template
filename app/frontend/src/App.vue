<template>
  <n-config-provider :theme=Theme>
    <n-message-provider>

      <n-layout has-sider>

        <n-layout>

          <n-layout-header bordered style="height: 64px; padding: 10px 20px; background-color: #f2f2f2">
            <TopMenu @update_theme="themeChange" @select="handleMenuSelect"/>
          </n-layout-header>

          <n-layout has-sider>
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
                  type="card"
                  animated
                  :value="activeTab"
                  @update:value="handleTabChange"
                  pane-style="padding-left: 4px; padding-right: 4px; box-sizing: border-box;"
              >

                <n-tab-pane v-for="tab in tabs" :key="tab.name" :name="tab.name" :tab="tab.label">
                  <template #tab>
                    <n-space align="center">
                      <n-icon :component="tab.icon"/>
                      {{ tab.label }}
                    </n-space>
                  </template>

                  <div class="tab-content">
                    <component :is="tab.component"/>
                  </div>>

                </n-tab-pane>

              </n-tabs>


            </n-layout-content>
          </n-layout>

          <Footer/>

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
} from 'naive-ui'
import {HomeOutline, SettingsOutline} from '@vicons/ionicons5'
import SideMenu from './components/SideMenu.vue'
import TopMenu from './components/TopMenu.vue'
import Footer from './components/Footer.vue'
import Settings from './components/Settings.vue'

const activeTab = ref('settings')
const collapsed = ref(false)
const Theme = ref("light")

// 左侧竖排菜单
const menuOptions = [

  {
    label: '首页',
    key: 'shouye',
    icon: () => h(HomeOutline),
    tabs: [
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
const handleMenuSelect = (key) => {
  activeTab.value = key
}

// 点击tab切换
const handleTabChange = (name) => {
  activeTab.value = name
}
const themeChange = (theme) => {
  Theme.value = theme
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