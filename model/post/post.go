package post

import "github.com/multiverse-os/webkit/model"

// TODO: Record is naturally going to need to be tied to DB and collection but
// knowing myself I already did that through the model.Record call
type Record struct {
	model.Record

	Name        string
	Description string
}

// TODO: I added attributes and methods, but with our record we don't need that
// we could just add our attributes to the record itself and the same with
// methods. Methods in the other way did allow for eventually restricting access
// more easily but we can still figure out a method that merges the benefits of
// each
