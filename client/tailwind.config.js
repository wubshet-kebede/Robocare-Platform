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
        "muted-foreground": "#7a6f6c",
        card: "#150f0e",
        background: "#0f0f10",
        foreground: "#f4f4f5",

        card: "#171718",
        "card-foreground": "#f4f4f5",

        popover: "#171718",
        "popover-foreground": "#f4f4f5",

        primary: "#d97706",
        "primary-foreground": "#0f0f10",

        secondary: "#1f1f21",
        "secondary-foreground": "#f4f4f5",

        muted: "#1e1e20",
        "muted-foreground": "#8a8a8d",

        accent: "#222224",
        "accent-foreground": "#f4f4f5",

        destructive: "#dc2626",
        "destructive-foreground": "#fafafa",

        success: "#16a34a",
        "success-foreground": "#0f0f10",

        warning: "#f59e0b",
        "warning-foreground": "#0f0f10",

        border: "#2a2a2d",
        input: "#2a2a2d",
        ring: "#d97706",

        chart: {
          1: "#d97706",
          2: "#0891b2",
          3: "#16a34a",
          4: "#7c3aed",
          5: "#f59e0b",
        },

        sidebar: "#141416",
        "sidebar-foreground": "#f4f4f5",
        "sidebar-primary": "#d97706",
        "sidebar-primary-foreground": "#ffffff",
        "sidebar-accent": "#222224",
        "sidebar-accent-foreground": "#f4f4f5",
        "sidebar-border": "#2a2a2d",
        "sidebar-ring": "#d97706",
      },
    },
  },
  plugins: [],
};
