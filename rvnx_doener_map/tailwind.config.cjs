/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        'purple': {
          'twitch' : '#9146FF'
        }
      }
    },
  },
  plugins: [],
}
