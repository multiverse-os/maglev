package app

// TODO
// It would be nice if we only needed to interact with framework, and framework
// dealt with the submodules itself to simplify things as much as possible
import (
	"fmt"

	controller "github.com/multiverse-os/maglev/controller"
	model "github.com/multiverse-os/maglev/model"

	framework "github.com/multiverse-os/webkit"
)

type App struct {
	framework.Framework
}

func Init(cfg framework.Config) App {
	app := App{framework.Init(cfg)}

	// Database Initialization
	app.KV(framework.ModelStore)
	app.KV(framework.CacheStore)

	app.Framework.CacheDB().Store.Put([]byte("key"), []byte("value"))
	val, _ := app.Framework.CacheDB().Store.Get([]byte("key"))
	fmt.Printf("GET[on]CacheDB app.Framework.CacheDB().Store.Get([]byte('key')):", string(val))

	// Model
	app.NewModel("user")

	// Controller
	app.NewController("app")
	app.NewController("session")

	// NEW CONTROLLER DESIGN
	//   The new design is implementing a controller type that is separate from
	//   routes and the server.

	//   This controller object will implement serverMux, Handler() by having a
	//   ServeHTTP() method.

	//   This controller will hold model/database access and be able to create a
	//   limited access context that is passed to views in a standardized way (not
	//   random parameters on views

	//   (database is separate, only interaction with db is via model interaction)

	//   router=controller, and you build the complete router by creating a tree
	//   of encapsualted routes in routers that are linked together by the way
	//   they path.

	// **IMPORTANT** We dont ever assign to the "fallback" handler (the global
	// one), we create our own router and set it and the handler when setting up
	// the http server. (IDEALLY we create it in a way we could easily have a
	// ServerDNS() handler to setup a dns server in a near identical way.

	//router.Use(middleware.RealIP)
	//router.Use(middleware.Logger)
	//router.Use(middleware.Recoverer)
	//router.Use(middleware.DefaultCompress)
	//router.Use(middleware.Timeout(60 * time.Second))

	//// Set up our root handlers
	//router.Get("/", HelloWorld)

	//// Set up our API
	//router.Mount("/api/v1/", v1.NewRouter())

	// Server
	app.Routes()
	app.HTTP().UseRouter(app.Router)

	// Hooks
	//app.AppendToShutdown(app.TestShutdownProcess)

	return app
}

func (app App) Controller(name string) controller.Controller {
	return controller.Controller(app.Controllers[name])
}

func (app App) Model(name string) model.Model {
	return model.Model(app.Models[name])
}
