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

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import svgLoader from "vite-svg-loader";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 5174,
  },
  minify: 'terser', // 或 'esbuild'（更快，但压缩率稍低）
  terserOptions: {
    compress: {
      drop_console: true, // 移除 console.log
      pure_funcs: ['console.info'], // 指定移除特定的函数调用
    },
  },
  publicDir: 'public',
  plugins: [
    vue(),svgLoader()
  ],
})
