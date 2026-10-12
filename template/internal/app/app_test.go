package app_test

import (
	"bytes"
	"context"
	"slices"
	"strings"
	"testing"

	"github.com/standards-lab/go-cli-sdk"
	"github.com/standards-lab/go-core/lifecycle"
	"github.com/standards-lab/go-core/process"

	"github.com/standards-lab/go-cli-sdk-template/template/internal/app"
)

// shutdownTimeout is the one variable the template reads.
var shutdownTimeout = lifecycle.NewEnv("APP").ShutdownTimeout

// newApp returns the App over fresh buffers, and the log of the nodes its
// runs build, recorded with Observe on its published graph.
func newApp(t *testing.T) (a *app.App, stdout, stderr *bytes.Buffer, built *[]string) {
	t.Helper()
	stdout, stderr, built = new(bytes.Buffer), new(bytes.Buffer), new([]string)
	a = app.New(cli.Streams{Stdin: strings.NewReader(""), Stdout: stdout, Stderr: stderr})
	a.Graph().Observe(func(name string) { *built = append(*built, name) })
	return a, stdout, stderr, built
}

// probe adds a leaf at a's root that declares the lifecycle config node and
// reports whether it ran.
func probe(a *app.App) *bool {
	ran := new(bool)
	a.Root().Add((&cli.Command{
		Name:    "probe",
		Summary: "Build the lifecycle config",
		Run: func(context.Context, *cli.Invocation) error {
			*ran = true
			return nil
		},
	}).Use(a.Nodes().LifecycleConfig))
	return ran
}

func TestRun_RootHelp(t *testing.T) {
	for _, args := range [][]string{nil, {"-h"}, {"--help"}} {
		t.Run(strings.Join(args, " "), func(t *testing.T) {
			// An unreadable configuration proves the run reads none.
			t.Setenv(shutdownTimeout, "not-a-duration")
			a, stdout, stderr, built := newApp(t)

			code := a.Run(context.Background(), args)

			if code != process.ExitUsage {
				t.Errorf("code = %d, want %d; stderr = %q", code, process.ExitUsage, stderr)
			}
			if !strings.Contains(stdout.String(), "version") {
				t.Errorf("stdout = %q, want the root help listing version", stdout)
			}
			if stderr.Len() != 0 {
				t.Errorf("stderr = %q, want empty", stderr)
			}
			if len(*built) != 0 {
				t.Errorf("nodes built = %q, want none", *built)
			}
		})
	}
}

func TestRun_UnknownCommand(t *testing.T) {
	t.Setenv(shutdownTimeout, "not-a-duration")
	a, stdout, stderr, built := newApp(t)

	code := a.Run(context.Background(), []string{"bogus"})

	if code != process.ExitUsage {
		t.Errorf("code = %d, want %d", code, process.ExitUsage)
	}
	if want := `app: unknown command "bogus"`; !strings.HasPrefix(stderr.String(), want) {
		t.Errorf("stderr = %q, want it to start with %q", stderr, want)
	}
	if stdout.Len() != 0 {
		t.Errorf("stdout = %q, want empty", stdout)
	}
	if len(*built) != 0 {
		t.Errorf("nodes built = %q, want none", *built)
	}
}

func TestRun_Version(t *testing.T) {
	t.Setenv(shutdownTimeout, "not-a-duration")
	a, stdout, stderr, built := newApp(t)

	code := a.Run(context.Background(), []string{"version"})

	if code != process.ExitOK {
		t.Errorf("code = %d, want %d; stderr = %q", code, process.ExitOK, stderr)
	}
	// A test binary carries no main module version.
	if stdout.String() != "(devel)\n" {
		t.Errorf("stdout = %q, want %q", stdout, "(devel)\n")
	}
	if stderr.Len() != 0 {
		t.Errorf("stderr = %q, want empty", stderr)
	}
	if len(*built) != 0 {
		t.Errorf("nodes built = %q, want none", *built)
	}
}

func TestNew_IsCold(t *testing.T) {
	t.Setenv(shutdownTimeout, "not-a-duration")

	_, stdout, stderr, built := newApp(t)

	if stdout.Len() != 0 || stderr.Len() != 0 {
		t.Errorf("New() wrote %q, %q, want nothing", stdout, stderr)
	}
	if len(*built) != 0 {
		t.Errorf("nodes built = %q, want none", *built)
	}
}

func TestRun_CommandBuildsTheLifecycleConfig(t *testing.T) {
	t.Setenv(shutdownTimeout, "")
	a, _, stderr, built := newApp(t)
	ran := probe(a)

	code := a.Run(context.Background(), []string{"probe"})

	if code != process.ExitOK {
		t.Errorf("code = %d, want %d; stderr = %q", code, process.ExitOK, stderr)
	}
	if !*ran {
		t.Error("probe did not run")
	}
	if want := []string{"lifecycle config"}; !slices.Equal(*built, want) {
		t.Errorf("nodes built = %q, want %q", *built, want)
	}
}

func TestRun_InvalidShutdownTimeoutFailsTheCommand(t *testing.T) {
	t.Setenv(shutdownTimeout, "not-a-duration")
	a, _, stderr, _ := newApp(t)
	ran := probe(a)

	code := a.Run(context.Background(), []string{"probe"})

	if code != process.ExitFailure {
		t.Errorf("code = %d, want %d", code, process.ExitFailure)
	}
	if *ran {
		t.Error("probe ran, want it stopped by the Build")
	}
	if want := "app probe: lifecycle config: APP_SHUTDOWN_TIMEOUT"; !strings.HasPrefix(stderr.String(), want) {
		t.Errorf("stderr = %q, want it to start with %q", stderr, want)
	}
}
