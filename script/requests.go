package script

import "github.com/qlova/script/language"

//DataRequestJS is the js defined data_request function.
const DataRequestJS = `
function data_request (url) {
	if (url.charAt(0) == "/") url = host+url;
	return new Promise(function (resolve, reject) {
		var xhr = new XMLHttpRequest();
		xhr.open("GET", url, true);
		xhr.responseType = 'json';
		xhr.onload = function () {
			if (this.status >= 200 && this.status < 300) {
				resolve(xhr.response);
			} else {
				reject({
					status: this.status,
					statusText: xhr.statusText,
					response: xhr.response
				});
			}
		};
		xhr.onerror = function () {
			reject({
				status: this.status,
				statusText: xhr.statusText,
				response: xhr.response
			});
		};
		xhr.send();
	});
}
`

//DataResponse returns the resulting value from the specified key inside of the data result object.
func (q Ctx) DataResponse(key string) String {
	return q.Value(`rpc_result["` + key + `"]`).String()
}

//DataRequest makes a new request execting JSON data from the specified URL.
func (q Ctx) DataRequest(url String) Promise {
	q.Require(DataRequestJS)

	var variable = Unique()

	q.Raw("Javascript", language.Statement(`let `+variable+` = data_request(`+url.LanguageType().Raw()+`);`))

	return Promise{q.Value(variable).Native(), q}
}
