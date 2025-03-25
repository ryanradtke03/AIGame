import { RoomPlayer } from "../../types/types";
import { API_BASE } from "../../utils/env/env";

export const getPlayersInRoom = async (roomId: string) => {
  const res = await fetch(`${API_BASE}/players/in-room/${roomId}`, {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  });
  const data: RoomPlayer[] = await res.json();
  return data;
};
