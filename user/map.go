package user

import (
	"errors"
	"sync"
)

var active = make(map[string]Ctx)
var mutex sync.RWMutex

//Save saves the user-context under the given key which can be retrieved at a later time with Load.
func Save(u Ctx, name string) error {
	mutex.Lock()
	defer mutex.Unlock()

	if u.conn == nil {
		return errors.New("cannot save: user.Ctx wasn't created with CtxFromSocket")
	}

	active[name] = u

	return nil
}

//Load loads a user under the given name.
func Load(name string) (Ctx, error) {
	mutex.RLock()
	defer mutex.RUnlock()

	u, ok := active[name]
	if !ok {
		return u, errors.New("user is not connected")
	}

	return u, nil
}
