package resource

import (
	"common/databases/mysql"
	"common/databases/redis"
)

type Manager struct {
	Redis *redis.RedisManager
	Db    *mysql.Db
}

func NewManager() *Manager {
	return &Manager{
		Redis: redis.NewRedis(),
		Db:    mysql.NewDb(),
	}
}

func (m *Manager) Close() {
	if m.Redis != nil {
		m.Redis.Close()
	}
}
