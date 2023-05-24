/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    fontFamily :{
      'sans': ['Open Sans', 'ui-sans-serif', 'system-ui']
    },
    extend: {
      colors: {
        'background': '#051139',
        'foreground': '#111C44'
      }
    },
  },
  plugins: []
}
