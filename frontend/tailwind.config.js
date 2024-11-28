/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,jsx,ts,tsx}", // Ensures Tailwind processes your React files
  ],
  theme: {
    extend: {
      boxShadow: {
        shadow: "4px 4px 0px 0px #000",
      },
      colors: {
        garconia: {
          red: "#A31621",
          background: "#211A1E",
          offwhite: "#B7B5B3",
          white: "#F3F6F5",
        },
      },
    },
  },
  plugins: [],
};
