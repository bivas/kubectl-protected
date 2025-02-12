package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/bivas/kubectl-protected/internal/pkg/protected"
)

func main() {
	opts := &protected.Options{}
	// Create the Cobra command
	var rootCmd = &cobra.Command{
		Use:           "protected",
		Short:         "Check if the current Kubernetes cluster is protected",
		Args:          cobra.NoArgs,
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return protected.RunCheck(opts)
		},
	}

	// Define flags
	rootCmd.Flags().BoolVar(
		&opts.SilenceOnProtected,
		"silence-error-on-protected",
		false,
		"Exit with error if the cluster is protected")
	rootCmd.Flags().StringVar(
		&opts.ProtectedFilePath,
		"protected-file-path",
		protected.DefaultProtectedFilePath,
		"Path to the file containing the list of protected clusters")

	// Execute the command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("❌ Error:", err)
		os.Exit(1)
	}
}
