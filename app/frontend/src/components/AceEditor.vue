<template>
  <div :id="editorId" class="ace-editor" :style="{ height: props.height, width: props.width }"></div>
</template>

<script setup>
import {nextTick, onBeforeUnmount, onMounted, ref, watch} from 'vue'
import * as ace from 'ace-builds'

// ===========需要支持的语言需要自行导入===========
import 'ace-builds/src-noconflict/mode-text'
import 'ace-builds/src-noconflict/mode-json'
// import 'ace-builds/src-noconflict/mode-javascript'
import 'ace-builds/src-noconflict/mode-golang'
import 'ace-builds/src-noconflict/mode-python'
import 'ace-builds/src-noconflict/mode-yaml'
// import 'ace-builds/src-noconflict/mode-html'
// import 'ace-builds/src-noconflict/mode-css'
// import 'ace-builds/src-noconflict/mode-vue'
import 'ace-builds/src-noconflict/mode-toml'
// import 'ace-builds/src-noconflict/mode-rust'
import 'ace-builds/src-noconflict/mode-markdown'
// import 'ace-builds/src-noconflict/mode-nginx'
import 'ace-builds/src-noconflict/mode-sql'

import 'ace-builds/src-noconflict/worker-json'
import 'ace-builds/src-noconflict/worker-base'


import 'ace-builds/src-noconflict/theme-monokai'
// import 'ace-builds/src-noconflict/theme-clouds_midnight'
import 'ace-builds/src-noconflict/ext-language_tools'

// 定义 props
const props = defineProps({
  value: {
    type: String,
    default: ''
  },
  mode: {
    type: String,
    default: 'text'
  },
  theme: {
    type: String,
    default: 'monokai'
  },
  fontSize: {
    type: Number,
    default: 15
  },
  options: {
    type: Object,
    default: () => ({
      tabSize: 2,
      useSoftTabs: true,
      showPrintMargin: false,
      enableLiveAutocompletion: true,
      animatedScroll: true
    })
  },
  enableAutocomplete: {
    type: Boolean,
    default: true
  },
  showLineNumbers: {
    type: Boolean,
    default: true
  },
  readOnly: {
    type: Boolean,
    default: false
  },
  wrap: {
    type: Boolean,
    default: false
  },
  width: {
    type: String,
    default: '100%'
  },
  height: {
    type: String,
    default: '400px'
  },
  placeholder: {
    type: String,
    default: ''
  }
})

// 定义 emits
const emit = defineEmits(['update:value', 'change', 'ready'])

// 生成唯一编辑器 ID
const editorId = `ace-editor-${Math.random().toString(36).substr(2, 9)}`
const editor = ref(null)

// 配置：说明：默认值：推荐值
// 编辑器选项 (Editor Options)
// selectionStyle: String - 选择高亮样式，"line"整行高亮，"text"仅文本 ("line", "line")
// highlightActiveLine: Boolean - 高亮当前行 (true, true)
// highlightSelectedWord: Boolean - 高亮选中的单词 (true, true)
// readOnly: Boolean - 编辑器只读模式 (false, false)
// cursorStyle: String - 光标样式，"ace"标准，"slim"细，"smooth"平滑，"wide"宽 ("ace", "ace")
// mergeUndoDeltas: Boolean | String - 合并撤销操作，"always"总是合并 (true, true)
// behavioursEnabled: Boolean - 启用自动配对括号等行为 (true, true)
// wrapBehavioursEnabled: Boolean - 自动换行时保持行为 (true, true)
// autoScrollEditorIntoView: Boolean - 自动滚动使编辑器可见 (false, false)
// copyWithEmptySelection: Boolean - 允许复制空选择 (false, false)
// useSoftTabs: Boolean - 使用空格代替制表符 (true, true)
// navigateWithinSoftTabs: Boolean - 光标在软制表符内导航 (false, false)
// enableMultiselect: Boolean - 允许多重选择 (true, true)
// enableAutoIndent: Boolean - 启用自动缩进 (true, true)
// enableKeyboardAccessibility: Boolean - 启用键盘可访问性 (false, false)
// keyboardHandler: Function | null - 自定义键盘处理程序 (null, null)
// enableBasicAutocompletion: Boolean - 启用基本代码补全，需 ext-language_tools.js (false, true)
// enableLiveAutocompletion: Boolean - 启用实时代码补全，需 ext-language_tools.js (false, false)
// enableSnippets: Boolean - 启用代码片段，需 ext-language_tools.js (false, false)
// enableEmmet: Boolean - 启用Emmet支持，需 ext-emmet.js (false, false)
// useElasticTabstops: Boolean - 启用弹性制表位，需 ext-elastic_tabstops_lite.js (false, false)

// 渲染器选项 (Renderer Options)
// animatedScroll: Boolean - 启用滚动动画 (false, false)
// showInvisibles: Boolean - 显示不可见字符（如空格、制表符） (false, false)
// showPrintMargin: Boolean - 显示打印边距线 (true, true)
// printMarginColumn: Number - 打印边距列位置 (80, 80)
// showGutter: Boolean - 显示行号侧栏 (true, true)
// displayIndentGuides: Boolean - 显示缩进指南 (true, true)
// fontSize: Number | String - 字体大小，单位px或CSS单位 (12, 14)
// fontFamily: String - 字体家族，CSS字体栈 ("monospace", "Consolas, monospace")
// maxLines: Number - 最大显示行数，0禁用 (0, 0)
// minLines: Number - 最小显示行数，0禁用 (0, 0)
// scrollPastEnd: Boolean | Number - 允许滚动超出文档末尾，数字表示倍数 (0, 0.5)
// fixedWidthGutter: Boolean - 固定行号侧栏宽度 (false, false)
// theme: String - 主题路径，如"ace/theme/monokai" (null, "ace/theme/monokai")
// showFoldWidgets: Boolean - 显示代码折叠控件 (true, true)
// fadeFoldWidgets: Boolean - 折叠控件淡入淡出效果 (false, false)
// showLineNumbers: Boolean - 显示行号 (true, true)
// highlightGutterLine: Boolean - 高亮行号侧栏当前行 (true, true)
// hScrollBarAlwaysVisible: Boolean - 水平滚动条始终可见 (false, false)
// vScrollBarAlwaysVisible: Boolean - 垂直滚动条始终可见 (false, false)
// scrollSpeed: Number - 滚动速度倍数 (2, 2)
// useTextareaForIME: Boolean - 使用textarea处理输入法 (true, true)
// useSvgGutterIcons: Boolean - 行号侧栏使用SVG图标 (true, true)
// useWorker: Boolean - 启用后台语法检查 (true, true)
// placeholder: 占位符

// 会话选项 (Session Options)
// firstLineNumber: Number - 起始行号 (1, 1)
// overwrite: Boolean - 覆盖模式，插入覆盖现有字符 (false, false)
// newLineMode: String - 新行模式，"auto"|"unix"|"windows" ("auto", "unix")
// useWrapMode: Boolean - 启用自动换行 (false, true)
// wrap: Number | Boolean | String - 换行列数，"off"|40|"free"|"printMargin" (false, 80)
// tabSize: Number - 制表符宽度 (4, 2)
// indentUnit: Number - 缩进单位空格数 (4, 2)
// foldStyle: String - 折叠样式，"markbegin"|"markbeginend"|"manual" ("markbegin", "markbegin")
// mode: String - 语言模式，如"ace/mode/javascript" (null, "ace/mode/text")

// 鼠标处理器选项 (MouseHandler Options)
// scrollSpeed: Number - 鼠标滚动速度倍数 (2, 2)
// dragDelay: Number - 拖拽延迟，毫秒 (150, 150)
// dragEnabled: Boolean - 启用鼠标拖拽选择 (true, true)
// focusTimeout: Number - 鼠标聚焦延迟，毫秒 (0, 0)
// tooltipFollowsMouse: Boolean - 工具提示跟随鼠标 (true, true)


// 初始化 Ace 编辑器
onMounted(async () => {
  await nextTick()
  ace.config.set('basePath', '/node_modules/ace-builds/src-noconflict')

  editor.value = ace.edit(editorId, {
    mode: `ace/mode/${props.mode}`,
    theme: `ace/theme/${props.theme}`,
    value: props.value,
    fontSize: props.fontSize,
    enableBasicAutocompletion: props.enableAutocomplete,
    enableLiveAutocompletion: props.enableAutocomplete,
    enableSnippets: props.enableAutocomplete,
    showLineNumbers: props.showLineNumbers,
    readOnly: props.readOnly,
    placeholder: props.placeholder,
    wrap: props.wrap,
    ...props.options
  })

  // 监听内容变化
  editor.value.on('change', () => {
    const content = editor.value.getValue()
    emit('update:value', content)
    emit('change', content)
  })

  // 监听失去焦点事件，仅当 mode 为 json 时格式化
  editor.value.on('blur', () => {
    if (props.mode === 'json') {
      let currentValue = editor.value.getValue()
      let formatted;
      try {
        const parsed = JSON.parse(currentValue)
        formatted = JSON.stringify(parsed, null, 2)
      } catch (e) {
        formatted = currentValue
      }

      // 更新编辑器内容
      editor.value.setValue(formatted, -1)
      emit('update:value', formatted)
      emit('change', formatted)
    }
  })

  // 触发 ready 事件
  emit('ready', editor.value)
})

// 同步外部传入的 value
watch(() => props.value, (newValue) => {
  if (newValue !== editor.value?.getValue()) {
    editor.value?.setValue(newValue, -1)
  }
})

// 监听模式变化
watch(() => props.mode, (newMode) => {
  editor.value?.session.setMode(`ace/mode/${newMode}`)
  // 模式切换时清除错误提示
  editor.value?.getSession().clearAnnotations()
})

// 监听主题变化
watch(() => props.theme, (newTheme) => {
  editor.value?.setTheme(`ace/theme/${newTheme}`)
})

// 清理
onBeforeUnmount(() => {
  editor.value?.destroy()
  editor.value?.container.remove()
})

// 定义提示词数据
// const completions = [
// {
//   caption: "console.log", // 用户看到的名称（可选，如果和 value 相同可以省略）
//   value: "console.log(${1:value})", // 实际插入的值
//   score: 100, // 权重（越高越靠前）
//   meta: "JavaScript" // 分类（如 "JavaScript"、"CSS"、"Custom"）
// },
// ];

// 暴露方法给父组件
defineExpose({
  getEditor: () => editor.value,
  setValue: (value) => editor.value?.setValue(value, -1),
  getValue: () => editor.value?.getValue(),
  setOptions: (options) => editor.value?.setOptions(options),
  addCompleter: (completer) => {
    if (props.enableAutocomplete) {
      editor.value?.completers.push(completer)
    }
  },
  setCompleter: (completer) => {
    if (props.enableAutocomplete) {
      editor.value.completers = [completer]
    }
  }
})
</script>

<style>
/* 隐藏ace编辑器的脱离聚焦时携带的光标 */
.ace_editor:not(.ace_focus) .ace_cursor {
  opacity: 0 !important;
}

/* 使主题支持placeholder */
.ace_editor .ace_placeholder {
  position: absolute;
  z-index: 10;
}
</style>