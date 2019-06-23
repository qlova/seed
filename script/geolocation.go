package script

type Location struct {
	Native
}

func (q Script) RequestGeoLocation() Promise {
	return Promise{
		`requestGeoLocation()`, q,
	}
}

func (q Script) GeoLocation() Location {
	return Location{
		q.Value("geoLocation").Native(),
	}
}
