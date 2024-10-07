
// 渲染图标给菜单
import {h} from "vue";
import {NIcon} from "naive-ui";

export function renderIcon(icon) {
  return () => h(NIcon, null, {default: () => h(icon)});
}

