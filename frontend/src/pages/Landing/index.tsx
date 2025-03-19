import React from "react";

const Landing: React.FC = () => {
    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-black text-green-500">
            {/* Outer terminal-like container */}
            <div className="bg-black p-6 rounded-lg shadow-lg border-4 border-green-500 w-[80%] max-w-3xl text-center">

                {/* AI Game Title */}
                <div className="border border-green-500 p-4 mb-6 flex justify-center items-center w-full">
                    <h1 className="text-7xl md:text-8xl font-bold text-green-500">{`{ Ai Game }`}</h1>
                </div>

                {/* Play Game Button */}
                <div className="border border-green-500 p-3 inline-block cursor-pointer hover:bg-green-700 hover:text-black transition duration-300">
                    <button className="text-lg">&gt;&gt; Play Game! &lt;&lt;</button>
                </div>

                {/* Terminal Cursor Icon */}
                <div className="absolute bottom-10 left-10 text-4xl font-bold">
                    <span className="text-green-500">&gt;_</span>
                </div>
            </div>
        </div>
    );
};

export default Landing;
