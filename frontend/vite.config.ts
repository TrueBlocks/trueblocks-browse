import react from "@vitejs/plugin-react";
import tsconfigPaths from 'vite-tsconfig-paths';
import { defineConfig } from "vite";
import mdPlugin from 'vite-plugin-md';

export default defineConfig({
  plugins: [
    react(),
    tsconfigPaths(),
    mdPlugin({
      mode: 'html', // 'html' mode will handle Markdown as raw HTML/text content
    }),
  ],
  server: {
    host: 'localhost',
    port: 5173,
    watch: {
      usePolling: true,
    },
    hmr: {
      host: 'localhost',
      port: 5173,
    },
  },
});
