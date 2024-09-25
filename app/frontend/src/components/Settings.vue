<template>
  <div>
    <n-form :model="config" label-placement="left" label-width="auto" require-mark-placement="right-hanging" :style="{ maxWidth: '640px' }">
      <n-form-item label="窗口宽度">
        <n-input-number v-model:value="config.width" :min="800" :max="1920" />
      </n-form-item>
      <n-form-item label="窗口高度">
        <n-input-number v-model:value="config.height" :min="600" :max="1080" />
      </n-form-item>
      <n-form-item label="语言">
        <n-select v-model:value="config.language" :options="languageOptions" />
      </n-form-item>
      <n-form-item>
        <n-button @click="saveConfig" type="primary">保存设置</n-button>
      </n-form-item>
    </n-form>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { NForm, NFormItem, NInputNumber, NSelect, NButton} from 'naive-ui'
import { useMessage } from 'naive-ui'

const message = useMessage()


const config = ref({
  width: 1024,
  height: 768,
  language: 'zh-CN'
})

const languageOptions = [
  {label: '中文', value: 'zh-CN'},
  {label: 'English', value: 'en-US'}
]
// onMounted(async () => {
//   // 从后端加载配置
//   const loadedConfig = await LoadConfig()
//   if (loadedConfig) {
//     config.value = loadedConfig
//   }
// })

const saveConfig = async () => {
  // 可以添加保存成功的提示
  message.success("保存成功")
}
</script>