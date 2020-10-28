package api

import (
	"net/http"

	"qlova.org/seed"
)

type harvester struct {
	Map map[string]http.Handler
}

func newHarvester() harvester {
	return harvester{
		Map: make(map[string]http.Handler),
	}
}

func (h harvester) harvest(c seed.Seed) map[string]http.Handler {
	var data data
	c.Load(&data)

	for route, handler := range data.handlers {
		h.Map[route] = handler
	}

	for _, child := range c.Children() {
		h.harvest(child)
	}

	return h.Map
}

func Routes(root seed.Seed) map[string]http.Handler {
	return newHarvester().harvest(root)
}
