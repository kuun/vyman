import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
  ],
  build: {
    sourcemap: true
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  },
  //本地开发时打开可以远程调试前端代码，调试时改为自己的服务器地址和端口即可
  server: {
    proxy: {
      '/api': {
        target: 'http://192.168.0.237:8000/api',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, '')
      },
      '/ui': {
        target: 'http://192.168.0.237:8000/ui',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/ui/, '')
      },
      '/assets': {
        target: 'http://192.168.0.237:8000/assets',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/assets/, '')
      }
    }
  }
})
