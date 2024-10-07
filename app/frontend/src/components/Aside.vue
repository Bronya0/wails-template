<template>
  <div class="aside-btn">
    <n-flex vertical>
      <n-menu
          :mode="'vertical'"
          :value="value"
          @update:value="handleMenuSelect"
          :options="options"
      />
<!--      <n-button text ghost :focusable="false" @click="checkForUpdates" :render-icon="renderIcon(RefreshOutline)"-->
<!--                :loading="update_loading"/>-->
    </n-flex>
  </div>
</template>


<script setup>
import {renderIcon} from "../utils/common";
import {RefreshOutline} from "@vicons/ionicons5";
import {GetVersion} from "../../wailsjs/go/main/App";
import {CheckUpdate} from "../../wailsjs/go/system/Update";
import {ref} from "vue";
import {useMessage, useNotification} from "naive-ui";

defineProps({
  options: {},
  value: {},
});
const emit = defineEmits(['checkForUpdates'])
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