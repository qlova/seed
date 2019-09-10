package script

import qlova "github.com/qlova/script"

//Go calls a Go function with the provided arguments.
func (q Script) Go(function interface{}, args ...qlova.Type) Promise {
	var Promise = q.rpc(function, "undefined", nil, args...)
	q.Javascript(Promise.expression + `.then(function(response) {
	if (response.charAt(0) != "{") return;
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

	for (let namespace in json.Evaluations) {
		for (let instruction of json.Evaluations[namespace]) {
			eval(instruction);
		}
	}
}).catch(function(){});
	`)
	return Promise
}
