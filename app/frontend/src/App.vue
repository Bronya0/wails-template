<template>
  <n-config-provider
      :theme="Theme"
  >
    <!--https://www.naiveui.com/zh-CN/os-theme/components/layout-->
    <n-message-provider>
      <n-notification-provider placement="bottom-right">
        <n-layout has-sider position="absolute" style="height: 100vh;" :class="headerClass">
          <!--header-->
          <n-layout-header bordered style="height: 42px; bottom: 0; padding: 0; ">
            <Header
                :value="activeItem.label"
                :options="menuOptions"
                @update:value="handleMenuSelect"
                @update_theme="themeChange"
            />
          </n-layout-header>
          <!--side + content-->
          <n-layout has-sider position="absolute" style="top: 42px; bottom: 0;">
            <n-layout-sider
                bordered
                collapsed
                collapse-mode="width"
                :collapsed-width="60"
                style="--wails-draggable:drag"
            >
              <Aside
                  :collapsed-width="60"
                  :value="activeItem.label"
                  @update:value="handleMenuSelect"
                  :options="sideMenuOptions"
              />

            </n-layout-sider>
            <n-layout-content style="padding: 16px;">
              <keep-alive>
                <component :is="activeItem.component"></component>
              </keep-alive>

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
import emitter from "./utils/eventBus";

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

  emitter.on('update_theme', themeChange)
})
// 左侧菜单
const sideMenuOptions = [
  {
    label: '主页',
    key: '主页',
    icon: renderIcon(HomeOutline),
    tab_icon: HomeOutline,
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
    tab_icon: SettingsOutline,
    component: Settings
  },

]


// 顶部菜单
const menuOptions = []


const activeItem = shallowRef(sideMenuOptions[0].children[0])

// 切换菜单
function handleMenuSelect(key, item) {
  activeItem.value = item;
}

let Theme = shallowRef(lightTheme)

// 主题切换
function themeChange(newTheme) {
  console.log(newTheme.name)
  Theme.value = newTheme
  headerClass = newTheme === lightTheme ? "lightTheme" : "darkTheme"
  SaveTheme(newTheme.name)
}

// 自定义主题
// /**
//  * @type import('naive-ui').GlobalThemeOverrides
//  */
// const themeOverrides = {
//   common: {
//     bodyColor: '#FDFCFF'
//   }
// }

</script>

<style>
body {
  margin: 0;
  font-family: sans-serif;

}

.lightTheme .n-layout-header {
  background-color: #f7f7fa;
}

.lightTheme .n-layout-sider {
  background-color: #f7f7fa !important;
}
</style>