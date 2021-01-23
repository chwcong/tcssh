package handler

import "github.com/desertbit/grumble"

type Handler interface {
	Handle (*grumble.Context) error
}