package httpapi

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
)

type RoomHttpAPI struct {
	pattern  string
	handler  *chi.Mux
	upgrader *websocket.Upgrader
}

func NewRoomHttpAPI() *RoomHttpAPI {
	return &RoomHttpAPI{
		pattern:  "/room",
		handler:  chi.NewRouter(),
		upgrader: &websocket.Upgrader{},
	}
}

func (h *RoomHttpAPI) WithPattern(p string) *RoomHttpAPI {
	h.pattern = p
	return h
}

func (h *RoomHttpAPI) App() (string, http.Handler) {
	h.handler.Post("/create", h.createRoom)
	h.handler.HandleFunc("/join", h.JoinRoom)
	return h.pattern, h.handler
}

func (h *RoomHttpAPI) createRoom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("room created!"))
}

func (h *RoomHttpAPI) JoinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.WriteJSON(map[string]string{"hello": "world"})
	if err != nil {
		return
	}
}
