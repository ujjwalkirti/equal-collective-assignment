import type { Metadata } from "next";
import { Space_Grotesk } from "next/font/google";
import "./globals.css";

const display = Space_Grotesk({ subsets: ["latin"], variable: "--font-display" });

export const metadata: Metadata = {
  title: "X-Ray Dashboard",
  description: "Inspect decision traces across multi-step pipelines",
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className={`${display.variable} bg-canvas text-slate-50 antialiased`}>{children}</body>
    </html>
  );
}
