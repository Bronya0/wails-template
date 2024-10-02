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

<!--      <n-button quaternary  :focusable="false" @click="minimizeWindow" :render-icon="renderIcon(Remove)" style="margin-left: 8px"/>-->
<!--      <n-button quaternary  :focusable="false" @click="resizeWindow">-->
<!--        <template #icon>-->
<!--          <n-icon size="16px">-->
<!--            <SquareOutline v-if="!isMaximized"/>-->
<!--            <CopyOutline v-else/>-->
<!--          </n-icon>-->
<!--        </template>-->
<!--      </n-button>-->
<!--      <n-button quaternary  :focusable="false" @click="closeWindow" :render-icon="renderIcon(Close)" />-->

    </div>
  </div>
</template>

<script setup>
import {
  NMenu, NButton, NIcon, NAvatar
} from 'naive-ui'
import {
  RefreshOutline,
} from '@vicons/ionicons5'
import logo from '../assets/images/logo.svg'
// import {ref} from "vue";
// import {Quit, WindowMaximise, WindowMinimise, WindowUnmaximise} from "../../wailsjs/runtime";


const props = defineProps({
  options: {
    type: Array,
    required: true
  }
});

const emit = defineEmits(['select'])


const handleSelect = (key, item) => {
  emit('select', key, item)
}

const checkForUpdates = () => {
  console.log('Checking for updates...')
}


// const isMaximized = ref(false);
//
// const minimizeWindow = () => {
//   WindowMinimise()
// }
//
// const resizeWindow = () => {
//   isMaximized.value = !isMaximized.value;
//   if (isMaximized.value) {
//     WindowMaximise();
//   } else {
//     WindowUnmaximise();
//   }
// }
//
// const closeWindow = () => {
//   Quit()
// }

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