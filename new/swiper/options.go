package swiper

import "qlova.org/seed"

type PaginationOptions struct {
	Element string `json:"el,omitempty"`
}

type Options struct {
	Effect string `json:"effect,omitempty"`

	CoverflowEffect CoverflowEffect `json:"coverflowEffect,omitempty"`

	Pagination PaginationOptions `json:"pagination,omitempty"`

	Observer             bool `json:"observer,omitempty"`
	ObserveParents       bool `json:"observeParents,omitempty"`
	ObserveSlideChildren bool `json:"observeSlideChildren,omitempty"`
}

type CoverflowEffect struct {
	Rotate       float64 `json:"rotate"`
	Stretch      float64 `json:"stretch,omitempty"`
	Depth        int     `json:"depth"`
	Modifier     int     `json:"modifier,omitempty"`
	SlideShadows bool    `json:"slideShadows"`
}

func (CoverflowEffect) AddTo(c seed.Seed) {}
