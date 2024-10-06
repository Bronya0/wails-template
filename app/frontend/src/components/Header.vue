<template>
  <div class="top-menu">
    <div class="left-section">
      <n-avatar size="small" :src="logo"/>
      <h1 class="title">测试APP</h1>
      <n-h5> {{ version.tag_name }}</n-h5>
    </div>
    <div class="right-section">
      <n-menu mode="horizontal" :value="props.value" :options="props.options" @update:value="handleSelect"/>

      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <n-button quaternary :focusable="false" :loading="loading" @click="checkForUpdates" :render-icon="renderIcon(RefreshOutline)"/>
        </template>
        <span> 检查版本 {{ version.tag_name}} </span>
      </n-tooltip>

      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <n-button quaternary :focusable="false" @click="minimizeWindow" :render-icon="renderIcon(Remove)"/>
        </template>
        <span> 最小化 </span>
      </n-tooltip>

      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <n-button quaternary :focusable="false" @click="resizeWindow" :render-icon="renderIcon(MaxMinIcon)"/>
        </template>
        <span> {{ isMaximized ? '还原' : '最大化'}} </span>
      </n-tooltip>

      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <n-button quaternary :focusable="false" @click="closeWindow" :render-icon="renderIcon(Close)"
                    style="margin-right: 10px;"/>
        </template>
        <span> 关闭 </span>
      </n-tooltip>

    </div>
  </div>
</template>

<script setup>
import {NAvatar, NButton, NIcon, NMenu} from 'naive-ui'
import {RefreshOutline, SquareOutline, CopyOutline, Close, Remove} from '@vicons/ionicons5'
import logo from '../assets/images/logo.svg'
import {onMounted, ref} from "vue";
import {Quit, WindowMaximise, WindowMinimise, WindowUnmaximise} from "../../wailsjs/runtime";
import {CheckUpdate} from '../../wailsjs/go/system/Update'
import {GetVersion} from '../../wailsjs/go/main/App'
import {useNotification} from 'naive-ui'
import {renderIcon} from "../utils/common";

const props = defineProps({
  options: {},
  value: {}
});

const loading = ref(false)


const emit = defineEmits(['select'])
let version = ref({
  tag_name: "",
  body: "",
})

const notification = useNotification()

const handleSelect = (key, item) => {
  emit('select', key, item)
}

const checkForUpdates = async () => {
  loading.value = true
  try {
    console.info("check version……")
    version.value.tag_name = await GetVersion()
    console.info(version.value.tag_name)
    const resp = await CheckUpdate()
    console.info(resp)
    if (resp.tag_name !== version.value.tag_name) {
      notification.info({
        title: '发现新版本' + resp.tag_name,
        content: version.value.body,
      })
    }
  } finally {
    loading.value = false
  }

}

onMounted(async () => {
  await checkForUpdates()

})

const isMaximized = ref(false);
const MaxMinIcon = ref(SquareOutline)

const minimizeWindow = () => {
  WindowMinimise()
}

const resizeWindow = () => {
  isMaximized.value = !isMaximized.value;
  // 最大化
  if (isMaximized.value) {
    WindowMaximise();
    MaxMinIcon.value = CopyOutline;
  } else {
    WindowUnmaximise();
    MaxMinIcon.value = SquareOutline;
  }
  console.log(isMaximized.value)
}

const closeWindow = () => {
  Quit()
}

</script>

<style scoped>

.top-menu {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 0 0 20px;
  height: 100%;
}

.left-section {
  display: flex;
  align-items: center;
}

.title {
  margin-left: 10px;
  font-size: 1.2em;
}

.right-section {
  display: flex;
  align-items: center;
}

.right-section .n-button {
  padding: 0 8px;
}
</style>