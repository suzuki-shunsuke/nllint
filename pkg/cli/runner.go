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
	args := &Args{}
	return (&cli.Command{ //nolint:wrapcheck
		Name:    "nllint",
		Usage:   "Check newlines at the end of files",
		Version: ldFlags.VersionString(),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "fix",
				Aliases:     []string{"f"},
				Usage:       "Fix files",
				Destination: &args.Fix,
			},
			&cli.BoolFlag{
				Name:        "trim-space",
				Aliases:     []string{"s"},
				Usage:       "Disallow leading and trailing white spaces in files",
				Destination: &args.TrimSpace,
			},
			&cli.BoolFlag{
				Name:        "trim-trailing-space",
				Aliases:     []string{"S"},
				Usage:       "Disallow trailing white spaces in each line",
				Destination: &args.TrimTrailingSpace,
			},
			&cli.BoolFlag{
				Name:        "ignore-notfound",
				Aliases:     []string{"i"},
				Usage:       "Ignore not found files",
				Destination: &args.IgnoreNotFound,
			},
		},
		Arguments: []cli.Argument{
			&cli.StringArgs{
				Name:        "file",
				Destination: &args.Files,
				Max:         -1,
			},
		},
		Action: func(ctx context.Context, _ *cli.Command) error {
			return r.run(ctx, logger.Logger, args)
		},
	}).Run(ctx, env.Args)
}

type Runner struct {
	Stdout io.Writer
}

type Args struct {
	Fix               bool
	TrimSpace         bool
	TrimTrailingSpace bool
	IgnoreNotFound    bool

	Files []string
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

func (r *Runner) run(ctx context.Context, logger *slog.Logger, args *Args) error {
	param := &controller.ParamRun{
		Fix:             args.Fix,
		IsTrimSpace:     args.TrimSpace,
		IsTrailingSpace: args.TrimTrailingSpace,
		IgnoreNotFound:  args.IgnoreNotFound,
		Args:            args.Files,
	}

	ctrl := controller.New(afero.NewOsFs(), r.Stdout)
	return ctrl.Run(ctx, logger, param) //nolint:wrapcheck
}
