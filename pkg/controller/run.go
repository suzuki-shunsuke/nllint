package controller

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/logrus-error/logerr"
)

type ParamRun struct {
	Fix            bool
	IsTrimSpace    bool
	IgnoreNotFound bool
	Args           []string
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
		if errors.Is(err, os.ErrNotExist) && param.IgnoreNotFound {
			logE.Warn("ignore a file because it doesn't exist")
			return nil
		}
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
		if param.IsTrimSpace {
			return strings.TrimSpace(content) + "\n", nil
		}
		return content + "\n", nil
	}
	if !param.IsTrimSpace {
		return "", nil
	}
	newContent := strings.TrimSpace(content) + "\n"
	if content == newContent {
		return "", nil
	}
	if !param.Fix {
		return "", errors.New("leading and trailing white spaces in files should be trimmed")
	}
	logE.Warn("leading and trailing white spaces in files should be trimmed")
	return newContent, nil
}
