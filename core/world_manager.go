package core

import "sync"

/**
 * 当前世界的管理模块
 */
type WorldManager struct {
	AOIManager *AOIManager
	Players map[int32] *Player
	pLock sync.RWMutex
}

// 提供一个对外的世界管理模块句柄（全局）
var WorldManagerObj *WorldManager

func init() {
	WorldManagerObj = &WorldManager{
		// 创建世界AOI地图规划
		AOIManager: NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_X, AOI_MAX_Y, AOI_CNTS_Y),
		// 初始化Plaers集合
		Players:make(map[int32]*Player),
	}
}

func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Players[player.Pid] = player
	wm.pLock.Unlock()

	// 将player添加到AOIManager中
	wm.AOIManager.AddPidToGridByPos(int(player.Pid), player.X, player.Z)
}

func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	// 将player添加到AOIManager中
	player := wm.Players[pid]
	wm.AOIManager.RemovePidFromGridByPos(int(pid), player.X, player.Z)

	wm.pLock.Lock()
	delete(wm.Players, pid)
	wm.pLock.Unlock()
}

func (wm *WorldManager) GetPlayerByPid(pid int32) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	return wm.Players[pid]
}

func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	players := make([]*Player,0)

	for _, p := range wm.Players {
		players = append(players, p)
	}

	return players
}

