/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
  ],
  theme: {
    fontFamily: {
      'sans': ['Open Sans', 'ui-sans-serif', 'system-ui'],
      'dots': ['dotsfont'],
      'poppins-bold': ['Poppins-Bold'],
      'poppins-light': ['Poppins-Light']
    },
    screens: {
      'mh-1080': { 'raw': '(min-height: 1080px)' }
    },
    extend: {
      colors: {
        'background': '#191019',
        'foreground': '#140D14',
        'primary': '#EA5600',
      }
    }
  },
  plugins: []
}
