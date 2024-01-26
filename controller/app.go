package controller

import (
	"net/http"
	"time"

	view "github.com/multiverse-os/maglev/view"
)

// Example:
// In Gin

// RouteInfo
//  Method       string
//  Path         string
//  Handler      string
//  HandlerFunc  HanderFunc

// Engine
//   Routergroup (parent mux)
//   various options like redirect trailing slash

// Handler ->

// TODO: Okay for NOW we stick with this design and the current design of the
// view so we can sort out the model portion of the mVC; then we can zero on on
// a router/http server solution which will finalize everything

// NOTE: VIEWS should contain folders like application or sessions, and they
// should contain all the helpers needed to build the session views in the
// session file in views. This way in views its ONLY the completed pages to
// serve up.

func (c Controller) Root(body http.ResponseWriter, request *http.Request) {
	defer c.Framework.Benchmark(time.Now(), "Root()")
	// TODO: Where in the world would this be getting set??? realistically
	// where?>??
	// We were successfully just using the function directly; and we could have
	// jsut as easily passed any necessary params through the view function
	// parameter; arguably this will allow for richer access, and more fine
	// grained control over perahps exactly what wshould be all;owed to be used vs
	// free form and letting the developer make every choice which seems beter
	// I do like the idea of encapsulating the view concept tho-ugh
	body.Write(view.Root().Bytes()) // Because it could be .JSON ?
}
