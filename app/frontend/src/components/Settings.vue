<template>
  <div>
    <n-form :model="config" label-placement="left">
      <n-form-item label="窗口宽度">
        <n-input-number v-model:value="config.width" :min="800" :max="1920" :style="{ maxWidth: '120px' }"/>
      </n-form-item>
      <n-form-item label="窗口高度">
        <n-input-number v-model:value="config.height" :min="600" :max="1080" :style="{ maxWidth: '120px' }"/>
      </n-form-item>
      <n-form-item label="语言">
        <n-select v-model:value="config.language" :options="languageOptions" :style="{ maxWidth: '120px' }"/>
      </n-form-item>

      <n-form-item label="主题">
        <n-space>
          <n-button strong  @click="theme=darkTheme">
            dark
          </n-button>
          <n-button strong  type="tertiary" @click="theme=lightTheme">
            light
          </n-button>

        </n-space>
      </n-form-item>


      <n-form-item>
        <n-button @click="saveConfig" strong type="primary">保存设置</n-button>
      </n-form-item>


    </n-form>
  </div>
</template>

<script setup>
import {onActivated, onBeforeUnmount, onMounted, ref} from 'vue'
import {darkTheme, lightTheme, NButton, NForm, NFormItem, NInputNumber, NSelect, useMessage} from 'naive-ui'
import {GetConfig, SaveConfig} from '../../wailsjs/go/config/AppConfig'
import {WindowSetSize} from "../../wailsjs/runtime";
import {onDeactivated} from "@vue/runtime-core";

const message = useMessage()

const emit = defineEmits(['update_theme'])
let theme = lightTheme

const config = ref({
  width: 1024,
  height: 768,
  language: 'zh-CN',
  theme: theme.name,
})

const languageOptions = [
  {label: '中文', value: 'zh-CN'},
  {label: 'English', value: 'en-US'}
]

onMounted(async () => {
  console.info("初始化settings……")

  // 从后端加载配置
  const loadedConfig = await GetConfig()
  console.log(loadedConfig)
  if (loadedConfig) {
    config.value = loadedConfig
    await WindowSetSize(loadedConfig.width, loadedConfig.height)

    theme = loadedConfig.theme === 'dark' ? darkTheme : lightTheme
    emit('update_theme', theme)
  }
})

const saveConfig = async () => {
  config.value.theme = theme.name
  const err = await SaveConfig(config.value)
  if (err !== ""){
    message.error("保存失败：" + err)
    return
  }

  await WindowSetSize(config.value.width, config.value.height)

  emit('update_theme', theme)
  // 可以添加保存成功的提示
  message.success("保存成功")
  config.value = await GetConfig()

}

//
// onBeforeUnmount(() => {
//   console.info("卸载settings组件……")
// })
//
// onActivated(() => {
//   console.info("激活settings组件……")
// })
// onDeactivated(() => {
//   console.info("失活settings组件……")
// })

</script>