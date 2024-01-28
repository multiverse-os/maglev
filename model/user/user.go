package user

import (
	model "github.com/multiverse-os/webframe/model"
)

// AH BUT THIS IS GLOBAL; omg omg lets make omg omg a fucntion?
// ARG now we need access to app object; this is getting worse.

// So if we do like database.OpenKV() we would likely create a new instance
// interacting wtih the files; we want to share existing handlers. Or maybe we
// dont initialize them in app!
// TODO: This would need our New() function for our concept of using just
// app/models(should be model); then we can do
//    u = model.User.New("test")
// Then we have
//    u.Save() => bool
// built in ready to go. Now yes this isnt exactly active record
// but feels like a really good starting point and skeleton to build on and
// giving us the basic MVC we are looking for.

// TODO: So this cant be like this, it has to be a struct, then in models we
// initialize it presumably; and so Collection should be local and
// model.Collection needs to be interface
// -- wRONG-- we still have collection to make our initialziation sensible; and
// maybe just use instance
type Model struct {
	model.Collection
}

// TODO: Ok something LIKE this but not this. And we would do some step where we
// iterate over each of the models and put them into the database object as
// collections and that may or may not end up in the final app (now thinking the
// app object should be slim slim slim
func Init() (newModel Model) {
	newModel.Collection = model.NewCollection("user")
	return newModel
}

// ^ if we can get above into model mostly that would be beaut

// TODO: NICE design is going well when we desperately try to make it work and
// the logic flows and makes sense naturally; creating new from our collection
// gives us a record
func (self Model) New(username string, password string) *Record {
	userRecord := &Record{
		Username: username,
		Password: password,
	}

	// TODO: *** So we need some way to well do below, laod the record to
	// instance, and we need to create an Id (using muid) assign collection
	// could do this on first save.
	//userRecord.Collection = self.Collection
	userRecord.Collection = self.Collection.UseDB(model.Database)

	userRecord.New() // => This should result either result in save4 and return bool or it should be gereated a valid empty object
	userRecord.Instance = userRecord
	return userRecord
}

// TODO: Maybe here we specify the database the model uses otherwise maybe it
// defaults? Even so we should implement simply first

// TODO: Cant realistically do the login system without models worked out!!!!

// TODO: Though if we dont do the embedded style below we cant get the All(),
// First() etc built in and tahts important but we still need to get basics
// going.

// TODO: How do we get sensible user.Collection and get this inherit from
// model.Collection

// Model
type Record struct {
	model.Record

	Username string
	Password string

	Name string
}

// Save()
// New()
// Delete()

// DEV function
func IsModelInstance(m model.Instance) bool {
	return true
}

// This conecept of Users will let us create selections easier and provide
// First() Last() but there must be a way we can more easily get that
// functionality

// Methods
