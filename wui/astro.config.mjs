// @ts-check
import { defineConfig } from 'astro/config';

import tailwind from '@astrojs/tailwind';

import htmx from 'astro-htmx';

import icon from 'astro-icon';

// https://astro.build/config
export default defineConfig({
  integrations: [tailwind(), icon(), htmx()],
  outDir: '../build/public',
  // Esto es para que astro pueda tomar páginas dentro de astro.
  output:'server'
});