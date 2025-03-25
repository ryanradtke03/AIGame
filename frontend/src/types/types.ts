export interface RoomPlayer {
  id: number;
  sessionID: string;
  username: string;
  isAI: boolean;
  points: number;
  score: number;
  isOwner: boolean;
}
