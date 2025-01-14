import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import { nodeResolve } from '@rollup/plugin-node-resolve';

export default defineConfig({
	plugins: [
        sveltekit(),
        nodeResolve({
            browser: true,
        }),
    ]
});
