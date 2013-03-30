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
	"errors"
	"github.com/hfeeki/bingo/urlrouter/trie"
	"net/url"
	"regexp"
	"sort"
	"strings"
)

// Methods required by sort.Interface.
type matcherArray []*regexp.Regexp

func specificity(matcher *regexp.Regexp) int {
	return len(matcher.String())
}
func (this matcherArray) Len() int {
	return len(this)
}
func (this matcherArray) Less(i, j int) bool {
	// The sign is reversed below so we sort the matchers in descending order
	return specificity(this[i]) > specificity(this[j])
}
func (this matcherArray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type RegxRouter struct {
	// list of Routes, the order matters, if multiple Routes match, the first defined will be used.
	Routes   []Route
	patterns map[string]string
	matchers matcherArray
	handlers []interface{}
}

// Define the Routes. The order the Routes matters,
// if a request matches multiple Routes, the first one will be used.
// Note that the underlying router is https://github.com/ant0ine/go-urlrouter.
func (self *RegxRouter) AddRoutes(routes ...Route) error {

	for _, route := range routes {
		self.Routes = append(
			self.Routes,
			route,
		)
	}

	return nil
}

// This validates the Routes and prepares the Trie data structure.
// It must be called once the Routes are defined and before trying to find Routes.
func (self *RegxRouter) Start() error {

	// Compile the matchers
	for _, route := range self.Routes {
		pattern := strings.ToUpper(route.Method) + route.Pattern
		// compile the matchers
		self.patterns[pattern] = route
	}

	// Compile the matchers
	for matcher, _ := range self.patterns {
		// compile the matchers
		self.matchers = append([]*regexp.Regexp(self.matchers), regexp.MustCompile(matcher))
	}

	// sort 'em by descending length
	sort.Sort(self.matchers)

	// TODO validation of the PathExp ? start with a /
	// TODO url encoding

	return nil
}

// Return the first matching Route and the corresponding parameters for a given URL object.
func (self *RegxRouter) FindRouteFromURL(method string, url_obj *url.URL) (*Route, map[string]string) {

	for i, matcher := range self.matchers {
		matches := matcher.FindStringSubmatch(method + url_obj.Path)
		if len(matches) != 0 {
			// Matched a route; inject matches and return handler			
			// only return the first Route that matches			

			for _, match := range matches {
				pattern := match.String()
				route := self.patterns[pattern]
			}
		}
	}

	// only return the first Route that matches
	min_index := -1
	matches_by_index := map[int]*trie.Match{}

	for _, match := range matches {
		route := match.Route.(*Route)
		route_index := self.index[route]
		matches_by_index[route_index] = match
		if min_index == -1 || route_index < min_index {
			min_index = route_index
		}
	}

	if min_index == -1 {
		// no route found
		return nil, nil
	}

	// and the corresponding params
	match := matches_by_index[min_index]

	return match.Route.(*Route), match.Params
}

// Parse the url string (complete or just the path) and return the first matching Route and the corresponding parameters.
func (self *RegxRouter) FindRoute(method, url_str string) (*Route, map[string]string, error) {

	// parse the url
	url_obj, err := url.Parse(url_str)
	if err != nil {
		return nil, nil, err
	}

	route, params := self.FindRouteFromURL(url_obj)
	return route, params, nil
}
