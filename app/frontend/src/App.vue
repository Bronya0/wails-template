<template>
  <n-config-provider :theme="theme">
    <n-message-provider>

    <n-layout has-sider>

      <n-layout>

        <n-layout-header bordered style="height: 64px; padding: 10px 20px; background-color: #f2f2f2">
          <TopMenu @select="handleMenuSelect" />
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
            <SideMenu @select="handleMenuSelect" />
          </n-layout-sider>


          <n-layout-content>
            <n-tabs type="card" animated :value="activeTab" @update:value="handleTabChange">
              <n-tab-pane v-for="tab in tabs" :key="tab.name" :name="tab.name" :tab="tab.label">
                <template #tab>
                  <n-space align="center">
                    <n-icon :component="tab.icon" />
                    {{ tab.label }}
                  </n-space>
                </template>
                <component :is="tab.component" />
              </n-tab-pane>
            </n-tabs>



          </n-layout-content>
        </n-layout>

        <Footer />

      </n-layout>

    </n-layout>

    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
import {ref, reactive, shallowRef, onMounted} from 'vue'
import { NConfigProvider, NLayout, NLayoutSider, NLayoutHeader, NLayoutContent, NTabs, NTabPane, NMessageProvider} from 'naive-ui'
import { HomeOutline, PersonOutline, SettingsOutline } from '@vicons/ionicons5'
import { useOsTheme } from 'naive-ui'
import SideMenu from './components/SideMenu.vue'
import TopMenu from './components/TopMenu.vue'
import Footer from './components/Footer.vue'
import Settings from './components/Settings.vue'

const activeTab = ref('settings')

const osTheme = useOsTheme()
const theme = ref(osTheme.value)

const collapsed = ref(false)

const tabs = [
  { name: 'settings', label: '界面设置', icon: SettingsOutline, component: Settings},
]


const handleMenuSelect = (key) => {
  activeTab.value = key
}

const handleTabChange = (name) => {
  activeTab.value = name
}

// onMounted(async () => {
//   // 在应用启动时读取配置
//   await window.go.main.App.LoadConfig()
// })

</script>

<style>
body {
  margin: 0;
  font-family: sans-serif;
}

html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  overflow: hidden;
}

.n-layout-sider {
  height: 100vh;
}
</style>