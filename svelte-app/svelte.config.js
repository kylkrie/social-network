import adapter from "@sveltejs/adapter-static";
import { vitePreprocess } from "@sveltejs/vite-plugin-svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
  preprocess: vitePreprocess(),
  kit: {
    adapter: adapter({
      // `fallback` is the key config for SPA mode
      fallback: "index.html",
      // Specify the directory to output the built files
      pages: "build",
      assets: "build",
      precompress: false,
    }),
    prerender: {
      entries: [], // This ensures no pages are prerendered
    },
  },
};

export default config;
