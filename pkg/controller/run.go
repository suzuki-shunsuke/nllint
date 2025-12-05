package controller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"unicode"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/slog-error/slogerr"
)

type ParamRun struct {
	Fix             bool
	IsTrimSpace     bool
	IsTrailingSpace bool
	IgnoreNotFound  bool
	Args            []string
}

func (c *Controller) Run(_ context.Context, logger *slog.Logger, param *ParamRun) error {
	failed := false
	for _, arg := range param.Args {
		logger := logger.With("file_path", arg)
		if err := c.handleFile(logger, param, arg); err != nil {
			failed = true
			slogerr.WithError(logger, err).Error("a file is invalid")
		}
	}
	if failed {
		return errors.New("some files are invalid")
	}
	return nil
}

func (c *Controller) handleFile(logger *slog.Logger, param *ParamRun, filePath string) error {
	f, err := afero.ReadFile(c.fs, filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) && param.IgnoreNotFound {
			logger.Warn("ignore a file because it doesn't exist")
			return nil
		}
		return fmt.Errorf("open a file: %w", err)
	}
	content, err := handleFileContent(logger, param, string(f))
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

func handleFileContent(logger *slog.Logger, param *ParamRun, content string) (string, error) {
	lines := strings.Split(content, "\n")
	if param.IsTrailingSpace {
		for i, line := range lines {
			newL := strings.TrimRightFunc(line, unicode.IsSpace)
			if newL != line {
				logger.Warn("trailing white spaces in a line are found", "line_number", i+1)
			}
			lines[i] = newL
		}
	}
	if lines[len(lines)-1] != "" {
		// Check if the last line has a newline character
		logger.Warn("a newline at the end of file is missing")
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
		return "", errors.New("a file should be fixed")
	}
	logger.Info("a file is fixed")
	return newContent, nil
}
