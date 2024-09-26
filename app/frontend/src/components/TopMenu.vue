<template>
  <div class="top-menu">
    <div class="left-section">
      <n-avatar size="small" :src="logo" />
      <h1 class="title">测试APP</h1>
    </div>
    <div class="right-section">
      <n-menu mode="horizontal" :options="menuOptions" @update:value="handleSelect" />
      <n-button text @click="checkForUpdates">
        <template #icon>
          <n-icon><refresh-outline /></n-icon>
        </template>
        Check for Updates
      </n-button>
      <n-button text @click="toggleTheme">
        <template #icon>
          <n-icon>
            <moon-outline v-if="theme === 'light'" />
            <sunny-outline v-else />
          </n-icon>
        </template>
        {{ theme === 'light' ? 'Dark' : 'Light' }} Mode
      </n-button>
    </div>
  </div>
</template>

<script setup>
import { h, ref } from 'vue'
import { NMenu, NButton, NIcon, NAvatar } from 'naive-ui'
import { RefreshOutline, MoonOutline, SunnyOutline,
  SettingsOutline, HelpCircleOutline, LogOutOutline } from '@vicons/ionicons5'
import logo from '../assets/images/logo.svg'

const emit = defineEmits(['select', 'update_theme'])

const theme = ref('light')

const menuOptions = [
  {
    label: 'File',
    key: 'file',
    children: [
      {
        label: 'New',
        key: 'new'
      },
      {
        label: 'Open',
        key: 'open'
      },
      {
        label: 'Save',
        key: 'save'
      }
    ]
  },
  {
    label: 'Edit',
    key: 'edit',
    children: [
      {
        label: 'Undo',
        key: 'undo'
      },
      {
        label: 'Redo',
        key: 'redo'
      },
      {
        label: 'Cut',
        key: 'cut'
      },
      {
        label: 'Copy',
        key: 'copy'
      },
      {
        label: 'Paste',
        key: 'paste'
      }
    ]
  },
  {
    label: 'View',
    key: 'view',
    children: [
      {
        label: 'Zoom In',
        key: 'zoomIn'
      },
      {
        label: 'Zoom Out',
        key: 'zoomOut'
      },
      {
        label: 'Reset Zoom',
        key: 'resetZoom'
      }
    ]
  },
  {
    label: () =>
        h(
            'a',
            {
              href: 'https://github.com/yourusername/yourproject',
              target: '_blank',
              rel: 'noopenner noreferrer'
            },
            'GitHub'
        ),
    key: 'github'
  },
  {
    label: 'More',
    key: 'more',
    children: [
      {
        label: 'Settings',
        key: 'settings',
        icon: () => h(NIcon, null, {default: () => h(SettingsOutline)})
      },
      {
        label: 'Help',
        key: 'help',
        icon: () => h(NIcon, null, {default: () => h(HelpCircleOutline)})
      },
      {
        label: 'Logout',
        key: 'logout',
        icon: () => h(NIcon, null, {default: () => h(LogOutOutline)})
      }
    ]
  }
]

const handleSelect = (key) => {
  emit('select', key)
}

const checkForUpdates = () => {
  // Implement your update checking logic here
  console.log('Checking for updates...')
}

const toggleTheme = () => {
  theme.value = theme.value === 'light' ? 'dark' : 'light'
  emit('update_theme', theme.value)
}
</script>

<style scoped>
.top-menu {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
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