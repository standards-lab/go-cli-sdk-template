//go:build integration

package integration_test

import (
	"testing"

	"github.com/standards-lab/go-cli-sdk-template/template/integration"
)

func TestMain(m *testing.M) { integration.Main(m) }
