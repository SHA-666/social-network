package ws

import (
	"encoding/json"
	"log"
	"sync"
)

// OK pour 10-50k connexions
// sharder les hubs (ex : avec Redis Pub/Sub ou NATS) pour synchroniser les événements entre plusieurs instances backend.
type Hub struct {
	Clients    map[string]map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	stop       chan struct{}
	mu         sync.RWMutex
}

func InitHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]map[*Client]bool),
		stop:       make(chan struct{}),
	}
}

func (h *Hub) DisconnectUser(userId string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if userClients, ok := h.Clients[userId]; ok {
		log.Printf("🔌 Déconnexion de l'utilisateur %s (%d connexions)", userId, len(userClients))
		for client := range userClients {
			select {
			case client.Send <- []byte("logout"):
			default:
				close(client.Send)
			}
			client.Conn.Close()
		}
		delete(h.Clients, userId)
		log.Printf("✅ Utilisateur %s déconnecté du Hub", userId)
	}
}

func (h *Hub) Run() {
	log.Printf("🚀 Hub démarré")

	for {
		select {
		case <-h.stop:
			log.Println("🛑 Hub en cours d'arrêt...")
			h.mu.Lock()
			for _, clients := range h.Clients {
				for client := range clients {
					close(client.Send)
				}
			}
			h.mu.Unlock()
			return

		case client := <-h.Register:
			h.mu.Lock()
			if h.Clients[client.UserId] == nil {
				h.Clients[client.UserId] = make(map[*Client]bool)
			}
			h.Clients[client.UserId][client] = true
			connectionsCount := len(h.Clients[client.UserId])
			h.mu.Unlock()

			log.Printf("✅ Client enregistré pour %s (Total: %d) ,,,,,,", client.UserId, connectionsCount, client.ClientId)
		case client := <-h.Unregister:
			h.mu.Lock()
			if userClients, ok := h.Clients[client.UserId]; ok {
				if _, exists := userClients[client]; exists {
					delete(userClients, client)
					close(client.Send)
					log.Printf("❌ Client désenregistré pour %s (Restant: %d)", client.UserId, len(userClients))
				}

				if len(userClients) == 0 {
					delete(h.Clients, client.UserId)
					log.Printf("⚠️ Plus de connexions pour %s", client.UserId)
				}
			}
			h.mu.Unlock()
		case message := <-h.Broadcast:
			h.mu.RLock()
			// id _
			for _, userClients := range h.Clients {
				for client := range userClients {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(userClients, client)
					}
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) Stop() {
	close(h.stop)
}

func (h *Hub) BroadcastToUser(userID string, message any) {
    data, err := json.Marshal(message)
    if err != nil {
        log.Printf("Erreur marshalling message: %v", err)
        return
    }  
    h.mu.RLock()
    defer h.mu.RUnlock()   
    if clients, exists := h.Clients[userID]; exists {
        for client := range clients {
            select {
            case client.Send <- data:
            default:
                close(client.Send)
                delete(clients, client)
                if len(clients) == 0 {
                    delete(h.Clients, userID)
                }
            }
        }
    }
}
