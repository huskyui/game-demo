package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManager(t *testing.T) {
	manager := NewAOIManager(0, 250, 5, 0, 250, 5)
	fmt.Println(manager)
}

func TestAOIManager_GetSurroundGridsByGid(t *testing.T) {
	aoiManager := NewAOIManager(0, 250, 5, 0, 250, 5)
	for gid, _ := range aoiManager.grids {
		grids := aoiManager.GetSurroundGridsByGid(gid)
		fmt.Println("gid:", gid, " grids len = ", len(grids))
		gidArr := make([]int, 0, len(grids))
		for _, grid := range grids {
			gidArr = append(gidArr, grid.GID)
		}
		fmt.Println("surrounding grid Ids are ", gidArr)
	}
}
