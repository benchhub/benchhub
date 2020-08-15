// Code generated by tqbuilder from github.com/benchhub/benchhub/core/services/git/schema/ddl DO NOT EDIT.

package gitschema

import (
	"github.com/benchhub/benchhub/lib/tqbuilder/sql/ddl"
)

// ----------------------------------------------------------------------------
// GitOwner

var GitOwner = newGitOwner()

// GitOwnerTable is generated from table git_owners
type GitOwnerTable struct {
	Id     ddl.ColumnDef // id
	Type   ddl.ColumnDef // type
	HostId ddl.ColumnDef // host_id
	Name   ddl.ColumnDef // name
}

func newGitOwner() GitOwnerTable {
	// TODO: fill in column def etc.
}

// ----------------------------------------------------------------------------
// GitRepo

var GitRepo = newGitRepo()

// GitRepoTable is generated from table git_repos
type GitRepoTable struct {
	Id       ddl.ColumnDef // id
	Type     ddl.ColumnDef // type
	OwnerId  ddl.ColumnDef // owner_id
	Owner    ddl.ColumnDef // owner
	Name     ddl.ColumnDef // name
	Goimport ddl.ColumnDef // goimport
}

func newGitRepo() GitRepoTable {
	// TODO: fill in column def etc.
}
