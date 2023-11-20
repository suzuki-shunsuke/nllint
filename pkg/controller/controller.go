package controller

import (
	"io"

	"github.com/spf13/afero"
)

type Controller struct {
	fs     afero.Fs
	stdout io.Writer
}

func New(fs afero.Fs, stdout io.Writer) *Controller {
	return &Controller{
		fs:     fs,
		stdout: stdout,
	}
}
