package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllFormatsConvertable(t *testing.T) {
	assertions := []traitAssertion{
		assertStdoutLengthGreaterThan(1000),
		assertSuccessfulReturnCode,
	}

	tests := []struct {
		to       string
		from     string
		template string
		env      map[string]string
	}{
		{to: "buildx-json", from: "spdx-json"},
		{to: "buildx-json", from: "cyclonedx-json"},
		{to: "spdx-json", from: "buildx-json"},
		{to: "template", from: "buildx-json", template: "test-fixtures/csv.template"},
		{to: "spdx-json", from: "cyclonedx-json"},
		{to: "cyclonedx-json", from: "buildx-json"},
		{to: "cyclonedx-json", from: "spdx-json"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("from %s to %s", test.from, test.to), func(t *testing.T) {
			sbomArgs := []string{"dir:./test-fixtures/image-pkg-coverage", "-o", test.from}
			cmd, stdout, stderr := runBuildx(t, test.env, sbomArgs...)
			if cmd.ProcessState.ExitCode() != 0 {
				t.Log("STDOUT:\n", stdout)
				t.Log("STDERR:\n", stderr)
				t.Log("COMMAND:", strings.Join(cmd.Args, " "))
				t.Fatalf("failure executing buildx creating an sbom")
				return
			}

			tempDir := t.TempDir()
			sbomFile := filepath.Join(tempDir, "sbom.json")
			require.NoError(t, os.WriteFile(sbomFile, []byte(stdout), 0666))

			convertArgs := []string{"convert", sbomFile, "-o", test.to}
			if test.template != "" {
				convertArgs = append(convertArgs, "--template", test.template)
			}
			cmd, stdout, stderr = runBuildx(t, test.env, convertArgs...)
			for _, traitFn := range assertions {
				traitFn(t, stdout, stderr, cmd.ProcessState.ExitCode())
			}
			logOutputOnFailure(t, cmd, stdout, stderr)
		})
	}
}
