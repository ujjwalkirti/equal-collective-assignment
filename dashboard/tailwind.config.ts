import type { Config } from "tailwindcss";
import { fontFamily } from "tailwindcss/defaultTheme";

const config: Config = {
  content: [
    "./app/**/*.{ts,tsx}",
    "./components/**/*.{ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        canvas: {
          DEFAULT: "#0b1021",
          subtle: "#10162e",
          edge: "#1f2a4d",
        },
        accent: {
          mint: "#5cf0c8",
          coral: "#ff7a7a",
          amber: "#ffc857",
        },
      },
      fontFamily: {
        display: ["var(--font-display)", ...fontFamily.sans],
      },
      boxShadow: {
        glow: "0 20px 80px rgba(92, 240, 200, 0.15)",
      },
      backgroundImage: {
        mesh: "radial-gradient(circle at 20% 20%, rgba(92, 240, 200, 0.08), transparent 25%), radial-gradient(circle at 80% 10%, rgba(255, 200, 87, 0.09), transparent 20%), radial-gradient(circle at 60% 80%, rgba(255, 122, 122, 0.08), transparent 18%)",
      },
    },
  },
  plugins: [],
};

export default config;
