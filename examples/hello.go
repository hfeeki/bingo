package main

import (
	"fmt"
	"github.com/paulbellamy/mango"
	"reflect"
	//"github.com/robfig/revel"
)

type RouteEntry struct {
	method string
	path   string
	app    mango.App
}

type Routes map[string]RouteEntry

// var m = map[string]RouteEntry{
// 	"Bell Labs": Vertex{
// 		"GET", -74.39967,
// 	},
// 	"Google": Vertex{
// 		"GET", -122.08408,
// 	},
// }

// Data Model
type Dish struct {
	Id     int
	Name   string
	Origin string
	Query  func()
}

// Example of how to use Go's reflection
// Print the attributes of a Data Model
func attributes(m interface{}) map[string]reflect.Type {
	typ := reflect.TypeOf(m)
	// if a pointer to a struct is passed, get the type of the dereferenced object
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	// create an attribute data structure as a map of types keyed by a string.
	attrs := make(map[string]reflect.Type)
	// Only structs are supported so return an empty result if the passed object
	// isn't a struct
	if typ.Kind() != reflect.Struct {
		fmt.Printf("%v type can't have attributes inspected\n", typ.Kind())
		return attrs
	}

	// loop through the struct's fields and set the map
	for i := 0; i < typ.NumField(); i++ {
		p := typ.Field(i)
		if !p.Anonymous {
			attrs[p.Name] = p.Type
		}
	}

	return attrs
}

func Hello(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	r := env.Request()
	env.Logger().Println("Got a ", r.Method, " request for ", r.RequestURI)
	name := r.URL.Query().Get("name")
	body := fmt.Sprintf("Hello %s!", name)
	return 200, mango.Headers{}, mango.Body(body)
}

// Our handler for /goodbye
func Goodbye(env mango.Env) (mango.Status, mango.Headers, mango.Body) {
	return 200, mango.Headers{}, mango.Body("Goodbye World!")
}

func main() {
	stack := new(mango.Stack)
	stack.Address = ":3300"

	// Route all requests for /goodbye to the Goodbye handler
	routes := map[string]mango.App{"/goodbye(.*)": Goodbye}
	stack.Middleware(mango.Routing(routes))
	stack.Middleware(mango.ShowErrors("ERROR!!"))

	//stack.Run(Hello)
	fmt.Println("type of stack: ", reflect.TypeOf(stack))
	//fmt.Println("type of m: ", reflect.TypeOf(m))
	fmt.Println("type of hello: ", reflect.TypeOf(Hello))

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
