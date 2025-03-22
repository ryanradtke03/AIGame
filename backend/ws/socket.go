package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"sync"
	"encoding/json"
)

// RoomConnections holds all connections for a room tracked by room code
var RoomConnections = make(map[string][]*websocket.Conn)
// Mutex for RoomConnections
var mu sync.Mutex

// WebSocketHandler handles socket connections
// Check if client is trying to connect to a websocket :roomCode
// Extract room code from URL and pass it to next handler
// Stores roomCode in the conext for later use
// Passes control to the next handler (JoinRoomSocket)
// If not, return 426 Upgrade Required
func WebSocketHandler(c *fiber.Ctx) error {
	// Get room code from URL
	roomCode := c.Params("roomCode")

	// Check if room code is provided
	if roomCode == "" {
		return c.Status(400).SendString("Room code required")
	}

	// Check if upgrade to websocket (this means the client is requesting a websocket connection)
	// If not, return 426 Upgrade Required
	if websocket.IsWebSocketUpgrade(c) {
		// Pass room code to next handler
		c.Locals("roomCode", roomCode)
		// Continue to next handler
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}


func JoinRoomSocket(c *websocket.Conn) {
	roomCode := c.Locals("roomCode").(string)

	// Add connection to room list
	// Retrieve roomCode from RoomConnections 
	mu.Lock()
	RoomConnections[roomCode] = append(RoomConnections[roomCode], c)
	mu.Unlock()

	type IncomingMessage struct {
		Event    string `json:"event"`
		Message  string `json:"message,omitempty"`
		Username string `json:"username,omitempty"`
	}


	// Loop through messages and broadcast to others in room
	// If there is an error, break the loop
	for{
		_, raw, err := c.ReadMessage()
		if err != nil {
			break
		}

		var msg IncomingMessage
		if err := json.Unmarshal(raw, &msg); err != nil {
			continue // skip Invalid JSON
		}

		// Handle different event types
		switch msg.Event {
		// Chat event
		case "chat":
			// Broadcast chat message to room
			outgoing := map[string]string{
				"event":    "chat",
				"username": msg.Username,
				"message":  msg.Message,
			}
			BroadcastJSON(roomCode, outgoing)
		
		// Add more event types here


		// Default case for invalid event type
		default:
			outgoing := map[string]string{
				"event": "error",
				"message": "Invalid event type",
			}
			c.WriteJSON(outgoing)
		}
	}

	// Remove on disconnect
	mu.Lock()
	conns := RoomConnections[roomCode]
	for i, conn := range conns {
		if conn == c {
			RoomConnections[roomCode] = append(conns[:i], conns[i+1:]...)
			break
		}
	}
	mu.Unlock()
}


// BroadcastJSON sends a JSON message to all connections in a room
func BroadcastJSON(roomCode string, data interface{}) {
	mu.Lock()
	defer mu.Unlock()

	for _, conn := range RoomConnections[roomCode] {
		conn.WriteJSON(data)
	}
}

// Sends a string
func broadcastToRoom(roomCode string, msg []byte) {
	mu.Lock()
	defer mu.Unlock()
	for _, conn := range RoomConnections[roomCode] {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}