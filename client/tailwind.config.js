/** @type {import('tailwindcss').Config} */
export default {
  content: [
    // all directories and extensions will correspond to your Nuxt config
    "./components/**/*.{vue,js,ts}",
    "./layouts/**/*.{vue,js,ts}",
    "./pages/**/*.{vue,js,ts}",
    "./plugins/**/*.{js,ts}",
    "./composables/**/*.{js,ts}",
    "./utils/**/*.{js,ts}",
    "./App.{vue,js,ts}",
    "./Error.{vue,js,ts}",
    "./app.config.{js,ts}",
    "./app/spa-loading-template.html",
  ],
  theme: {
    extend: {
      colors: {
        "sidebar-robocare": "#f7f0f1",
        "sidebar-primary": "#ea6c5a",
        "card-foreground": "#f7f0f1",
        background: "#0c0809",
      },
    },
  },
  plugins: [],
};
