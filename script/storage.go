package script

import (
	qlova "github.com/qlova/script"
	"github.com/qlova/seed/user"
)

//GetCookie is the required JS code for getting cookies.
const GetCookie = `
	function getCookie(cname) {
		var name = cname + "=";
		var decodedCookie = decodeURIComponent(document.cookie);
		var ca = decodedCookie.split(';');
		for(var i = 0; i <ca.length; i++) {
		var c = ca[i];
		while (c.charAt(0) == ' ') {
			c = c.substring(1);
		}
		if (c.indexOf(name) == 0) {
			return c.substring(name.length, c.length);
		}
		}
		return "";
	}
`

//UserData retuns the cookie with the given name.
func (q Script) UserData(name user.Data) qlova.String {
	q.Require(GetCookie)
	return q.wrap(`getCookie("` + string(name) + `")`)
}

//SetCookie is the required code for setting cookies.
const SetCookie = `
	function setCookie(cname, cvalue, exdays) {
		var d = new Date();
		d.setTime(d.getTime() + (exdays*24*60*60*1000));
		var expires = "expires="+ d.toUTCString();
		if (production) {
			document.cookie = cname + "=" + cvalue + ";" + expires + ";secure;path=/";
		} else {
			document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
		}
	}
`

//SetUserData sets the cookie with the given name to the given value.
func (q Script) SetUserData(name user.Data, value qlova.String) {
	q.Require(SetCookie)
	q.Javascript(`setCookie("` + string(name) + `", ` + raw(value) + `, 365);`)
}
