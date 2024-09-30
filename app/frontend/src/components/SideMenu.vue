<template>
  <n-menu
      :collapsed-width="64"
      :collapsed-icon-size="22"
      :mode="'vertical'"
      @update:value="handleSelect"
  />
</template>

<script setup>
import { NMenu } from 'naive-ui'
import {ref} from "vue";

const emit = defineEmits(['select'])
const collapsed = ref(false)

const handleSelect = (key) => {
  emit('select', key)
}

</script>

<style scoped>
.side-menu {
  height: 100%;
}
</style>