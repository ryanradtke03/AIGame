import { WS_BASE } from "../../utils/env/env";

export const createRoomSocket = (roomCode: string): WebSocket => {
  const protocol = window.location.protocol === "https:" ? "wss" : "ws";
  const socket = new WebSocket(`${protocol}://${WS_BASE}/ws/${roomCode}`);
  return socket;
};
