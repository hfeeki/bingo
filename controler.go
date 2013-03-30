package bingo

import (
	"fmt"
)

type MyResourceHandler struct {
	ResourceHandler
}

type ResourceController interface {
	Show(Env) (Status, Headers, Body)
	Index(Env) (Status, Headers, Body)
	Create(Env) (Status, Headers, Body)
	Update(Env) (Status, Headers, Body)
	Delete(Env) (Status, Headers, Body)
}

func (self *MyResourceHandler) AddResource(name string, c ResourceController) error {
	show_func := func(e Env) (s Status, h Headers, b Body) {
		s, h, b = c.Show(e)
		return s, h, b
	}
	index_func := func(e Env) (s Status, h Headers, b Body) {
		return c.Index(e)
	}
	create_func := func(e Env) (s Status, h Headers, b Body) {
		return c.Create(e)
	}
	update_func := func(e Env) (s Status, h Headers, b Body) {
		return c.Update(e)
	}
	delete_func := func(e Env) (Status, Headers, Body) {
		return c.Delete(e)
	}

	err := self.ResourceHandler.SetRoutes(
		Route{"GET", fmt.Sprintf("/%s/:id", name), show_func},
		Route{"GET", fmt.Sprintf("/%s", name), index_func},
		Route{"POST", fmt.Sprintf("/%s", name), create_func},
		Route{"PUT", fmt.Sprintf("/%s/:id", name), update_func},
		Route{"DELETE", fmt.Sprintf("/%s/:id", name), delete_func},
	)
	return err
}
