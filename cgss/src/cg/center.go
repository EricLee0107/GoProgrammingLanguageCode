package cg

import (
	"encoding/json"
	"errors"
	"ipc"
	"sync"
)

var _ ipc.Server = &CenterServer{} // 确认实现了Server接口

type Message struct {
	From    string "from"
	To      string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string]ipc.Server
	player  []*Player
	rooms   []*Room
	mutex   sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	player := make([]*Player, 0)
	return &CenterServer{servers: servers, players: players}
}

func (server *CenterServer) addPlayer(params string) error {
	player := NewPlayer()
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()

	//偷懒了没有重复登录验证
	server.player = append(server.players, player)
	return nil
}

func (server *CenterServer) removePlayer(params string) error {

}
