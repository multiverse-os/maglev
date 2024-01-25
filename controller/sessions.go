package controller

import (
	"fmt"
	"net/http"
	"time"

	view "github.com/multiverse-os/maglev-app/view"
)

func (c Controller) Login(body http.ResponseWriter, request *http.Request) {
	defer c.Framework.Benchmark(time.Now(), "Login()")
	c.Framework.Log("c.Framework.Config.Name=(%v)\n", c.Framework.Config.Name)
	c.Framework.Log("c.Framework.DB().CacheStore=(%v)", c.Framework.CacheDB().Store)

	body.Write(view.Login().Bytes())
}

func (c Controller) Register(body http.ResponseWriter, request *http.Request) {
	defer c.Framework.Benchmark(time.Now(), "Register()")
	body.Write(view.Register().Bytes())
}

func (c Controller) NewSession(body http.ResponseWriter, request *http.Request) {
	defer c.Framework.Benchmark(time.Now(), "NewSession()")
	// TODO: Do we really need to request.ParseForm()? Can this not be done in a
	// standardized framework middleware that checks if form's are capable of
	// being parsed and automatically handle preparing the data for use in the
	// controller to reduce any unnecessary code?
	request.ParseForm()

	uid := request.Form.Get("uid")
	fmt.Println("uid:", uid)

	pwd := request.Form.Get("pwd")
	fmt.Println("pwd:", pwd)

	fmt.Println("login controller")
	body.Write(view.Login().Bytes())
}
