import {sveltekit} from '@sveltejs/kit/vite';

/** @type {import('vite').UserConfig} */
const config = {
    plugins: [sveltekit()],
    server: {
        host: false,
        proxy: {
            "/api": {
                target: "http://localhost:8080",
                changeOrigin: true,
                secure: false,
            }
        }
    },
    ssr: {
        noExternal: [
            '@fortawesome/free-solid-svg-icons',
            '@fortawesome/free-brands-svg-icons',
        ]
    }
};

export default config;
