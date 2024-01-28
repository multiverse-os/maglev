package app

import (
	"net/http"
)

func (app *App) Routes() {

	// TODO: Right now the differation between the separate controllers is
	// meaningless because it carries no other data than the framework its
	// connected to. If this remains the case then we dont need the local
	// controller type (which may be useful because it greatly simplifies the
	// entire app portion of the project)

	// app.Controller.Root()

	// controller should eventually be the equivilent to the engine in the gin
	// framework; it should contain:
	//    * the router (mux)
	//    * the before/after action hooks
	//    * available actions
	//    * models, and whitelisting of field access
	//    * cached results
	//    * a reduced context (isnt whole app object or framework object)
	//    * ServeHTTP() method, so it can be  used as a handler

	//    Controller will need a way to handle mapping paths to actions
	//               and a way to interact with the models in a sane manner

	// TODO: THis looks right, we route a specific path to a controller action
	// and this will be very much like Rails which we are trying to mimic as
	// much as possible. SO many web frameworks claim to be rails like but this
	// web framework should feel natural to rails users ideally

	app.Router.Get("/", app.Controller("app").Root)
	app.Router.Get("/ducks", app.Controller("app").Root)
	app.Router.Get("/login", app.Controller("session").Login)

	//app.Router.Post("/sessions/new", controller.NewSession)
	//app.Router.Get("/register", sessionControlle.Register)

	app.Router.Get("/about", func(body http.ResponseWriter, r *http.Request) {
		body.Write([]byte("hi"))
		//app.Views = app.Views + 1
	})
}
