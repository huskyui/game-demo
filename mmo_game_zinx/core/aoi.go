package core

import "fmt"

type AOIManager struct {
	// 区域左边的坐标
	MinX int
	// 区域右边的坐标
	MaxX int
	// x轴方向的格子数量
	CntsX int
	MinY  int
	MaxY  int
	CntsY int
	// 当前格子的id,和格子对象
	grids map[int]*Grid
}

func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	manager := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int]*Grid),
	}

	gridWidth := manager.gridWidth()
	gridHeight := manager.gridHeight()

	// 初始化AOI格子
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			gid := y*cntsX + x
			manager.grids[gid] = NewGrid(gid,
				manager.MinX+x*gridWidth,
				manager.MinX+(x+1)*gridWidth,
				manager.MinY+y*gridHeight,
				manager.MinY+(y+1)*gridHeight)
		}
	}

	return manager
}

func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

func (m *AOIManager) gridHeight() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

func (m *AOIManager) String() string {
	s := fmt.Sprintf("AOIManager： \n MinX: %d,MaxX:%d,cntsX:%d,miny:%d,maxY:%d,cntsY:%d\n", m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}

func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	if _, ok := m.grids[gID]; !ok {
		return
	}
	grids = append(grids, m.grids[gID])
	idx := gID % m.CntsX
	if idx > 0 {
		grids = append(grids, m.grids[gID-1])
	}
	if idx < m.MaxX-1 {
		grids = append(grids, m.grids[gID+1])
	}

	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	for _, v := range gidsX {
		idy := v / m.CntsY
		if idy > 0 {
			grids = append(grids, m.grids[v-m.CntsX])
		}
		if idy < m.CntsY-1 {
			grids = append(grids, m.grids[v+m.CntsX])
		}
	}
	return grids

}
