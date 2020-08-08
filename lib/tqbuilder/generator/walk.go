package generator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dyweb/gommon/errors"
)

type WalkResult struct {
	DDLs []string
	DMLS []string
}

func Walk(root string) (WalkResult, error) {
	var (
		ddls []string
		dmls []string
	)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "schema/ddl") && info.IsDir() {
			ddls = append(ddls, path)
		}
		if strings.HasSuffix(path, "schema/dml") && info.IsDir() {
			dmls = append(dmls, path)
		}
		return nil
	})
	if err != nil {
		return WalkResult{}, err
	}
	return WalkResult{
		DDLs: ddls,
		DMLS: dmls,
	}, nil
}

type ExtractedPath struct {
	Path       string // relative path of the file
	Package    string // go package name
	ImportPath string // go import path
	OutputPath string // folder for generated file
}

func ExtractPath(p string, importPrefix string) (ExtractedPath, error) {
	// p is core/services/user/schema/ddl
	segs := strings.Split(p, "/")
	if len(segs) < 3 {
		return ExtractedPath{}, errors.Errorf("expect at least 3 segments in ddl path, got %d from %s", len(segs), p)
	}
	segs[len(segs)-1] = "generated"
	return ExtractedPath{
		Path:       p,                      // core/services/user/schema/ddl
		Package:    segs[len(segs)-3],      // user
		ImportPath: importPrefix + "/" + p, // github.com/benchhub/benchhub/core/services/user/schema/ddl
		OutputPath: filepath.Join(segs...), // core/services/user/schema/generated
	}, nil
}
