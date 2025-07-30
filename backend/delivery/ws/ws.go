package ws

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	application "social-network/application/usecases"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pingPeriod     = 30 * time.Second
	maxMessageSize = 5120
)

type Client struct {
	UserId   string
	Hub      *Hub
	Conn     *websocket.Conn
	Send     chan []byte
	ClientId string
	LastPing time.Time
}

func (c *Client) generateClientId() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// Upgrader pour convertir la requête HTTP en WebSocket.
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		_, err := url.Parse(origin)
		if err != nil {
			return false
		}
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request, getUserId func(context.Context, string) string, m application.MessageUsecase) {
	cookie, err := r.Cookie("Token")
	if err != nil || cookie.Value == "" {
		log.Printf("❌ Cookie Token introuvable: %v", err)
		http.Error(w, "Token cookie required", http.StatusUnauthorized)
		return
	}
	userId := getUserId(r.Context(), cookie.Value)
	if userId == "" {
		log.Printf("❌ ID utilisateur introuvable pour le token")
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	log.Printf("🔗 Nouvelle connexion WebSocket pour l'utilisateur: %s", userId)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("❌ Erreur lors de l'upgrade WebSocket: %v", err)
		return
	}
	log.Printf("✅ Connexion WebSocket établie pour l'utilisateur: %s", userId)
	client := &Client{
		UserId:   userId,
		Hub:      hub,
		Conn:     conn,
		Send:     make(chan []byte, 256),
		LastPing: time.Now(),
	}
	client.ClientId = client.generateClientId()
	hub.Register <- client

	go client.writePump()
	go client.readPump(m, r)
}

func (c *Client) readPump(m application.MessageUsecase, r *http.Request) {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("⚠️ Erreur de fermeture inattendue: %v", err)
			}
			break
		}
		fmt.Println("heeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", string(msg))
		// On parse le message
		var incoming struct {
			Type    string `json:"type"`
			To      string `json:"to,omitempty"`
			Content string `json:"content,omitempty"`
		}
		if err := json.Unmarshal(msg, &incoming); err != nil {
			log.Printf("❌ Erreur JSON: %v", err)
			continue
		}
		switch incoming.Type {
		case "message":
			outgoing := []byte(fmt.Sprintf(
				`{"from":"%s","to":"%s","content":"%s"}`,
				c.UserId, incoming.To, incoming.Content,
			))
			m.SaveMessage(c.UserId, incoming.To, incoming.Content)
			c.Hub.SendToUser(incoming.To, outgoing)
		case "broadcast":
			c.Hub.Broadcast <- msg
		default:
			log.Printf("⚠️ Type inconnu: %s", incoming.Type)
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// Envoi du message
			w, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(c.Send)
			for range n {
				w.Write([]byte("\n"))
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (h *Hub) SendToUser(userID string, message []byte) {
	fmt.Println("iiiiiiiiiiiiiiii", string(message))
	h.mu.RLock()
	defer h.mu.RUnlock()
	if userClients, ok := h.Clients[userID]; ok {
		for client := range userClients {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(userClients, client)
			}
		}
	}
}
