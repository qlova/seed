package swiper

import "github.com/qlova/seed"

type Effect string
const Coverflow = "coverflow"

type CoverflowEffectOptions struct {
	Rotate float64 `json:"rotate,omitempty"`
	Stretch float64 `json:"stretch,omitempty"`
	Depth int `json:"depth,omitempty"`
	Modifier int `json:"modifier,omitempty"` 
	SlideShadows bool `json:"slideShadows"` 
}

type PaginationOptions struct {
	Element seed.Seed `json:"el,omitempty"` 
}

type Options struct {
	Effect Effect `json:"effect,omitempty"` 
	
	CoverflowEffect CoverflowEffectOptions `json:"coverflowEffect,omitempty"` 
	
	Pagination PaginationOptions `json:"pagination,omitempty"` 
} 
