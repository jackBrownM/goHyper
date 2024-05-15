import jiti from "file:///E:/developProject/go-hyper/web/node_modules/.pnpm/jiti@1.21.0/node_modules/jiti/lib/index.js";

/** @type {import("E:/developProject/go-hyper/web/internal/vite-config/src/index")} */
const _module = jiti(null, {
  "esmResolve": true,
  "interopDefault": true,
  "alias": {
    "@vben/vite-config": "E:/developProject/go-hyper/web/internal/vite-config"
  }
})("E:/developProject/go-hyper/web/internal/vite-config/src/index.ts");

export const defineApplicationConfig = _module.defineApplicationConfig;
export const definePackageConfig = _module.definePackageConfig;