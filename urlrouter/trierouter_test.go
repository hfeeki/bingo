package urlrouter

import (
	"net/url"
	"testing"
)

func TestFindRouteAPI(t *testing.T) {

	router := TrieRouter{
		Routes: []Route{
			Route{
				Method:  "GET",
				Pattern: "/",
				Action:  "root",
			},
		},
	}

	err := router.Start()
	if err != nil {
		t.Fatal()
	}

	// net/url can parse GET/abc/xxx as url
	// full url string
	input := "http://example.org/"
	route, params, err := router.FindRoute("GET", input)
	if err != nil {
		t.Fatal()
	}
	// println(route)
	// println(params)
	if route.Action != "root" {
		t.Error()
	}
	if len(params) != 0 {
		t.Error()
	}

	// part of the url string
	input = "http://example.org/"
	route, params, err = router.FindRoute("GET", input)
	if err != nil {
		t.Fatal()
	}
	if route.Action != "root" {
		t.Error()
	}
	if len(params) != 0 {
		t.Error()
	}

	// url object
	url_obj, err := url.Parse("http://example.org/")
	if err != nil {
		t.Fatal()
	}
	route, params = router.FindRouteFromURL("GET", url_obj)
	if route.Action != "root" {
		t.Error()
	}
	if len(params) != 0 {
		t.Error()
	}
}

func TestNoRoute(t *testing.T) {

	router := TrieRouter{
		Routes: []Route{},
	}

	err := router.Start()
	if err != nil {
		t.Fatal()
	}

	input := "http://example.org/notfound"
	route, params, err := router.FindRoute("GET", input)
	if err != nil {
		t.Fatal()
	}

	if route != nil {
		t.Error("should not be able to find a route")
	}
	if params != nil {
		t.Error("params must be nil too")
	}
}

func TestDuplicatedRoute(t *testing.T) {

	router := TrieRouter{
		Routes: []Route{
			Route{
				Method:  "GET",
				Pattern: "/",
				Action:  "root",
			},
			Route{
				Method:  "GET",
				Pattern: "/",
				Action:  "the_same",
			},
		},
	}

	err := router.Start()
	if err == nil {
		t.Error("expected the duplicated route error")
	}
}

func TestRouteOrder(t *testing.T) {

	router := TrieRouter{
		Routes: []Route{
			Route{
				Method:  "GET",
				Pattern: "/r/:id",
				Action:  "first",
			},
			Route{
				Method:  "GET",
				Pattern: "/r/*rest",
				Action:  "second",
			},
		},
	}

	err := router.Start()
	if err != nil {
		t.Fatal()
	}

	input := "http://example.org/r/123"
	route, params, err := router.FindRoute("GET", input)
	if err != nil {
		t.Fatal()
	}

	if route.Action != "first" {
		t.Errorf("both match, expected the first defined, got %s", route.Action)
	}
	if params["id"] != "123" {
		t.Error()
	}
}

func TestSimpleExample(t *testing.T) {

	router := TrieRouter{
		Routes: []Route{
			Route{
				Method:  "GET",
				Pattern: "/resources/:id",
				Action:  "one_resource",
			},
			Route{
				Method:  "GET",
				Pattern: "/resources",
				Action:  "all_resources",
			},
		},
	}

	err := router.Start()
	if err != nil {
		t.Fatal()
	}

	input := "http://example.org/resources/123"
	route, params, err := router.FindRoute("GET", input)
	if err != nil {
		t.Fatal()
	}

	if route.Action != "one_resource" {
		t.Error()
	}
	if params["id"] != "123" {
		t.Error()
	}
}
