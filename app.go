package app

// TODO
// It would be nice if we only needed to interact with framework, and framework
// dealt with the submodules itself to simplify things as much as possible

// TODO: So this should kinda feel like application.rb in the root folder where
// one can basically setup the application somewhere around initialization
// and this is initialization.
import (
	"fmt"

	framework "github.com/multiverse-os/webframe"
)

type App struct {
	framework.Framework
}

func Init(cfg framework.Config) App {
	app := App{framework.Init(cfg)}

	// Database Initialization
	// TODO: This is a alias to setup all three of these basic databases
	app.InitializeDBs()
	//app.KV(framework.ModelStore)
	//app.KV(framework.CacheStore)
	//app.KV(framework.SessionStore)

	app.Framework.Cache().Put([]byte("key"), []byte("value"))

	val, _ := app.Framework.Cache().Get([]byte("key"))

	fmt.Printf(
		"GET[on]CacheDB app.Framework.Cache().Get([]byte('key')): %v\n",
		string(val),
	)

	// Model
	app.NewModel("user")
	app.NewModel("post")
	app.NewModel("task")

	// Controller
	app.NewController("app")
	app.NewController("session")

	// TODO: We really want to be able to establish middleware here, the routes
	// on the otherhand should be in routes, maybe the middleware should be
	// there too
	//router.Use(middleware.RealIP)
	//router.Use(middleware.Logger)
	//router.Use(middleware.Recoverer)
	//router.Use(middleware.DefaultCompress)
	//router.Use(middleware.Timeout(60 * time.Second))

	// Server
	// TODO: This should be automatic and hidden, we want as much as possible
	// hidden that isn't changed by the developer, here we only want things the
	// developer can affect
	app.Routes()
	app.HTTP().UseRouter(app.Router)

	// Hooks
	//app.AppendToShutdown(app.TestShutdownProcess)

	return app
}

func (app App) Controller(name string) Controller {
	return Controller(app.Controllers[name])
}

func (app *App) Model(name string) Model {
	return app.Models[name]
}
