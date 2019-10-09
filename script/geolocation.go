package script

//Location is a type that contains geolocation data.
type Location struct {
	Native
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
	return Promise{
		q.Value(`requestGeoLocation()`).Native(), q,
	}
}

//GeoLocation returns the current Location.
func (q Ctx) GeoLocation() Location {
	q.Require(GeoLocation)
	return Location{
		q.Value("geoLocation").Native(),
	}
}
