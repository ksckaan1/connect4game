package httpapi

import (
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

func TestRoomHttpAPI_JoinRoom(t *testing.T) {
	router := chi.NewRouter()
	roomAPI := NewRoomHttpAPI().WithPattern("/room")
	router.Mount(roomAPI.App())
	server := httptest.NewServer(router)
	defer server.Close()

	t.Log(server.URL)

	uri, err := url.Parse(server.URL)
	require.NoError(t, err)
	uri.Scheme = "ws"
	uri.Path = "/room/join"

	conn, _, err := websocket.DefaultDialer.Dial(uri.String(), nil)
	require.NoError(t, err)
	defer conn.Close()

	resp := map[string]string{}
	err = conn.ReadJSON(&resp)
	require.NoError(t, err)
	t.Log(resp)
}
