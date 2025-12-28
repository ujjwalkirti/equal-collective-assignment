import * as React from "react";
import { cn } from "../utils";

const styles = {
  muted: "bg-white/5 text-slate-200",
  accent: "bg-accent-mint/15 text-accent-mint",
  warn: "bg-amber-500/20 text-amber-200",
  error: "bg-coral-500/20 text-coral-100",
  success: "bg-emerald-500/20 text-emerald-100",
} as const;

type Variant = keyof typeof styles;

export function Badge({ variant = "muted", className, ...props }: React.HTMLAttributes<HTMLSpanElement> & { variant?: Variant }) {
  return (
    <span
      className={cn(
        "inline-flex items-center gap-2 rounded-full px-3 py-1 text-xs font-semibold uppercase tracking-wide",
        styles[variant],
        className
      )}
      {...props}
    />
  );
}
