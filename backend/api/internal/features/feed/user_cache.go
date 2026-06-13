package feed

import (
	"Server/internal/features/users"
	"sync"
)

type userCache struct {
	sync.Mutex
	users map[string]*users.User
}

func (uc *userCache) Get(id string, fetch func(string) (*users.User, error)) (*users.User, error) {
	uc.Lock()
	defer uc.Unlock()
	if u, ok := uc.users[id]; ok {
		return u, nil
	}
	u, err := fetch(id)
	if err != nil {
		return nil, err
	}
	uc.users[id] = u
	return u, nil
}
