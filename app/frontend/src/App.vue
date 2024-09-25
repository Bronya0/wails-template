<template>
  <n-config-provider :theme="theme">
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

            <n-tabs type="line" justify-content="space-evenly" animated v-model:value="activeTab">
              <n-tab-pane v-for="tab in tabs" :key="tab.name" :name="tab.name">
                <template #tab>
                  <n-space align="center">
                    <n-icon :component="tab.icon" />
                    {{ tab.title }}
                  </n-space>
                </template>
                {{ tab.content }}
              </n-tab-pane>
            </n-tabs>

          </n-layout-content>
        </n-layout>

        <Footer />

      </n-layout>

    </n-layout>
  </n-config-provider>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { NConfigProvider, NLayout, NLayoutSider, NLayoutHeader, NLayoutContent, NTabs, NTabPane } from 'naive-ui'
import { HomeOutline, PersonOutline, SettingsOutline } from '@vicons/ionicons5'

import { useOsTheme } from 'naive-ui'
import SideMenu from './components/SideMenu.vue'
import TopMenu from './components/TopMenu.vue'
import Footer from './components/Footer.vue'

const osTheme = useOsTheme()
const theme = ref(osTheme.value)

const collapsed = ref(false)
const activeTab = ref('tab1')
const tabs = [
  {
    name: 'home',
    title: 'Home',
    content: 'Home content',
    icon: HomeOutline
  },
  {
    name: 'profile',
    title: 'Profile',
    content: 'Profile content',
    icon: PersonOutline
  },
  {
    name: 'settings',
    title: 'Settings',
    content: 'Settings content',
    icon: SettingsOutline
  }
]

const handleMenuSelect = (key) => {
  // Handle menu item selection
  console.log('Selected menu item:', key)
}

const handleClose = (name) => {
  const index = tabs.findIndex(tab => tab.name === name)
  if (index !== -1) {
    tabs.splice(index, 1)
  }
}


const handleAdd = () => {
  const newTab = {
    name: `tab${tabs.length + 1}`,
    title: `New Tab ${tabs.length + 1}`,
    content: `Content of New Tab ${tabs.length + 1}`
  }
  tabs.push(newTab)
  activeTab.value = newTab.name
}
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
</style>