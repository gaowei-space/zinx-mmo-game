package core

import (
	"fmt"
	"testing"
)

// func TestNewAOIManager(t *testing.T) {
// 	aoiMgr := NewAOIManager(0, 250, 5, 0, 250, 5)
// 	fmt.Println(aoiMgr)
// }

func TestGetSurroundGridsByGid(t *testing.T) {
	aoiMgr := NewAOIManager(0, 250, 5, 0, 250, 5)

	for gid, _ := range aoiMgr.Grids {
		// 得到当前九宫格的所有格子信息
		grids := aoiMgr.GetSurroundGridsByGid(gid)
		fmt.Println("gid:", gid, " grids len:", len(grids))
		gIDs := make([]int, 0, len(grids))
		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}
		fmt.Printf("grid ID: %d, surrounding grid IDs are %v\n", gid, gIDs)
	}
}
