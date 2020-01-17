package user

import "sync"

var active = make(map[string]User)
var mutex sync.RWMutex

//Save saves the user under the given name which can be loaded at a later time with Load.
func Save(u User, name string) {
	mutex.Lock()
	defer mutex.Unlock()

	active[name] = u
}

//Load loads a user under the given name.
func Load(name string) User {
	mutex.RLock()
	defer mutex.RUnlock()

	return active[name]
}
