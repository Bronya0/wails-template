<template>
  <n-config-provider :theme="theme">
    <n-layout has-sider>

      <n-layout>

        <n-layout-header bordered style="height: 64px; padding: 10px 20px; background-color: #f6f8fa">
          <TopMenu @select="handleMenuSelect" />
        </n-layout-header>

        <n-layout has-sider>
          <n-layout-sider
              bordered
              collapse-mode="width"
              :collapsed-width="64"
              :width="200"
              :collapsed="collapsed"
              show-trigger
              @collapse="collapsed = true"
              @expand="collapsed = false"
          >
            <SideMenu @select="handleMenuSelect" />
          </n-layout-sider>


          <n-layout-content>
            <n-tabs type="card" animated v-model:value="activeTab" @close="handleClose" @add="handleAdd">
              <n-tab-pane v-for="tab in tabs" :key="tab.name" :name="tab.name" :tab="tab.title" >
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
import { useOsTheme } from 'naive-ui'
import SideMenu from './components/SideMenu.vue'
import TopMenu from './components/TopMenu.vue'
import Footer from './components/Footer.vue'

const osTheme = useOsTheme()
const theme = ref(osTheme.value)

const collapsed = ref(false)
const activeTab = ref('tab1')
const tabs = reactive([
  { name: 'tab1', title: 'Tab 1', content: 'Content of Tab 1' },
  { name: 'tab2', title: 'Tab 2', content: 'Content of Tab 2' }
])

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