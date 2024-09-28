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

      <n-form-item>
        <n-space>
          <n-button @click="theme.value='dark'">
            深色
          </n-button>
          <n-button @click="theme.value='light'">
            浅色
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
import {onMounted, ref} from 'vue'
import {NButton, NForm, NFormItem, NInputNumber, NSelect, useMessage} from 'naive-ui'
import {GetConfig, SaveConfig} from '../../wailsjs/go/config/AppConfig'
import {WindowSetSize} from "../../wailsjs/runtime";

const message = useMessage()
const props = defineProps({
  theme: {
    type: String,
    default: "light"
  }
})
const emit = defineEmits(['update_theme'])
const theme = ref(props.theme)

const config = ref({
  width: 1024,
  height: 768,
  language: 'zh-CN',
  theme: "light",
})

const languageOptions = [
  {label: '中文', value: 'zh-CN'},
  {label: 'English', value: 'en-US'}
]

onMounted(async () => {
  // 从后端加载配置
  const loadedConfig = await GetConfig()
  console.log(loadedConfig)
  if (loadedConfig) {
    config.value = loadedConfig
    await WindowSetSize(loadedConfig.width, loadedConfig.height)
    emit('update_theme', loadedConfig.theme)
  }
})

const saveConfig = async () => {
  const err = await SaveConfig(config.value)
  if (err !== ""){
    message.error("保存失败：" + err)
    return
  }

  await WindowSetSize(config.value.width, config.value.height)

  emit('update_theme', config.value.theme)
  // 可以添加保存成功的提示
  message.success("保存成功")
}
</script>