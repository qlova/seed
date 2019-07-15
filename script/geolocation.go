package script

type Location struct {
	Native
}

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

func (q Script) RequestGeoLocation() Promise {
	q.Require(GeoLocation)
	return Promise{
		`requestGeoLocation()`, q,
	}
}

func (q Script) GeoLocation() Location {
	q.Require(GeoLocation)
	return Location{
		q.Value("geoLocation").Native(),
	}
}
