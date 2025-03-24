import clsx from "clsx";
import React from "react";

type ButtonVariant = "primary" | "danger" | "terminal";
type ButtonSize = "sm" | "md" | "lg";

interface ButtonProps {
  children: React.ReactNode;
  onClick?: () => void;
  type?: "button" | "submit" | "reset";
  disabled?: boolean;
  variant?: ButtonVariant;
  size?: ButtonSize;
  className?: string;
}

const variantClasses: Record<ButtonVariant, string> = {
  primary: "bg-blue-500 hover:bg-blue-600 text-white",
  danger: "bg-terminal-red hover:bg-red-700 text-white",
  terminal:
    "bg-terminal-black text-terminal-green border border-terminal-green hover:bg-terminal-green/10",
};

const sizeClasses: Record<ButtonSize, string> = {
  sm: "text-sm px-3 py-1",
  md: "text-base px-4 py-2",
  lg: "text-lg px-6 py-3",
};

const Button: React.FC<ButtonProps> = ({
  // Default props
  children,
  className = "",
  variant = "terminal",
  size = "md",
  ...props
}) => {
  return (
    <button
      className={clsx(
        "rounded-md font-mono transition duration-200 disabled:opacity-50 w-full",
        variantClasses[variant],
        sizeClasses[size],
        className
      )}
      {...props}
    >
      {children}
    </button>
  );
};

export default Button;
