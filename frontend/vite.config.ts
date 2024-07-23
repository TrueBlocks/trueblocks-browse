import react from '@vitejs/plugin-react';
import tsconfigPaths from 'vite-tsconfig-paths';
import { defineConfig } from 'vite';

export default defineConfig({
  plugins: [react(), tsconfigPaths()],
  server: {
    watch: {
      usePolling: true,
    },
  },
});

//    "paths": {
//      "@assets/*": [
//        "./assets/*"
//      ],
//      "@components": [
//        "./components/index"
//      ],
//      "@views": [
//        "./views/index"
//      ],
//      "@hooks": [
//        "./hooks/index"
//      ],
//      "@gocode/*": [
//        "../wailsjs/go/*"
//      ],
//      "@runtime": [
//        "../wailsjs/runtime/runtime"
//      ],
//      "@/*": [
//        "./*"
//      ]
//
// import path from "path";
// import react from "@vitejs/plugin-react";
// import { defineConfig } from "vite";
// 
// export default defineConfig({
//   plugins: [react()],
//   resolve: {
//     alias: {
//       "@": path.resolve(__dirname, "./src"),
//       "@assets": path.resolve(__dirname, "./src/assets"),
//       "@components": path.resolve(__dirname, "./src/components/index"),
//       "@gocode": path.resolve(__dirname, "./wailsjs/go")
//     }
//   },
//   server: {
//     watch: {
//       usePolling: true,
//     },
//   },
// });
