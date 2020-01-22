package nodeutil

import (
	"github.com/rs/xid"
)

var uid string

// generate unique id
// https://github.com/benchhub/benchhub/issues/17
// https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html

func UID() string {
	return uid
}

func NewUID() string {
	id := xid.New()
	//log.Info(id.String(), id.Pid(), id.Machine(), id.Time())
	return id.String()
}

func init() {
	uid = NewUID()
}
