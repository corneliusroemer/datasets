package datasets

import (
	"fmt"
	"os"

	// cmdflags "datasets_cli/v2/datasets/flags"

	"github.com/gosuri/uiprogress"
	// openapi "datasets/openapi/v2"
)

var (
	// AppVersion is the application version string whose value is set at build time
	progress *uiprogress.Progress
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// fmt.Fprintln(os.Stderr, "DEBUG: Starting Execute()")
	// exitval := 0
	//err := rootCmd.Execute()
	// fmt.Fprintln(os.Stderr, "DEBUG: rootCmd.Execute() returned")
	// if err != nil {
	// 	exitval = 1
	// }
	// if userMessage != "" {
	// 	exitval = 1
	// }
	fmt.Fprintln(os.Stderr, "DEBUG: About to call os.Exit()")
	os.Exit(0)
	fmt.Fprintln(os.Stderr, "DEBUG: Code should never reach here")
}
