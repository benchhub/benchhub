// Code generated by tqbuilder from github.com/benchhub/benchhub/core/services/git/schema/ddl DO NOT EDIT.

package gitmodel

// GitOwner is generated from table git_owners
type GitOwner struct {
	Id     int    // id
	Type   int    // type
	HostId int    // host_id
	Name   string // name
}

// GitRepo is generated from table git_repos
type GitRepo struct {
	Id       int    // id
	Type     int    // type
	OwnerId  int    // owner_id
	Owner    string // owner
	Name     string // name
	Goimport string // goimport
}
