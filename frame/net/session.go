package net

import "sync"

type Session struct {
	sync.RWMutex
	Uid  int
	Cid  string
	data map[string]any
	all  map[string]any
}

func NewSession(cid string) *Session {
	return &Session{
		Cid:  cid,
		all:  make(map[string]any),
		data: make(map[string]any),
	}
}

func (s *Session) Put(k string, v any) {
	s.Lock()
	defer s.Unlock()
	s.data[k] = v
}

func (s *Session) Get(k string) (v any, ok bool) {
	s.RLock()
	s.RUnlock()
	v, ok = s.data[k]
	return
}

func (s *Session) SetData(uid int, data map[string]any) {
	s.Lock()
	defer s.Unlock()

	if uid == s.Uid {
		for k, v := range data {
			s.data[k] = v
		}
	}
}
