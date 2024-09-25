<template>
  <n-menu :options="menuOptions" @update:value="handleSelect" />
</template>

<script setup>
import { h } from 'vue'
import { NMenu } from 'naive-ui'
import { BookOutline, SettingsOutline } from '@vicons/ionicons5'

const menuOptions = [
  {
    label: '集群',
    key: 'docs',
    icon: () => h(BookOutline)
  },
  {
    label: '节点',
    key: '节点',
    icon: () => h(SettingsOutline)
  },
  {
    label: '设置',
    key: 'settings',
    icon: () => h(SettingsOutline)
  },
  {
    label: '存储',
    key: '存储',
    icon: () => h(SettingsOutline)
  },
  {
    label: '授权',
    key: '授权',
    icon: () => h(SettingsOutline)
  },
  {
    label: '认证',
    key: '认证',
    icon: () => h(SettingsOutline)
  },
  {
    label: 'Settings',
    key: 'settings',
    icon: () => h(SettingsOutline)
  },
  {
    label: 'Settings',
    key: 'settings',
    icon: () => h(SettingsOutline)
  },
  {
    label: 'Settings',
    key: 'settings',
    icon: () => h(SettingsOutline)
  }
]

const emit = defineEmits(['select'])

const handleSelect = (key) => {
  emit('select', key)
}
</script>

<style scoped>
.side-menu {
  height: 100%;
}
</style>