<template>
  <div class="top-menu">
    <div class="left-section">
      <n-avatar size="small" :src="logo"/>
      <h1 class="title">测试APP</h1>
      <n-h5>    {{version.tag_name}}</n-h5>
    </div>
    <div class="right-section">
      <n-menu mode="horizontal" :value="props.value" :options="props.options" @update:value="handleSelect"/>
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
import {NAvatar, NButton, NIcon, NMenu} from 'naive-ui'
import {RefreshOutline,} from '@vicons/ionicons5'
import logo from '../assets/images/logo.svg'
import {onMounted, ref, shallowRef} from "vue";
// import {ref} from "vue";
// import {Quit, WindowMaximise, WindowMinimise, WindowUnmaximise} from "../../wailsjs/runtime";
import {CheckUpdate} from '../../wailsjs/go/system/Update'
import {GetVersion} from '../../wailsjs/go/main/App'
import { useNotification } from 'naive-ui'

const props = defineProps({
  options: {},
  value: {}
});

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
  console.info("check version……")
  version.value.tag_name = await GetVersion()
  console.info(version.value.tag_name)
  const resp = await CheckUpdate()
  console.info(resp)
  if (resp.tag_name !== version.value.tag_name){
    notification.info({
      title: '发现新版本' + resp.tag_name,
      content: version.value.body,
    })
  }
}

onMounted(async () => {
  await checkForUpdates()

})

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