// src/pages/Landing/index.tsx
import React from "react";

const Voting: React.FC = () => {
    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
            <h1 className="text-4xl font-bold text-blue-600">Welcome to the Voting Page!</h1>
            <p className="text-lg mt-4">This is a simple page created in React using TypeScript.</p>
        </div>
    );
};

export default Voting;