package controller

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/logrus-error/logerr"
)

type ParamRun struct {
	Fix       bool
	EmptyLine bool
	Args      []string
}

func (c *Controller) Run(_ context.Context, logE *logrus.Entry, param *ParamRun) error {
	failed := false
	for _, arg := range param.Args {
		logE := logE.WithField("file_path", arg)
		if err := c.handleFile(logE, param, arg); err != nil {
			failed = true
			logerr.WithError(logE, err).Error("a file is invalid")
		}
	}
	if failed {
		return errors.New("some files are invalid")
	}
	return nil
}

func (c *Controller) handleFile(logE *logrus.Entry, param *ParamRun, filePath string) error {
	f, err := afero.ReadFile(c.fs, filePath)
	if err != nil {
		return fmt.Errorf("open a file: %w", err)
	}
	content, err := c.handleFileContent(logE, param, string(f))
	if err != nil {
		return err
	}
	if content == "" {
		return nil
	}
	stat, err := c.fs.Stat(filePath)
	if err != nil {
		return fmt.Errorf("get a file stat: %w", err)
	}
	if err := afero.WriteFile(c.fs, filePath, []byte(content), stat.Mode()); err != nil {
		return fmt.Errorf("edit a file: %w", err)
	}
	fmt.Fprintln(c.stdout, filePath)
	return nil
}

func (c *Controller) handleFileContent(logE *logrus.Entry, param *ParamRun, content string) (string, error) {
	lines := strings.Split(content, "\n")
	numOfLines := len(lines)
	if lines[numOfLines-1] != "" {
		if !param.Fix {
			return "", errors.New("a newline at the end of file is missing")
		}
		logE.Warn("a newline at the end of file is missing")
		return content + "\n", nil
	}
	if !param.EmptyLine {
		return "", nil
	}
	if lines[numOfLines-2] != "" {
		return "", nil
	}
	if !param.Fix {
		return "", errors.New("empty lines at the end of file should be trimmed")
	}
	logE.Warn("empty lines at the end of file should be trimmed")
	return strings.TrimSuffix(content, "\n") + "\n", nil
}
