package controller

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/logrus-error/logerr"
)

type ParamRun struct {
	Fix             bool
	IsTrimSpace     bool
	IsTrailingSpace bool
	IgnoreNotFound  bool
	Args            []string
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
	content, err := handleFileContent(logE, param, string(f))
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

func handleFileContent(logE *logrus.Entry, param *ParamRun, content string) (string, error) {
	lines := strings.Split(content, "\n")
	if param.IsTrailingSpace {
		for i, line := range lines {
			newL := strings.TrimRightFunc(line, unicode.IsSpace)
			if newL != line {
				logE.WithField("line_number", i+1).Warn("trailing white spaces in a line are found")
			}
			lines[i] = newL
		}
	}
	if lines[len(lines)-1] != "" {
		// Check if the last line has a newline character
		logE.Warn("a newline at the end of file is missing")
		lines = append(lines, "")
	}
	var newContent string
	if param.IsTrimSpace {
		newContent = strings.TrimSpace(strings.Join(lines, "\n")) + "\n"
	} else {
		newContent = strings.Join(lines, "\n")
	}
	if content == newContent {
		return "", nil
	}
	if !param.Fix {
		return "", errors.New("white spaces in files should be trimmed")
	}
	logE.Info("white spaces in files are trimmed")
	return newContent, nil
}
