package testutil

import (
	"github.com/dyweb/gommon/util/runtimeutil"
	"strings"
)

const Project = "github.com/benchhub/benchhub"

// ProjectRoot returns absolute path of project using runtime library
func ProjectRoot() string {
	frame := runtimeutil.GetCallerFrame(0)
	file := frame.File
	return file[:strings.Index(file, Project)+len(Project)]
}

// PkgRoot returns absolute path of the pkg package
func PkgRoot() string {
	return ProjectRoot() + "/pkg"
}

// CentralRoot returns absolute path of pkg/central package
func CentralRoot() string {
	return PkgRoot() + "/central"
}

// CentralTestdata returns absolute path of pkg/central/testdata/{file}
func CentralTestdata(file string) string {
	return CentralRoot() + "/testdata/" + file
}
