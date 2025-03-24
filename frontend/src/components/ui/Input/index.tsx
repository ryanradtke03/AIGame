import clsx from "clsx";
import React from "react";

type InputVariant = "primary" | "danger" | "terminal";
type InputSize = "sm" | "md" | "lg";

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  variant?: InputVariant;
  inputSize?: InputSize;
  className?: string;
}

const variantClasses: Record<InputVariant, string> = {
  primary: "bg-white text-black border border-blue-500 focus:ring-blue-500",
  danger: "bg-terminal-red text-white border border-red-700 focus:ring-red-500",
  terminal:
    "bg-terminal-black text-terminal-green placeholder-terminal-greenLight border border-terminal-green focus:ring-terminal-green",
};

const sizeClasses: Record<InputSize, string> = {
  sm: "text-sm px-3 py-1",
  md: "text-base px-4 py-2",
  lg: "text-lg px-6 py-3",
};

const Input: React.FC<InputProps> = ({
  // Default props
  label,
  variant = "terminal",
  inputSize = "md",
  className = "",
  id,
  ...props
}) => {
  const inputId = id || props.name;

  return (
    <div className="flex flex-col space-y-1 w-full">
      {label && (
        <label
          htmlFor={inputId}
          className="text-terminal-green font-mono text-sm"
        >
          {label}
        </label>
      )}
      <input
        id={inputId}
        className={clsx(
          "rounded-md font-mono outline-none transition duration-200",
          "focus:ring-2 disabled:opacity-50 w-full",
          variantClasses[variant],
          sizeClasses[inputSize],
          className
        )}
        {...props}
      />
    </div>
  );
};

export default Input;
