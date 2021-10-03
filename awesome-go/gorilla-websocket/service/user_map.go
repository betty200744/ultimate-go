package main

import (
	"sync"
)

type UserMap struct {
	lock  sync.RWMutex
	users map[int64]*User
}

func NewUserMap() *UserMap {
	return &UserMap{users: map[int64]*User{}}
}

func (m *UserMap) set(user *User) bool {
	if user == nil || user.id == 0 {
		return false
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	m.users[user.id] = user
	return true
}

func (m *UserMap) get(id int64) (*User, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.users[id]
	return v, ok
}

func (m *UserMap) del(id int64) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.users, id)
}

func (m *UserMap) getAll() []*User {
	m.lock.RLock()
	defer m.lock.RUnlock()
	out := make([]*User, 0, len(m.users))
	for _, user := range m.users {
		out = append(out, user)
	}
	return out
}
