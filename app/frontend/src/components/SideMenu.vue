<template>
  <n-menu
      :collapsed-width="64"
      :collapsed-icon-size="22"
      :mode="'vertical'"
      @update:value="handleSelect"
  />
</template>

<script setup>
import {h, ref} from 'vue'
import { NMenu } from 'naive-ui'

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