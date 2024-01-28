package controller

import (
	"fmt"
	"net/http"
	"time"

	view "github.com/multiverse-os/maglev/view"
)

func (a Actions) Login(body http.ResponseWriter, request *http.Request) {
	defer a.Framework.Benchmark(time.Now(), "Login()")
	a.Framework.Log("a.Framework.Config.Name=(%v)\n", a.Framework.Config.Name)

	// TODO: But really we should let cache store precompiled HTML that doesn't
	// change throughout the session or other important data to be grabbed
	// easily

	a.Framework.Log("c.Framework.Cache()=(%v)", a.Framework.Cache())

	body.Write(view.Login().Bytes())
}

func (a Actions) Register(body http.ResponseWriter, request *http.Request) {
	defer a.Framework.Benchmark(time.Now(), "Register()")
	body.Write(view.Register().Bytes())
}

func (a Actions) NewSession(body http.ResponseWriter, request *http.Request) {
	defer a.Framework.Benchmark(time.Now(), "NewSession()")
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
