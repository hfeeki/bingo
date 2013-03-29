package bingo

import (
	"fmt"
	"github.com/ant0ine/go-json-rest"
)

type MyResourceHandler struct {
	rest.ResourceHandler
}

type ResourceController interface {
	Show(w *rest.ResponseWriter, req *rest.Request)
	Create(w *rest.ResponseWriter, req *rest.Request)
	Update(w *rest.ResponseWriter, req *rest.Request)
	Delete(w *rest.ResponseWriter, req *rest.Request)
}

func (self *MyResourceHandler) AddResource(name string, c ResourceController) error {
	show_func := func(w *rest.ResponseWriter, r *rest.Request) {
		c.Show(w, r)
	}
	create_func := func(w *rest.ResponseWriter, r *rest.Request) {
		c.Create(w, r)
	}
	update_func := func(w *rest.ResponseWriter, r *rest.Request) {
		c.Update(w, r)
	}
	delete_func := func(w *rest.ResponseWriter, r *rest.Request) {
		c.Delete(w, r)
	}

	err := self.ResourceHandler.SetRoutes(
		rest.Route{"GET", fmt.Sprintf("/%s/:id", name), show_func},
		rest.Route{"POST", fmt.Sprintf("/%s", name), create_func},
		rest.Route{"PUT", fmt.Sprintf("/%s/:id", name), update_func},
		rest.Route{"DELETE", fmt.Sprintf("/%s/:id", name), delete_func},
	)
	return err
}
