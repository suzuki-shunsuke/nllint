package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/nllint/pkg/controller"
	"github.com/urfave/cli/v2"
)

type Runner struct {
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
	LDFlags    *LDFlags
	LogE       *logrus.Entry
	Env        *Env
	IsTerminal bool
}

type Env struct {
	Config     string
	ConfigBody string
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func (l *LDFlags) VersionString() string {
	if l == nil {
		return "unknown"
	}
	if l.Version == "" {
		if l.Date == "" {
			return "unknown"
		}
		return fmt.Sprintf("(%s)", l.Date)
	}
	if l.Date == "" {
		return l.Version
	}
	return fmt.Sprintf("%s (%s)", l.Version, l.Date)
}

func (r *Runner) Run(ctx context.Context, args ...string) error {
	app := &cli.App{
		Name:  "nllint",
		Usage: "Check newlines at the end of files",
		CustomAppHelpTemplate: `nllint - Check newlines at the end of files

https://github.com/suzuki-shunsuke/nllint

Usage:
  nllint [-fix] [-empty-line] <file path> [<file path>...]

Options:
  -help, -h        Show help
  -version, -v     Show version
	-fix, -f         Fix files
	-empty-line, -e  Disallow leading and trailing white spaces in files
`,
		Version: r.LDFlags.VersionString(),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "fix",
				Aliases: []string{"f"},
				Usage:   "Fix files",
			},
			&cli.BoolFlag{
				Name:    "empty-line",
				Aliases: []string{"e"},
				Usage:   "Disallow empty lines at the end of files",
			},
		},
		Action: r.run,
	}

	return app.RunContext(ctx, args) //nolint:wrapcheck
}

func (r *Runner) run(c *cli.Context) error {
	param := &controller.ParamRun{
		Fix:       c.Bool("fix"),
		EmptyLine: c.Bool("empty-line"),
		Args:      c.Args().Slice(),
	}

	ctrl := controller.New(afero.NewOsFs(), r.Stdout)
	return ctrl.Run(c.Context, r.LogE, param) //nolint:wrapcheck
}
