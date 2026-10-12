// Command app is the template's command-line program. Its main sequence is
// the signal context and one call into the composition root, internal/app.
package main

import (
	"os"

	"github.com/standards-lab/go-cli-sdk"
	"github.com/standards-lab/go-core/process"

	"github.com/standards-lab/go-cli-sdk-template/template/internal/app"
)

func main() {
	os.Exit(run())
}

// run runs app and returns its exit code. It is separate from main so the
// deferred stop releases the signal registration before os.Exit, which
// skips deferred calls.
func run() int {
	ctx, stop := process.SignalContext()
	defer stop()
	return app.New(cli.Streams{Stdin: os.Stdin, Stdout: os.Stdout, Stderr: os.Stderr}).Run(ctx, os.Args[1:])
}
