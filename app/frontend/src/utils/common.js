/*
 * Copyright 2025 Bronya0 <tangssst@163.com>.
 * Author Github: https://github.com/Bronya0
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// 渲染图标给菜单
import {h} from "vue";
import {NIcon, NTooltip} from "naive-ui";
import {BrowserOpenURL} from "../../wailsjs/runtime";

// 渲染图标
export function renderIcon(icon) {
  return () => h(NIcon, null, {default: () => h(icon)});
}
export const renderSelect = ({ node, option }) => h(NTooltip, {
  placement: "right",
}, {trigger: () => node, default: () => `${option.label}`});

// 打开链接
export function openUrl(url) {
  BrowserOpenURL(url)
}

// 压扁json
export function flattenObject(obj, parentKey = '') {
  let flatResult = {};

  for (let key in obj) {
    if (obj.hasOwnProperty(key)) {
      let newKey = parentKey ? `${parentKey}.${key}` : key;

      if (typeof obj[key] === 'object' && obj[key] !== null && !Array.isArray(obj[key])) {
        // 如果当前值也是一个对象，则递归调用
        Object.assign(flatResult, flattenObject(obj[key], newKey));
      } else {
        // 否则直接赋值
        flatResult[newKey] = obj[key];
      }
    }
  }

  return flatResult;
}
// 格式化的 JSON 字符串
export function formattedJson(value) {
  if (!value) return ''
  return JSON.stringify(value, null, 1)
}

// 验证json
export function isValidJson(jsonString) {
  try {
    // 尝试解析 JSON 字符串
    JSON.parse(jsonString);
    return true; // 解析成功，是有效的 JSON
  } catch (error) {
    // 解析失败，不是有效的 JSON
    return false;
  }
}

// 单位处理
export function formatBytes(bytes, decimals = 2) {
  if (bytes === 0) return '0 Bytes';
  if (bytes === null) return '';

  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];

  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
}


// 创建 CSV 内容的函数
export function createCsvContent(allData, columns) {
  // 过滤掉没有 title 的列
  columns = columns.filter(col => col.title !== undefined);

  const headers = columns.map(col => col.title).join(',')
  const rows = allData.map(row =>
      columns.map(col => row[col.key]).join(',')
  ).join('\n')
  return `${headers}\n${rows}`
}

// 下载件的函数，csv type：'text/csv;charset=utf-8;'
export function download_file(content, fileName, type) {
  const blob = new Blob([content], { type: type })
  const link = document.createElement('a')
  if (link.download !== undefined) {
    const url = URL.createObjectURL(blob)
    link.setAttribute('href', url)
    link.setAttribute('download', fileName)
    link.style.visibility = 'hidden'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }
}
// 日期格式化函数
export function formatDate(date){
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

export function getCurrentDateTime() {
  const now = new Date();

  // 获取年份、月份、日期、小时、分钟、秒钟
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, '0'); // 月份从0开始，所以要加1
  const date = String(now.getDate()).padStart(2, '0');
  const hours = String(now.getHours()).padStart(2, '0');
  const minutes = String(now.getMinutes()).padStart(2, '0');
  const seconds = String(now.getSeconds()).padStart(2, '0');

  // 返回格式化的日期时间字符串
  return `${year}.${month}.${date}-${hours}.${minutes}.${seconds}`;
}

// 获取本地语言
export function getLocalLanguage() {
  return navigator.language;
}


/**
 * 智能处理表格列配置
 * minWidth 不会主动影响列的初始宽度，它只是作为拖拽调整时的最小宽度限制。
 *  初始宽度由以下优先级决定：
 *      width > 内容自适应宽度 > 表格容器的等分分配。
 * 1. 自动计算未配置 minWidth 的列（基于 title 长度）
 * 2. 默认允许拖动（除非显式禁用）
 * 3. 自动添加 ellipsis 省略效果
 * 4. 当允许拖动时，移除默认 width 配置（避免冲突）
 */
export function refColumns(columns) {
  return columns.map((column) => {
    // 深拷贝原始列配置（避免修改原对象）
    const processed = { ...column };
    if (processed.type === 'selection'){
      return processed;
    }
    if ("_ignore" in processed){
      delete processed._ignore;
      return processed;
    }

    if (!('sorter' in processed)) {
      processed.sorter = 'default';
    }
    // ===== 处理 resizable =====
    if (!('resizable' in processed)) {
      processed.resizable = true; // 默认允许拖动
    }

    if (processed.resizable){
      // ===== 处理 minWidth =====
      if (!('width' in processed)) {
        processed.width = calculateWidthByTitle(processed.title);
      }
    }

    // ===== 处理 ellipsis =====
    if (!('ellipsis' in processed)) {
      processed.ellipsis = {
        tooltip: {
          scrollable: true,
          style: { maxWidth: '800px' } // 自定义提示框最大宽度
        }
      };
    }
    return processed;
  });
}

/**
 * 根据标题文本计算建议宽度（中英文混合）
 */
function calculateWidthByTitle(title) {
  let width = 0;
  for (const char of title) {
    width += /[\u4e00-\u9fa5]/.test(char) ? 16 : 8; // 中文16px，英文8px
  }
  return Math.max(width + 24, 80); // 加 padding 且不低于 80px
}