<template>
    <n-menu
        :mode="'vertical'"
        :value="props.value"
        @update:value="handleMenuSelect"
        :options="props.options"
    />

<!--  <n-flex class="aside-btn">-->
<!--    <n-button quaternary :render-icon="renderIcon(LogoGithub)"/>-->
<!--  </n-flex>-->
</template>


<script setup>
import {lightTheme, NButton, useMessage, useNotification} from "naive-ui";
import {renderIcon} from "../utils/common";
import {LogoGithub, SunnyOutline} from "@vicons/ionicons5";

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