package controller_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/nllint/pkg/controller"
)

func TestController_Run(t *testing.T) { //nolint:funlen,gocognit,cyclop
	t.Parallel()
	data := []struct {
		name     string
		files    map[string]string
		expFiles map[string]string
		param    *controller.ParamRun
		stdout   string
		isErr    bool
	}{
		{
			name:  "no file",
			param: &controller.ParamRun{},
		},
		{
			name: "no change",
			files: map[string]string{
				"foo": "hello\n",
			},
			expFiles: map[string]string{
				"foo": "hello\n",
			},
			param: &controller.ParamRun{
				Args: []string{"foo"},
			},
		},
		{
			name: "newline is required",
			files: map[string]string{
				"foo": "hello",
			},
			expFiles: map[string]string{
				"foo": "hello",
			},
			param: &controller.ParamRun{
				Args: []string{"foo"},
			},
			isErr: true,
		},
		{
			name: "TrimSpace is disabled by default",
			files: map[string]string{
				"foo": "  hello\n\n",
			},
			expFiles: map[string]string{
				"foo": "  hello\n\n",
			},
			param: &controller.ParamRun{
				Args: []string{"foo"},
			},
		},
		{
			name: "enable TrimSpace",
			files: map[string]string{
				"foo": "  hello\n\n",
			},
			expFiles: map[string]string{
				"foo": "  hello\n\n",
			},
			param: &controller.ParamRun{
				Args:        []string{"foo"},
				IsTrimSpace: true,
			},
			isErr: true,
		},
		{
			name: "no change (fix)",
			files: map[string]string{
				"foo": "hello\n",
			},
			expFiles: map[string]string{
				"foo": "hello\n",
			},
			param: &controller.ParamRun{
				Args:        []string{"foo"},
				IsTrimSpace: true,
				Fix:         true,
			},
		},
		{
			name: "newline is required (fix)",
			files: map[string]string{
				"foo": "hello",
			},
			expFiles: map[string]string{
				"foo": "hello\n",
			},
			param: &controller.ParamRun{
				Args: []string{"foo"},
				Fix:  true,
			},
		},
		{
			name: "enable TrimSpace (fix)",
			files: map[string]string{
				"foo": "  hello\n\n",
			},
			expFiles: map[string]string{
				"foo": "hello\n",
			},
			param: &controller.ParamRun{
				Args:        []string{"foo"},
				IsTrimSpace: true,
				Fix:         true,
			},
		},
	}
	for _, d := range data {
		d := d
		t.Run(d.name, func(t *testing.T) {
			t.Parallel()
			fs := afero.NewMemMapFs()
			for k, v := range d.files {
				if err := afero.WriteFile(fs, k, []byte(v), 0o644); err != nil {
					t.Fatal(err)
				}
			}
			buf := &bytes.Buffer{}
			ctrl := controller.New(fs, buf)
			if err := ctrl.Run(context.Background(), logrus.NewEntry(logrus.New()), d.param); err != nil {
				if !d.isErr {
					t.Fatal(err)
				}
				stdout := buf.String()
				if diff := cmp.Diff(stdout, d.stdout); diff != "" {
					t.Fatalf("stdout:\n%s", diff)
				}
				for k := range d.files {
					b, err := afero.ReadFile(fs, k)
					if err != nil {
						t.Fatal(err)
					}
					a := d.expFiles[k]
					if diff := cmp.Diff(string(b), a); diff != "" {
						t.Fatalf("%s\n%s", k, diff)
					}
				}
				return
			}
			if d.isErr {
				t.Fatal("error must be returned")
			}
		})
	}
}
