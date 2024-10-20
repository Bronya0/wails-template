
// 渲染图标给菜单
import {h} from "vue";
import {NIcon} from "naive-ui";
import {BrowserOpenURL} from "../../wailsjs/runtime";

// 渲染图标
export function renderIcon(icon) {
  return () => h(NIcon, null, {default: () => h(icon)});
}

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
