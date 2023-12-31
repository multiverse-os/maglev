package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

var _ Router = &Mux{}

// Mux is a simple HTTP route multiplexer that parses a request path,
// records any URL params, and executes an end handler. It implements
// the http.Handler interface and is friendly with the standard library.
//
// Mux is designed to be fast, minimal and offer a powerful API for building
// modular and composable HTTP services with a large set of handlers. It's
// particularly useful for writing large REST API services that break a handler
// into many smaller parts composed of middlewares and end handlers.

// TODO: Now maybe our App object should be stored inside here? or maybe like
// a meta-app that is outside of framework and combines the model/database,
// outputs (logging or output), etc

type Mux struct {
	parent *Mux
	// The radix trie router
	// TODO: Not like root ?
	tree *node

	// Controls the behaviour of middleware chain generation when a mux
	// is registered as an inline group inside another mux.
	// TODO: for real? wtf
	inline bool

	// The middleware stack
	middlewares []func(http.Handler) http.Handler
	// The computed mux handler made of the chained middleware stack and
	// the tree router
	// TODO: if this is stored ehre, thennn
	handler                 http.Handler
	notFoundHandler         http.HandlerFunc
	methodNotAllowedHandler http.HandlerFunc

	// Routing context pool
	pool *sync.Pool

	// Custom route not found handler
	// TODO: Then dont we need like a map and then make several, also maybe an
	// offline but online style thing

	// Custom method not allowed handler
}

// NewMux returns a newly initialized Mux object that implements the Router
// interface.
func NewMux() *Mux {
	mux := &Mux{tree: &node{}, pool: &sync.Pool{}}
	// TODO: is this the best way>?
	mux.pool.New = func() interface{} {
		return NewRouteContext()
	}
	return mux
}

// ServeHTTP is the single method of the http.Handler interface that makes
// Mux interoperable with the standard library. It uses a sync.Pool to get and
// reuse routing contexts for each request.
func (mx *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Ensure the mux has some routes defined on the mux
	// TODO: Look how we end the route by using a handler and serving the http
	if mx.handler == nil {
		mx.NotFoundHandler().ServeHTTP(w, r)
		return
	}

	// Check if a routing context already exists from a parent router.
	rctx, _ := r.Context().Value(RouteCtxKey).(*Context)
	if rctx != nil {
		mx.handler.ServeHTTP(w, r)
		return
	}

	// Fetch a RouteContext object from the sync pool, and call the computed
	// mx.handler that is comprised of mx.middlewares + mx.routeHTTP.
	// Once the request is finished, reset the routing context and put it back
	// into the pool for reuse from another request.
	rctx = mx.pool.Get().(*Context)
	rctx.Reset()
	rctx.Routes = mx
	r = r.WithContext(context.WithValue(r.Context(), RouteCtxKey, rctx))
	mx.handler.ServeHTTP(w, r)
	mx.pool.Put(rctx)
}

// Use appends a middleware handler to the Mux middleware stack.
//
// The middleware stack for any Mux will execute before searching for a matching
// route to a specific handler, which provides opportunity to respond early,
// change the course of the request execution, or set request-scoped values for
// the next http.Handler.
func (mx *Mux) Use(middlewares ...func(http.Handler) http.Handler) {
	if mx.handler != nil {
		panic("all middlewares must be defined before routes on a mux")
	}
	mx.middlewares = append(mx.middlewares, middlewares...)
}

// Handle adds the route `pattern` that matches any http method to
// execute the `handler` http.Handler.
// TODO: Maybe include a controller name, or even action name, in rails its like
// [controller_name, action_name]
// controller_name+_+action_name+_path
// TODO: Also think how we get databases or meta-app style stuff into our action
//       AND session stuff!

func (mx *Mux) Handle(pattern string, handler http.Handler) {
	mx.handle(mALL, pattern, handler)
}

// HandleFunc adds the route `pattern` that matches any http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) HandleFunc(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mALL, pattern, handlerFn)
}

// TODO: Why this exist? We already have r.Get(), r.Post(), ...
// Method adds the route `pattern` that matches `method` http method to
// execute the `handler` http.Handler.
func (mx *Mux) Method(method, pattern string, handler http.Handler) {
	m, ok := methodMap[strings.ToUpper(method)]
	if !ok {
		panic(fmt.Sprintf("[fatal error] '%s' http method is not supported.", method))
	}
	mx.handle(m, pattern, handler)
}

// TODO: Just seems this is only needed for custom OPTION
// MethodFunc adds the route `pattern` that matches `method` http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) MethodFunc(method, pattern string, handlerFn http.HandlerFunc) {
	mx.Method(method, pattern, handlerFn)
}

// Connect adds the route `pattern` that matches a CONNECT http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Connect(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mCONNECT, pattern, handlerFn)
}

// Delete adds the route `pattern` that matches a DELETE http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Delete(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mDELETE, pattern, handlerFn)
}

// Get adds the route `pattern` that matches a GET http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Get(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mGET, pattern, handlerFn)
}

// Head adds the route `pattern` that matches a HEAD http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Head(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mHEAD, pattern, handlerFn)
}

// Options adds the route `pattern` that matches a OPTIONS http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Options(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mOPTIONS, pattern, handlerFn)
}

// Patch adds the route `pattern` that matches a PATCH http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Patch(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mPATCH, pattern, handlerFn)
}

// Post adds the route `pattern` that matches a POST http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Post(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mPOST, pattern, handlerFn)
}

// Put adds the route `pattern` that matches a PUT http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Put(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mPUT, pattern, handlerFn)
}

// Trace adds the route `pattern` that matches a TRACE http method to
// execute the `handlerFn` http.HandlerFunc.
func (mx *Mux) Trace(pattern string, handlerFn http.HandlerFunc) {
	mx.handle(mTRACE, pattern, handlerFn)
}

// TODO: WE MUST have way to define the 404 view. And this is the model for
// custom endings, and like maybe building in some advanced cluster managment
// NotFound sets a custom http.HandlerFunc for routing paths that could
// not be found. The default 404 handler is `http.NotFound`.
func (mx *Mux) NotFound(handlerFn http.HandlerFunc) {
	// Build NotFound handler chain
	m := mx
	hFn := handlerFn
	if mx.inline && mx.parent != nil {
		m = mx.parent
		hFn = Chain(mx.middlewares...).HandlerFunc(hFn).ServeHTTP
	}

	// Update the notFoundHandler from this point forward
	m.notFoundHandler = hFn
	m.updateSubRoutes(func(subMux *Mux) {
		if subMux.notFoundHandler == nil {
			subMux.NotFound(hFn)
		}
	})
}

// MethodNotAllowed sets a custom http.HandlerFunc for routing paths where the
// method is unresolved. The default handler returns a 405 with an empty body.
func (mx *Mux) MethodNotAllowed(handlerFn http.HandlerFunc) {
	// Build MethodNotAllowed handler chain
	m := mx
	hFn := handlerFn
	if mx.inline && mx.parent != nil {
		m = mx.parent
		hFn = Chain(mx.middlewares...).HandlerFunc(hFn).ServeHTTP
	}

	// Update the methodNotAllowedHandler from this point forward
	m.methodNotAllowedHandler = hFn
	m.updateSubRoutes(func(subMux *Mux) {
		if subMux.methodNotAllowedHandler == nil {
			subMux.MethodNotAllowed(hFn)
		}
	})
}

// With adds inline middlewares for an endpoint handler.
func (mx *Mux) With(middlewares ...func(http.Handler) http.Handler) Router {
	// Similarly as in handle(), we must build the mux handler once further
	// middleware registration isn't allowed for this stack, like now.
	if !mx.inline && mx.handler == nil {
		mx.buildRouteHandler()
	}

	// Copy middlewares from parent inline muxs
	var mws Middlewares
	if mx.inline {
		mws = make(Middlewares, len(mx.middlewares))
		copy(mws, mx.middlewares)
	}
	mws = append(mws, middlewares...)

	im := &Mux{pool: mx.pool, inline: true, parent: mx, tree: mx.tree, middlewares: mws}

	return im
}

// Group creates a new inline-Mux with a fresh middleware stack. It's useful
// for a group of handlers along the same routing path that use an additional
// set of middlewares. See _examples/.
func (m *Mux) Group(fn func(r Router)) Router {
	im := m.With().(*Mux)
	if fn != nil {
		fn(im)
	}
	return im
}

// Route creates a new Mux with a fresh middleware stack and mounts it
// along the `pattern` as a subrouter. Effectively, this is a short-hand
// call to Mount. See _examples/.
func (m *Mux) Route(pattern string, fn func(r Router)) Router {
	subRouter := New()
	if fn != nil {
		fn(subRouter)
	}
	m.Mount(pattern, subRouter)
	return subRouter
}

// Mount attaches another http.Handler or Router as a subrouter along a routing
// path. It's very useful to split up a large API as many independent routers and
// compose them as a single service using Mount. See _examples/.
//
// Note that Mount() simply sets a wildcard along the `pattern` that will continue
// routing at the `handler`, which in most cases is another router.Router. As a result,
// if you define two Mount() routes on the exact same pattern the mount will panic.
func (m *Mux) Mount(pattern string, handler http.Handler) {
	// Provide runtime safety for ensuring a pattern isn't mounted on an existing
	// routing pattern.
	if m.tree.findPattern(pattern+"*") || m.tree.findPattern(pattern+"/*") {
		panic(fmt.Sprintf("[fatal error] attempting to Mount() a handler on an existing path, '%s'", pattern))
	}

	// Assign sub-Router's with the parent not found & method not allowed handler if not specified.
	subr, ok := handler.(*Mux)
	if ok && subr.notFoundHandler == nil && m.notFoundHandler != nil {
		subr.NotFound(m.notFoundHandler)
	}
	if ok && subr.methodNotAllowedHandler == nil && m.methodNotAllowedHandler != nil {
		subr.MethodNotAllowed(m.methodNotAllowedHandler)
	}

	// Wrap the sub-router in a handlerFunc to scope the request path for routing.
	mountHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rctx := RouteContext(r.Context())
		rctx.RoutePath = m.nextRoutePath(rctx)
		handler.ServeHTTP(w, r)
	})

	// TODO empty or last is /
	if len(pattern) == 0 || pattern[len(pattern)-1] != '/' {
		m.handle(mALL|mSTUB, pattern, mountHandler)
		m.handle(mALL|mSTUB, pattern+"/", mountHandler)
		pattern += "/"
	}

	method := mALL
	subroutes, _ := handler.(Routes)
	if subroutes != nil {
		method |= mSTUB
	}
	n := m.handle(method, pattern+"*", mountHandler)

	if subroutes != nil {
		n.subroutes = subroutes
	}
}

// TODO: THIS IS WHAT I WANT FOR THE OUTPUT thats like rake routes in our cmd
//
//	AND I want to create a slice of Route objects and eys they should have a
//	STring() for outputting each in a nice output but also so we can see where
//	they point, CHANGE it, add them, remove them, // it would also be nice to
//	have all the actions (controller endpoints) in a slice, and all the models,
//	maybe views but not sure how that woudl work
//
// Routes returns a slice of routing information from the tree,
// useful for traversing available routes of a router.
func (m *Mux) Routes() []Route {
	return m.tree.routes()
}

// Middlewares returns a slice of middleware handler functions.
func (m *Mux) Middlewares() Middlewares {
	return m.middlewares
}

// Match searches the routing tree for a handler that matches the method/path.
// It's similar to routing a http request, but without executing the handler
// thereafter.
//
// Note: the *Context state is updated during execution, so manage
// the state carefully or make a NewRouteContext().
func (m *Mux) Match(rctx *Context, method, path string) bool {
	mm, ok := methodMap[method]
	if !ok {
		return false
	}

	node, _, h := m.tree.FindRoute(rctx, mm, path)

	if node != nil && node.subroutes != nil {
		rctx.RoutePath = m.nextRoutePath(rctx)
		return node.subroutes.Match(rctx, method, rctx.RoutePath)
	}

	return h != nil
}

// NotFoundHandler returns the default Mux 404 responder whenever a route
// cannot be found.
func (m *Mux) NotFoundHandler() http.HandlerFunc {
	if m.notFoundHandler != nil {
		return m.notFoundHandler
	}
	return http.NotFound
}

// MethodNotAllowedHandler returns the default Mux 405 responder whenever
// a method cannot be resolved for a route.
func (m *Mux) MethodNotAllowedHandler() http.HandlerFunc {
	if m.methodNotAllowedHandler != nil {
		return m.methodNotAllowedHandler
	}
	return methodNotAllowedHandler
}

// buildRouteHandler builds the single mux handler that is a chain of the middleware
// stack, as defined by calls to Use(), and the tree router (Mux) itself. After this
// point, no other middlewares can be registered on this Mux's stack. But you can still
// compose additional middlewares via Group()'s or using a chained middleware handler.
func (m *Mux) buildRouteHandler() {
	m.handler = chain(m.middlewares, http.HandlerFunc(m.routeHTTP))
}

// handle registers a http.Handler in the routing tree for a particular http method
// and routing pattern.
func (m *Mux) handle(method methodType, pattern string, handler http.Handler) *node {
	if len(pattern) == 0 || pattern[0] != '/' {
		panic(fmt.Sprintf("[fatal error] routing pattern must begin with '/' in '%s'", pattern))
	}

	// Build the final routing handler for this Mux.
	if !m.inline && m.handler == nil {
		m.buildRouteHandler()
	}

	// Build endpoint handler with inline middlewares for the route
	var h http.Handler
	if m.inline {
		m.handler = http.HandlerFunc(m.routeHTTP)
		h = Chain(m.middlewares...).Handler(handler)
	} else {
		h = handler
	}

	// Add the endpoint to the tree and return the node
	return m.tree.InsertRoute(method, pattern, h)
}

// routeHTTP routes a http.Request through the Mux routing tree to serve
// the matching handler for a particular http method.
func (m *Mux) routeHTTP(w http.ResponseWriter, r *http.Request) {
	// Grab the route context object
	rctx := r.Context().Value(RouteCtxKey).(*Context)

	// The request routing path
	routePath := rctx.RoutePath
	if len(routePath) == 0 {
		if len(r.URL.RawPath) != 0 {
			routePath = r.URL.RawPath
		} else {
			routePath = r.URL.Path
		}
	}

	// Check if method is supported by the router
	if len(rctx.RouteMethod) == 0 {
		rctx.RouteMethod = r.Method
	}
	method, ok := methodMap[rctx.RouteMethod]
	if !ok {
		m.MethodNotAllowedHandler().ServeHTTP(w, r)
		return
	}

	// Find the route
	if _, _, h := m.tree.FindRoute(rctx, method, routePath); h != nil {
		h.ServeHTTP(w, r)
		return
	}
	if rctx.methodNotAllowed {
		m.MethodNotAllowedHandler().ServeHTTP(w, r)
	} else {
		m.NotFoundHandler().ServeHTTP(w, r)
	}
}

// TODO: SHOULD be using this since URLs are now UTF8 and we are in golang
// utf8.DecodeLastRuneInString(s

func (m *Mux) nextRoutePath(rctx *Context) string {
	routePath := "/"
	nx := len(rctx.routeParams.Keys) - 1 // index of last param in list
	if nx >= 0 && rctx.routeParams.Keys[nx] == "*" && len(rctx.routeParams.Values) > nx {
		routePath += rctx.routeParams.Values[nx]
	}
	return routePath
}

// Recursively update data on child routers.
func (m *Mux) updateSubRoutes(fn func(subMux *Mux)) {
	for _, r := range m.tree.routes() {
		subMux, ok := r.SubRoutes.(*Mux)
		if !ok {
			continue
		}
		fn(subMux)
	}
}

// methodNotAllowedHandler is a helper function to respond with a 405,
// method not allowed.
func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Write(nil)
}
