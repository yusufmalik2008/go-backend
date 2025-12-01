/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./web/templates/**/*.{templ,gohtml}"],
  theme: { extend: {} },
  plugins: [require("daisyui")],
  daisyui: { themes: ["dark", "light"] }
}
