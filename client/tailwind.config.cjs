/** @type {import('tailwindcss').Config} */
module.exports = {
    mode: "jit",
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {
            fontFamily: {
                Vistol: ["Vistol", "sans-serif"],
            },
        },
    },
    plugins: [],
}
