package script

//Location is a type that contains geolocation data.
type Location struct {
	Q Ctx
	Native
}

//Latitude of the location.
func (loc Location) Latitude() Float {
	return loc.Q.Value("%v.coord.latitude", loc).Float()
}

//Longitude of the location.
func (loc Location) Longitude() Float {
	return loc.Q.Value("%v.coord.longitude", loc).Float()
}

//GeoLocation is the JS required for GeoLocation support.
const GeoLocation = `
	var geoLocation = null;
	var requestGeoLocation = function (options) {
		return new Promise(function (resolve, reject) {
			navigator.geolocation.getCurrentPosition(function(position) {
				geoLocation = position;
				resolve(position);
			}, reject, options);
		});
	}
`

//RequestGeoLocation requests GeoLocation information.
//This must be called before q.GeoLocation is called.
func (q Ctx) RequestGeoLocation() Promise {
	q.Require(GeoLocation)
	var variable = q.Value(`requestGeoLocation()`).Native().Var()
	return Promise{
		variable, q,
	}
}

//GeoLocation returns the current Location.
func (q Ctx) GeoLocation() Location {
	q.Require(GeoLocation)
	return Location{
		q,
		q.Value("geoLocation").Native(),
	}
}
