package storage

type Map map[string]string

func NewMap() Map {
	return make(Map)
}

//Create a view.
func (m Map) Create(view View) bool {
	return false
}

//Insert data into the node, returns an id, nil if unsuccessful.
func (m Map) Put(view View, data []byte) []byte {
	return nil
}

//Set the data at the specified key. Returns success state.
func (m Map) Set(view View, key []byte, data []byte) bool {
	return false
}

//Get the data at the specified key, this will copy the data into a slice that is safe to pass around.
//Nil if not found.
func (m Map) Get(view View, key []byte) (data []byte) {
	return nil
}

//Access the data at key, this is performance optimised. The handler will not run if the key is invalid.
func (m Map) Read(view View, key []byte, handler func(data []byte)) {
	
}

func (m Map) ForEach(view View, function func(key []byte, data []byte)) {
	
}