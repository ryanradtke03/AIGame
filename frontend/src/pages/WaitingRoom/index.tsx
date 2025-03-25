// Inputs
import RoomPlayerCard from "../../components/roomPlayer";
import Button from "../../components/ui/Button";

// Types
import { RoomPlayer } from "../../types/types";

// Hooks
import { useEffect, useState } from "react";

// Services
import { getPlayersInRoom } from "../../services/playerService";

// Socket
//import { createRoomSocket } from "../../services/socket/socket";

const WaitingRoom = () => {
  const roomCode = sessionStorage.getItem("roomCode");
  const roomId = sessionStorage.getItem("roomId");
  const sessionId = sessionStorage.getItem("sessionId");

  const [players, setPlayers] = useState<RoomPlayer[]>([]);
  const [isOwner, setIsOwner] = useState(false);

  // Grab players in room
  useEffect(() => {
    const fetchPlayers = async () => {
      if (roomId === null || !sessionId || roomCode === null) return;

      //const socket = createRoomSocket(roomCode);

      try {
        const data = await getPlayersInRoom(roomId);
        setPlayers(data);

        const self = data.find((p) => p.sessionID === sessionId);
        setIsOwner(!!self?.isOwner);
      } catch (err) {
        console.error("Failed to fetch players:", err);
      }
    };

    if (roomId && sessionId && roomCode) {
      console.log("Fetching players...");
      fetchPlayers();
    }
  }, [roomId, sessionId, roomCode]);

  const handleStartRoom = () => {
    // Start game
  };

  return (
    <div className="p-8 text-center text-xl font-bold">
      <div>
        Welcome to <span className="text-blue-500">Waiting Room</span>
      </div>
      <div>
        Room code: <span className="text-blue-500">{roomCode}</span>
      </div>

      {isOwner && (
        <div className="my-4">
          <Button onClick={handleStartRoom} variant="terminal" size="md">
            Start Game
          </Button>
        </div>
      )}

      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 mt-6">
        {players.map((player, idx) => (
          <RoomPlayerCard
            key={idx}
            username={player.username}
            points={player.points}
            score={player.score}
            isAI={player.isAI}
            variant={player.isOwner ? "primary" : "terminal"}
          />
        ))}
      </div>
    </div>
  );
};

export default WaitingRoom;
