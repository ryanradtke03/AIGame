import clsx from "clsx";
import React from "react";

type CardVariant = "primary" | "danger" | "terminal";
type CardSize = "sm" | "md" | "lg";

interface RoomPlayerCardProps {
  username: string;
  isAI?: boolean;
  points?: number;
  score?: number;
  children?: React.ReactNode;
  onClick?: () => void;
  variant?: CardVariant;
  size?: CardSize;
  className?: string;
}

// Variant styles like your Button
const variantClasses: Record<CardVariant, string> = {
  primary: "bg-white border-blue-500 text-blue-700",
  danger: "bg-white border-red-500 text-red-700",
  terminal: "bg-terminal-black border-terminal-green text-terminal-green",
};

// Size styles like your Button
const sizeClasses: Record<CardSize, string> = {
  sm: "p-2 text-sm",
  md: "p-4 text-base",
  lg: "p-6 text-lg",
};

const RoomPlayerCard: React.FC<RoomPlayerCardProps> = ({
  username,
  isAI = false,
  points = 0,
  score = 0,
  children,
  onClick,
  variant = "terminal",
  size = "md",
  className = "",
}) => {
  return (
    <div
      onClick={onClick}
      className={clsx(
        "rounded-xl shadow border text-left min-w-[180px] transition duration-200",
        variantClasses[variant],
        sizeClasses[size],
        className
      )}
    >
      <h2 className="font-semibold">
        {username}
        {isAI && <span className="ml-2 text-xs">(AI)</span>}
      </h2>
      <div className="mt-2 space-y-1 text-sm">
        <p>
          Points: <span className="font-medium">{points}</span>
        </p>
        <p>
          Score: <span className="font-medium">{score}</span>
        </p>
      </div>
      {children && <div className="mt-3">{children}</div>}
    </div>
  );
};

export default RoomPlayerCard;
