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

const handleSelect = (key, item) => {
  emit('select', key, item)
}

</script>

<style scoped>

</style>