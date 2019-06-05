package script

import qlova "github.com/qlova/script"

func (q Script) Go(function interface{}, args ...qlova.Type) Promise {
	var Promise = q.rpc(function, "undefined", args...)
	q.Javascript(Promise.expression + `.then(function(response) {
	let json = JSON.parse(response);
	for (let update in json.Document) {
		if (update.charAt(0) == "#") {
			let splits = update.split(".", 2)
			let id = splits[0];
			let property = splits[1];
			console.log("get('"+id.substring(1)+"')."+property+" = '"+json.Document[update]+"';");
			eval("get('"+id.substring(1)+"')."+property+" = '"+json.Document[update]+"';");
		}
	}
	for (let update in json.LocalStorage) {
		window.localStorage.setItem(update, json.LocalStorage[update]);
	}
});
	`)
	return Promise
}
