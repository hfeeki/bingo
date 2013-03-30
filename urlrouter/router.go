// Efficient URL routing using a Trie data structure.
//
// This Package implements a URL Router, but instead of using the usual
// "evaluate all the routes and return the first regexp that matches" strategy,
// it uses a Trie data structure to perform the routing. This is more efficient,
// and scales better for a large number of routes.
// It supports the :param and *splat placeholders in the route strings.
//
// Example:
//	router := urlrouter.Router{
//		Routes: []urlrouter.Route{
//			urlrouter.Route{
//				PathExp: "/resources/:id",
//				Dest:    "one_resource",
//			},
//			urlrouter.Route{
//				PathExp: "/resources",
//				Dest:    "all_resources",
//			},
//		},
//	}
//
//	err := router.Start()
//	if err != nil {
//		panic(err)
//	}
//
//	input := "http://example.org/resources/123"
//	route, params, err := router.FindRoute(input)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Print(route.Dest)  // one_resource
//	fmt.Print(params["id"])  // 123
//
// (Blog Post: http://blog.ant0ine.com/typepad/2013/02/better-url-routing-golang-1.html)
package urlrouter

import (
	"net/url"
)

// TODO
// support for http method routing ?
// support for #param placeholder ?
// replace map[string]string by a PathParams object for more flexibility, and support for multiple param with the same names, eg: /users/:id/blogs/:id ?

type Route struct {
	// A string defining the route, like "/resource/:id.json".
	// Placeholders supported are:
	// :param that matches any char to the first '/' or '.'
	// *splat that matches everything to the end of the string
	Method  string
	Pattern string
	// Can be anything useful to point to the code to run for this route.
	Action interface{}
}

type Router interface {
	AddRoutes(routes ...Route) error
	// AddResource(name string, c interface{}) error
	FindRouteFromURL(method string, url_obj *url.URL) (*Route, map[string]string)
	FindRoute(method, url_str string) (*Route, map[string]string, error)
}
