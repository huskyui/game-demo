package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	GID int

	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	playerIDs map[int]bool
	pIDLock   sync.RWMutex
}

func NewGrid(gId, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gId,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

func (g *Grid) Add(playerId int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.playerIDs[playerId] = true
}

func (g *Grid) Remove(playerId int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerId)
}

func (g *Grid) GetPlayerIds() (playerIDs []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.Unlock()
	for k, _ := range g.playerIDs {
		playerIDs = append(playerIDs, k)
	}
	return playerIDs
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid id: %d minX: %d maxX :%d minY :%d maxY: %d playerIDS %v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)
}
