package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/benchhub/benchhub/exp/qaq16/config"
	"github.com/dyweb/gommon/util/fsutil"
)

// $data/log
func NewLogDir(cfg config.Data, tm time.Time) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	t := fmt.Sprintf("%d-%02d-%02d/%02d-%02d", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute())
	p := filepath.Join(wd, cfg.Dir, "log", t)
	if err := os.MkdirAll(p, fsutil.DefaultDirPerm); err != nil {
		return p, err
	}
	return p, nil
}

func NewMountDir(prefix string, name string) (string, error) {
	p := filepath.Join(prefix, name)
	if err := os.Mkdir(p, fsutil.DefaultDirPerm); err != nil {
		return "", err
	}
	return p, nil
}

func FormatLog(prefix string, name string) string {
	return filepath.Join(prefix, name+".log")
}
