package generator

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dyweb/gommon/errors"
)

type WalkResult struct {
	DDLs []ExtractedPath
	DMLS []ExtractedPath
}

func Walk(root string, importPrefix string) (WalkResult, error) {
	var (
		ddls []ExtractedPath
		dmls []ExtractedPath
	)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// TODO： duplication for ddl and dml ...
		if strings.HasSuffix(path, "schema/ddl") && info.IsDir() {
			p, err := ExtractPath(path, importPrefix)
			if err != nil {
				return err
			}
			ddls = append(ddls, p)
		}
		if strings.HasSuffix(path, "schema/dml") && info.IsDir() {
			p, err := ExtractPath(path, importPrefix)
			if err != nil {
				return err
			}
			dmls = append(dmls, p)
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
