// import adapter from '@sveltejs/adapter-auto';
import adapter from '@sveltejs/adapter-static';

import { vitePreprocess } from '@sveltejs/kit/vite';


/** @type {import('@sveltejs/kit').Config} */
const config = {
	kit: {
		// adapter: adapter()
		adapter: adapter({
            // default options are shown. On some platforms
            // these options are set automatically — see below
            pages: 'build',
            assets: 'build',
            // fallback: null,
            fallback: '200.html', 
            precompress: false,
            strict: true
        }),
		// prerender: {
		// 	entries: ['/app/[slug]'],
		// }
	},
	preprocess: vitePreprocess()

};

export default config;
