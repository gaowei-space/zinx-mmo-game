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
		Grids: make(map[int]*Grid),
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

// 根据格子GID得到周边格子的集合
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	// 判断gID是否再AOIManager中
	if _, ok := m.Grids[gID]; !ok {
		return
	}

	// 初始化grids返回值切片, 将当前gID格子加入九宫格中
	grids = append(grids, m.Grids[gID])

	// 需要判断gID左边和右边是否有格子

	// 需要通过gID得到当前格子的x轴编号 idx = id % nx
	idx := gID % m.CntsX

	// 判断idx编号左边是否还有格子，如果有 放入gridX集合中
	if idx > 0 {
		grids = append(grids, m.Grids[gID-1])
	}
	// 判断idx编号右边是否还有格子，如果有 放入gridX集合中
	if idx < m.CntsX-1 {
		grids = append(grids, m.Grids[gID+1])
	}

	// 将x轴当前的格子都取出，进行遍历，再分别得到每个格子上下是否还有格子
	// 得到当前x轴格子的id集合
	gidsX := make([]int, 0, len(grids))

	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	// 遍历gidsX集合中每个格子的gid，确认上下是否还有格子
	for _, v := range gidsX {
		// 得到当前格子在Y轴的编号 idy = id / ny
		idy := v / m.CntsY

		if idy > 0 {
			grids = append(grids, m.Grids[v-m.CntsX])
		}

		if idy < m.CntsY-1 {
			grids = append(grids, m.Grids[v+m.CntsX])
		}
	}

	return
}

// 通过横纵坐标得到格子id
func (m *AOIManager) GetGidByPos(x, y float32) int {
	var idx, idy int
	idx = (int(x) - m.MinX) / m.gridWidth()
	idy = (int(y) - m.MinY) / m.gridHeight()

	return idy*m.CntsX + idx
}

// 通过横纵坐标得到周边所有玩家ids
func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	// 得到当前玩家的格子GID
	gID := m.GetGidByPos(x, y)
	// 再通过gid得到周边格子集合
	grids := m.GetSurroundGridsByGid(gID)
	// 将周边格子中的玩家累家到playerIDs
	for _, v := range grids {
		playerIDs = append(playerIDs, v.GetPlayers()...)
	}
	return
}

// 添加一个PlayerID到一个格子中
func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.Grids[gID].Add(pID)
}

// 移除一个格子中的PlayerID
func (m *AOIManager) RemovePidFromGrid(pID, gID int) {
	m.Grids[gID].Remove(pID)
}

// 通过 GID 获取全部的playerID
func (m *AOIManager) GetPidsByGid(gID int) (playerIDs []int) {
	playerIDs = m.Grids[gID].GetPlayers()
	return
}

// 通过 坐标把一个playerID加入某个格子中
func (m *AOIManager) AddPidToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	m.Grids[gID].Add(pID)
}

// 通过 坐标从一个格子中删除一个playerID
func (m *AOIManager) RemovePidFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	m.Grids[gID].Remove(pID)
}
