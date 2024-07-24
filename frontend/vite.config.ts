import react from "@vitejs/plugin-react";
import tsconfigPaths from 'vite-tsconfig-paths';
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [react(), tsconfigPaths()],
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
