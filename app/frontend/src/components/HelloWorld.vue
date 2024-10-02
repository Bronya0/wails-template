<template>

  <main>
    <div class="result">{{ data.resultText }}</div>
    <div class="input-box">
      <n-input v-model:value="data.name" type="text" placeholder="input name" class="input"/>
      <n-button type="info" @click="greet"> Greet </n-button>
      <n-button type="info" @click="$emit('selectTab', data.name)"> Select </n-button>
    </div>
  </main>
</template>


<script setup>
import {onActivated, onBeforeUnmount, onMounted, reactive} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'
import { useMessage } from 'naive-ui'
import {onDeactivated} from "@vue/runtime-core";

const message = useMessage()

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
    message.info(result)
  })
}


onMounted(() => {
  console.info("初始化HelloWorld……")
})

onBeforeUnmount(() => {
  console.info("卸载HelloWorld……")
})

onActivated(() => {
  console.info("激活HelloWorld组件……")
})
onDeactivated(() => {
  console.info("失活HelloWorld组件……")
})

</script>


<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}
.input-box {
  text-align: center;
}
.input {
  width: 200px;
  margin-right: 10px;
}
</style>
