<template>
  <div class="top-menu">
    <div class="left-section">
      <n-avatar size="small" :src="logo"/>
      <h1 class="title">测试APP</h1>
      <n-h5> {{ version.tag_name }}</n-h5>
    </div>
    <div class="right-section">
      <n-menu mode="horizontal" :value="props.value" :options="props.options" @update:value="handleMenuSelect"/>

      <n-tooltip placement="bottom" trigger="hover">
        <template #trigger>
          <n-button quaternary :focusable="false" :loading="update_loading" @click="checkForUpdates"
                    :render-icon="renderIcon(RefreshOutline)"/>
        </template>
        <span> 检查版本 {{ version.tag_name }} </span>
      </n-tooltip>

      <n-button quaternary :focusable="false" @click="minimizeWindow" :render-icon="renderIcon(Remove)"/>
      <n-button quaternary :focusable="false" @click="resizeWindow" :render-icon="renderIcon(MaxMinIcon)"/>
      <n-button quaternary style="font-size: 22px" :focusable="false" @click="closeWindow">
        <n-icon><Close /></n-icon>
      </n-button>

    </div>
  </div>
</template>

<script setup>
import {NAvatar, NButton, NMenu, useMessage} from 'naive-ui'
import {RefreshOutline, SquareOutline, CopyOutline, Close, Remove} from '@vicons/ionicons5'
import logo from '../assets/images/logo.svg'
import {onMounted, ref, shallowRef} from "vue";
import {Quit, WindowMaximise, WindowMinimise, WindowUnmaximise} from "../../wailsjs/runtime";
import {CheckUpdate} from '../../wailsjs/go/system/Update'
import {GetVersion} from '../../wailsjs/go/main/App'
import {useNotification} from 'naive-ui'
import {renderIcon} from "../utils/common";

const props = defineProps(['options', 'value']);
const emit = defineEmits(['update:value'])


const update_loading = ref(false)

let version = ref({
  tag_name: "",
  body: "",
})

const notification = useNotification()
const message = useMessage()

const handleMenuSelect = (key, item) => {
  emit('update:value', key, item)
}

const checkForUpdates = async () => {
  update_loading.value = true
  try {
    const version = await GetVersion()
    const resp = await CheckUpdate()
    if (resp && resp.tag_name !== version) {
      notification.info({
        title: '发现新版本' + resp.tag_name,
        content: resp.body,
      })
    }
  } finally {
    update_loading.value = false
  }
}

onMounted(async () => {
  version.value.tag_name = await GetVersion()
  await checkForUpdates()

})

const isMaximized = ref(false);
const MaxMinIcon = shallowRef(SquareOutline)

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

</style>