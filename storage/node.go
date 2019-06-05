package storage

//TODO
/*
	FORLOOP

	var message Message
	for element := Messages.Range(); !element.End(); element.Next() {
		element.Update(message)

		message.Email
	}

*/

//A storage.Node can be thought of as a database or directory.
type Node interface {
	//Create a view.
	Create(view View) bool

	//Insert data into the node, returns an id, nil if unsuccessful.
	Put(view View, data []byte) []byte

	//Set the data at the specified key. Returns success state.
	Set(view View, key []byte, data []byte) bool

	//Get the data at the specified key, this will copy the data into a slice that is safe to pass around.
	//Nil if not found.
	Get(view View, key []byte) (data []byte)

	//Access the data at key, this is performance optimised. The handler will not run if the key is invalid.
	Read(view View, key []byte, handler func(data []byte))

	ForEach(view View, function func(key []byte, data []byte))
}
