package core

import "fmt"

// AOI 管理模块
type AOIManager struct {
	MinX  int // 区域的左边界坐标
	MaxX  int // 区域的右边界坐标
	CntsX int // X方向格子数量
	MinY  int // 区域的上边界坐标
	MaxY  int // 区域的下边界坐标
	CntsY int // Y方向格子数量

	// 当前区域有哪些格子 map: key 格子ID，value 格子对象
	Grids map[int]*Grid
}

// 初始化 AOI 区域管理模块
func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiManager := &AOIManager{
		MinX:  minX,  // 区域的左边界坐标
		MaxX:  maxX,  // 区域的右边界坐标
		CntsX: cntsX, // X方向格子数量
		MinY:  minY,  // 区域的上边界坐标
		MaxY:  maxY,  // 区域的下边界坐标
		CntsY: cntsY, // Y方向格子数量
	}

	// 初始化区域中所有格子
	for y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			// 格子编号：id=idY * cntX + idX
			gid := y*cntsX + x

			// 初始化格子
			aoiManager.Grids[gid] = NewGrid(
				gid,
				aoiManager.MinX+x*aoiManager.gridWidth(),
				aoiManager.MinX+(x+1)*aoiManager.gridWidth(),
				aoiManager.MinY+y*aoiManager.gridHeight(),
				aoiManager.MinY+(y+1)*aoiManager.gridHeight(),
			)
		}
	}

	return aoiManager
}

// 得到每个格子在X轴方向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}

// 得到每个格子在Y轴方向的高度
func (m *AOIManager) gridHeight() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

// 打印AOI信息
func (m *AOIManager) String() string {
	// 打印AOI信息
	s := fmt.Sprintf("AOIManager: \n minX: %d, maxX: %d, cntsX: %d, minY: %d, maxY: %d, cntsY: %d \n", m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)

	for _, grid := range m.Grids {
		// 打印AOI中每个格子的信息
		s += fmt.Sprintln(grid)
	}

	return s
}
