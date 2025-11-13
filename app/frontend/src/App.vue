<!--
  - Copyright 2025 Bronya0 <tangssst@163.com>.
  - Author Github: https://github.com/Bronya0
  -
  - Licensed under the Apache License, Version 2.0 (the "License");
  - you may not use this file except in compliance with the License.
  - You may obtain a copy of the License at
  -
  -     https://www.apache.org/licenses/LICENSE-2.0
  -
  - Unless required by applicable law or agreed to in writing, software
  - distributed under the License is distributed on an "AS IS" BASIS,
  - WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  - See the License for the specific language governing permissions and
  - limitations under the License.
  -->

<template>
  <n-config-provider
      :date-locale="dateZhCN"
      :locale="zhCN" :theme="Theme" :hljs="hljs"
  >
    <!--https://www.naiveui.com/zh-CN/os-theme/components/layout-->
    <n-message-provider container-style="word-break: break-all;">
      <n-notification-provider container-style="text-align: left;" placement="bottom-right">
        <n-dialog-provider>
          <n-loading-bar-provider>
            <n-layout :class="headerClass" has-sider position="absolute" style="height: 100vh;">
              <!--header-->
              <n-layout-header bordered style="height: 42px; bottom: 0; padding: 0; ">
                <Header :desc="OneSay"/>
              </n-layout-header>
              <!--side + content-->
              <n-layout has-sider position="absolute" style="top: 42px; bottom: 0;">
                <n-layout-sider
                    :collapsed="true"
                    :collapsed-width="60"
                    bordered
                    collapse-mode="width"
                    style="--wails-draggable:drag"
                >
                  <n-menu
                      :collapsed-width="60"
                      :options="sideMenuOptions"
                      :value="route.path"
                      mode='vertical'
                      style="--wails-draggable:no-drag"
                      @update:value="handleMenuSelect"
                  />

                </n-layout-sider>
                <n-layout-content style="padding: 8px">
                  <router-view v-slot="{ Component }">
                    <keep-alive>
                      <component :is="Component" />
                    </keep-alive>
                  </router-view>

                </n-layout-content>
              </n-layout>
            </n-layout>
          </n-loading-bar-provider>
        </n-dialog-provider>
      </n-notification-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<script setup>
import {onMounted, ref, shallowRef} from 'vue'
import {
  darkTheme,
  dateZhCN,
  lightTheme,
  NConfigProvider,
  NLayout,
  NLayoutContent,
  NLayoutHeader,
  NMessageProvider,
  zhCN,
} from 'naive-ui'
import Header from './components/Header.vue'
import {useRoute, useRouter} from 'vue-router'
import {GetConfig, SaveConfig} from "../wailsjs/go/service/AppConfig";

import emitter from "./utils/eventBus";
import hljs from 'highlight.js/lib/core'
import powershell from 'highlight.js/lib/languages/powershell'
import {routes} from "./router";
import {getLocalLanguage} from "./utils/common";
import enUS from "./i18n/en-US";

hljs.registerLanguage('powershell', powershell)

let headerClass = shallowRef('lightTheme')
let Theme = shallowRef(lightTheme)

// 是否解锁的状态
let config = {}


onMounted(async () => {

  // 从后端加载配置
  const loadedConfig = await GetConfig()
  sessionStorage.setItem('app_config', JSON.stringify(loadedConfig));
  // 设置主题
  themeChange(loadedConfig.theme === darkTheme.name ? darkTheme : lightTheme)

  await Promise.allSettled(
      [
      ]
  )

  // =====================注册事件监听=====================
  // 主题切换
  emitter.on('update_theme', themeChange)
  // 菜单切换
  emitter.on('menu_select', handleMenuSelect)
  // 语言切换
  emitter.on('language_change', handleLanguageChange)


})

const route = useRoute()
const router = useRouter()

// 左侧菜单
const sideMenuOptions = generateMenuOptions(routes)

function generateMenuOptions(routes) {
  const options = []
  for (const route of routes) {
    // 如果设置了不显示在菜单中，则跳过
    if (!Object.hasOwn(route, 'meta')) {
      continue
    }

    const menuOption = {
      label: route.meta?.label || route.name, // 优先使用 meta.label
      key: route.meta?.key || route.path,     // 优先使用 meta.key
      icon: route.meta?.icon
    }

    // 如果有子路由并且也需要在菜单中显示，则递归处理
    if (route.children && route.children.length > 0) {
      const childrenOptions = generateMenuOptions(route.children)
      if (childrenOptions.length > 0) {
        menuOption.children = childrenOptions
      }
    }

    options.push(menuOption)
  }
  return options
}

// 切换菜单
function handleMenuSelect(key) {
  // console.log(key)
  router.push(key);
}


// 主题切换
function themeChange(newTheme) {
  Theme.value = newTheme
  headerClass = newTheme === lightTheme ? "lightTheme" : "darkTheme"
}

// naive ui language
const languageMap = {
  'zh-CN': zhCN,
  'en-US': enUS,
}

// 语言切换
// 未配置语言时，检查本地语言，合法则使用；否则使用默认语言
async function handleLanguageChange(language) {
  console.info("language配置：" + language);

  const DEFAULT_LANGUAGE = 'en-US';
  const localLanguage = getLocalLanguage();
  const supportedLanguages = Object.keys(languageMap);

  if (language === "" || language === null || language === undefined){
    language =  supportedLanguages.includes(localLanguage) ? localLanguage : DEFAULT_LANGUAGE
    config.language = language;
    console.info("未配置语言，使用本地语言并保存配置：" + language);
    await SaveConfig(config);
  }

  console.info("selectedLanguage：" + language);
  locale.value = language;
  naive_language.value = languageMap[language];
}


</script>

<style>
body {
  margin: 0;
  font-family: sans-serif;
}

.n-data-table .n-data-table__pagination {
  justify-content: start !important;
}

.lightTheme .n-layout-header {
  background-color: #f7f7fa;
}

.lightTheme .n-layout-sider {
  background-color: #f7f7fa !important;
}
</style>