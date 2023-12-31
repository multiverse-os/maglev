package maglev

import (
	database "github.com/multiverse-os/maglev/database"
)

//type Database struct {
//	*database.Database
//
//	Framework
//
//	Name string
//	Path string
//
//	Type database.StoreType
//}

func (f *Framework) KV(storeType database.StoreType) *database.Database {
	if f.databases[storeType.String()] == nil {
		f.databases[storeType.String()] = database.OpenBitcask(storeType.String())
	}
	return f.databases[storeType.String()]
}

func (f *Framework) DB() *database.Database {
	if f.databases[database.Model.String()] == nil {
		f.databases[database.Model.String()] = database.OpenBitcask(database.Model.String())
	}
	return f.databases[database.Model.String()]
}
