package cli

import (
	"context"
	"io"
	"log/slog"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/nllint/pkg/controller"
	"github.com/suzuki-shunsuke/slog-util/slogutil"
	"github.com/suzuki-shunsuke/urfave-cli-v3-util/urfave"
	"github.com/urfave/cli/v3"
)

func Run(ctx context.Context, logger *slogutil.Logger, env *urfave.Env) error {
	r := &Runner{
		Stdout: env.Stdout,
	}
	ldFlags := &LDFlags{
		Version: env.Version,
	}
	return urfave.Command(env, &cli.Command{ //nolint:wrapcheck
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
		Version: ldFlags.VersionString(),
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
			return r.run(ctx, logger.Logger, c)
		},
	}).Run(ctx, env.Args)
}

type Runner struct {
	Stdout io.Writer
}

type LDFlags struct {
	Version string
}

func (l *LDFlags) VersionString() string {
	if l == nil {
		return "unknown"
	}
	if l.Version == "" {
		return "unknown"
	}
	return l.Version
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
