package app

// TODO
// It would be nice if we only needed to interact with framework, and framework
// dealt with the submodules itself to simplify things as much as possible

// TODO: So this should kinda feel like application.rb in the root folder where
// one can basically setup the application somewhere around initialization
// and this is initialization.
import (
	"github.com/multiverse-os/maglev/controller"
	framework "github.com/multiverse-os/webframe"
)

type App struct {
	framework.Framework
}

func Init(cfg framework.Config) App {
	app := App{framework.Init(cfg)}

	// Database Initialization
	// TODO, maybe I should pass the databases I want to intiailize so I can
	// be selective and get rid of the app.KV()

	app.InitializeDBs()
	//app.KV(framework.ModelStore)
	//app.KV(framework.CacheStore)
	//app.KV(framework.SessionStore)

	// Model
	app.NewModel("user")
	app.NewModel("post")
	app.NewModel("task")

	// Controller
	app.NewController("app")
	app.NewController("session")

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

func (app App) Controller(name string) controller.Actions {
	return controller.Actions(app.Controllers[name])
}

func (app *App) Model(name string) Model {
	return app.Models[name]
}
