// Compoents: Button, Input
import Button from "../../components/ui/Button";
import Input from "../../components/ui/Input";

// Hooks: useState
import { useState } from "react";
import { useNavigate } from "react-router-dom";

// Services: joinRoom, createRoom
import { createRoom, joinRoom } from "../../services/roomService";

const Landing = () => {
  // States
  const [username, setUsername] = useState("");
  const [roomCode, setRoomCode] = useState("");

  const navigate = useNavigate();

  // Handlers
  const handleJoinRoom = async () => {
    if (!username || !roomCode) {
      alert("Username and Room Code are required to join.");
      return;
    }

    try {
      const response = await joinRoom(username, roomCode);
      console.log("Joined Room:", response);
      // maybe navigate or update global state
    } catch (err) {
      console.error(err);
      alert("Failed to join room.");
    }
  };

  const handleCreateRoom = async () => {
    // Check for username
    if (!username) {
      alert("Username is required to create a room.");
      return;
    }

    try {
      const response = await createRoom(username);
      sessionStorage.setItem("roomCode", response.room_code);

      console.log("Room Created:", response);
      navigate("/chat");
    } catch (err) {
      console.error(err);
      alert("Failed to create room.");
    }
  };

  return (
    <div className="max-w-md mx-auto p-6 space-y-6">
      <h1 className="text-terminal-green text-2xl font-mono text-center">
        {"{ AI Game }"}
      </h1>
      {/* UserName */}
      <Input
        label="Username"
        placeholder="Enter your username"
        value={username}
        onChange={(e) => setUsername(e.target.value)}
      />
      {/* RoomCode */}
      <Input
        label="Room Code (if joining)"
        placeholder="Enter room code"
        value={roomCode}
        onChange={(e) => setRoomCode(e.target.value)}
      />
      <div className="flex gap-4">
        {/* Join Room */}
        <Button onClick={handleJoinRoom} variant="terminal" size="md">
          Join Room
        </Button>

        {/* Create ROom */}
        <Button onClick={handleCreateRoom} variant="primary" size="md">
          Create Room
        </Button>
      </div>
    </div>
  );
};

export default Landing;
