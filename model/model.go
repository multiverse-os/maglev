package model

import (
	framework "github.com/multiverse-os/webkit"
)

// TODO: I want to switch it to webframe to match framekit and cli by being
// more generic. Overtime ive grown to hate maglev
// imagine "github.com/multiverse-os/maglev/framework would even be better!
// maybe netframe or netapp (since its broader than web) nepp appframe

// TODO: SO what do we do here?? We define globals that are tied to the model?
// ITS very nice API

type Model framework.Model

// TODO: I didnt hate this beacuse it gave us model.User.new() or so one
//var User = user.Init()
