package service

import (
	_ "embed"

	"github.com/spf13/afero"
	"go.octolab.org/unsafe"
)

var FS afero.Fs

//go:embed .env
var env []byte

func init() {
	FS = afero.NewMemMapFs()
	file := unsafe.Return(FS.Create(".env")).(afero.File)
	unsafe.DoSilent(file.Write(env))
}
