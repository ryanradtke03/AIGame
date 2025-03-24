import { API_BASE } from "../../utils/env/env";

export const createRoom = async (username: string) => {
  const res = await fetch(`${API_BASE}/rooms/create`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username }),
  });
  return await res.json();
};

export const joinRoom = async (username: string, code: string) => {
  const res = await fetch(`${API_BASE}/rooms/join`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, code }),
  });
  return await res.json();
};
