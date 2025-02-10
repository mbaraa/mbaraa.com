/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./tmplrndr/html/**/*.html"],
  theme: {
    extend: {
      fontFamily: {
        Vistol: ["Vistol", "sans-serif"],
        Space: ["Space", "monospace"],
      },
      colors: {
        grey: "#212121",
        green: "#20DB8F",
        white: "#FFFFFF",
        black: "#000000",
      },
    },
  },
  plugins: [],
};
