<template>
  <n-menu :options="menuOptions" @update:value="handleSelect" />
</template>

<script setup>
import { h } from 'vue'
import { NMenu } from 'naive-ui'
import { BookOutline, SettingsOutline } from '@vicons/ionicons5'

const menuOptions = [
  {
    label: 'Docs',
    key: 'docs',
    icon: () => h(BookOutline)
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