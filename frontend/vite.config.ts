import path from "path";
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";

export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@assets": path.resolve(__dirname, "./src/assets"),
      "@components": path.resolve(__dirname, "./src/components/index"),
      "@gocode": path.resolve(__dirname, "./wailsjs/go")
    }
  },
  server: {
    watch: {
      usePolling: true,
    },
  },
});
