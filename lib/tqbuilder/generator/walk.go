package generator

import (
	"os"
	"path/filepath"
	"strings"
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
