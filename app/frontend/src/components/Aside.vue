<template>
    <n-menu
        :mode="'vertical'"
        :value="props.value"
        @update:value="handleMenuSelect"
        :options="props.options"
    />
</template>


<script setup>
import {useMessage, useNotification} from "naive-ui";

const props = defineProps(['options', 'value']);

const emit = defineEmits(['update:value'])
const message = useMessage()
const notification = useNotification()

const handleMenuSelect = (key, item) => {
  emit('update:value', key, item)
}

</script>

<style>
.aside-btn .n-button .n-button__icon {
  font-size: 22px;
}
</style>