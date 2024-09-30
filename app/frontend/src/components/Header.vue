<template>
  <div class="top-menu">
    <div class="left-section">
      <n-avatar size="small" :src="logo"/>
      <h1 class="title">测试APP</h1>
    </div>
    <div class="right-section">
      <n-menu mode="horizontal" :options="props.options" @update:value="handleSelect"/>
      <n-button text @click="checkForUpdates">
        <template #icon>
          <n-icon>
            <refresh-outline/>
          </n-icon>
        </template>
        更新
      </n-button>
      <n-button text @click="changeTheme">
        <template #icon>
          <n-icon>
            <moon-outline v-if="Theme === lightTheme"/>
            <sunny-outline v-else/>
          </n-icon>
        </template>
        {{ Theme.name }}
      </n-button>

    </div>
  </div>
</template>

<script setup>
import {
  NMenu, NButton, NIcon, NAvatar, darkTheme,
  lightTheme, NSpace
} from 'naive-ui'
import {
  RefreshOutline, MoonOutline, SunnyOutline,
  HelpCircleOutline, LogOutOutline, SquareOutline, Close, LogoGithub, Remove, CopyOutline
} from '@vicons/ionicons5'
import logo from '../assets/images/logo.svg'


const props = defineProps({
  options: {
    type: Array,
    required: true
  }
});

const emit = defineEmits(['select', 'update_theme'])

let Theme = lightTheme


const handleSelect = (key, item) => {
  emit('select', item)
}

const checkForUpdates = () => {
  console.log('Checking for updates...')
}

// 主题切换
const changeTheme = () => {
  console.log(Theme.name)
  Theme = Theme === lightTheme ? darkTheme : lightTheme
  console.log(Theme.name)
  emit('update_theme', Theme)
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
  margin-left: 10px;
}
</style>