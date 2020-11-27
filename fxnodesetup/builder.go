package fxnodesetup

import (
	"github.com/filecoin-project/lotus/node/repo"
	"go.uber.org/fx"
)

// From node/builder.go

// special is a type used to give keys to modules which
//  can't really be identified by the returned type
type Special struct{ Id int }

type Invoke int

type Settings struct {
	// modules is a map of constructors for DI
	//
	// In most cases the index will be a reflect. Type of element returned by
	// the constructor, but for some 'constructors' it's hard to specify what's
	// the return type should be (or the constructor returns fx group)
	Modules map[interface{}]fx.Option

	// invokes are separate from modules as they can't be referenced by return
	// type, and must be applied in correct order
	Invokes []fx.Option

	NodeType repo.RepoType

	Online bool // Online option applied
	Config bool // Config option applied

}
