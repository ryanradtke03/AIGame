/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        terminal: {
          black: "#000000",
          green: "#39FF14",
          greenLight: "#69FF66",
          greenDark: "#28A92E",
          red: "#FF3233",
        },
      },
    },
  },
};
