package cli

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/nllint/pkg/controller"
	"github.com/urfave/cli/v3"
)

type Runner struct {
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
	LDFlags    *LDFlags
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

func (r *Runner) Run(ctx context.Context, logger *slog.Logger, args ...string) error {
	app := &cli.Command{
		Name:  "nllint",
		Usage: "Check newlines at the end of files",
		CustomRootCommandHelpTemplate: `nllint - Check newlines at the end of files

https://github.com/suzuki-shunsuke/nllint

Usage:
  nllint [-fix (-f)] [-trim-space (-s)] [-trim-trailing-space (-S)] [-ignore-notfound (-i)] <file path> [<file path>...]

Options:
  -help, -h                 Show help
  -version, -v              Show version
  -fix, -f                  Fix files
  -trim-space, -s           Disallow leading and trailing white spaces in files
  -trim-trailing-space, -S  Disallow trailing white spaces in each line
  -ignore-notfound, -i      Ignore not found files
`,
		Version: r.LDFlags.VersionString(),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "fix",
				Aliases: []string{"f"},
				Usage:   "Fix files",
			},
			&cli.BoolFlag{
				Name:    "trim-space",
				Aliases: []string{"s"},
				Usage:   "Disallow leading and trailing white spaces in files",
			},
			&cli.BoolFlag{
				Name:    "trim-trailing-space",
				Aliases: []string{"S"},
				Usage:   "Disallow trailing white spaces in each line",
			},
			&cli.BoolFlag{
				Name:    "ignore-notfound",
				Aliases: []string{"i"},
				Usage:   "Ignore not found files",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			return r.run(ctx, logger, c)
		},
	}

	return app.Run(ctx, args) //nolint:wrapcheck
}

func (r *Runner) run(ctx context.Context, logger *slog.Logger, c *cli.Command) error {
	param := &controller.ParamRun{
		Fix:             c.Bool("fix"),
		IsTrimSpace:     c.Bool("trim-space"),
		IsTrailingSpace: c.Bool("trim-trailing-space"),
		IgnoreNotFound:  c.Bool("ignore-notfound"),
		Args:            c.Args().Slice(),
	}

	ctrl := controller.New(afero.NewOsFs(), r.Stdout)
	return ctrl.Run(ctx, logger, param) //nolint:wrapcheck
}
