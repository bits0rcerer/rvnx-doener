import postcss from 'rollup-plugin-postcss';
import resolve from 'rollup-plugin-node-resolve';

export default {
	plugins: [
		resolve({
			dedupe: ['svelte', 'svelte/transition', 'svelte/internal'] // important!
		}),
		// eslint-disable-next-line no-undef
		svelte({
			emitCss: true
		}),
		postcss({
			extract: true
		})
	]
};
