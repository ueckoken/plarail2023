import { defineConfig } from 'vite'
import path from 'path'
import react from '@vitejs/plugin-react-swc'
import svgr from 'vite-plugin-svgr'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), svgr()],
  resolve: {
    alias: {
      '@': path.join(__dirname, 'src')
    }
  },
  css: {
    postcss: null,
    preprocessorOptions: {
        scss: {
            additionalData: `
                @import "../src/_variables.scss";
            `,
        },
    },
  },
})
