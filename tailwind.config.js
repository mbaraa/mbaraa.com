module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        SFUI: ["SFUI", "sans-serif"],
      },
    },
  },
  plugins: [],
  corePlugins: {
    fontFamily: true,
  },
};
