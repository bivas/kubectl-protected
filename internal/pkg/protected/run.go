package protected

import (
	"fmt"
	"os"
	"regexp"

	"github.com/bivas/kubectl-protected/internal/pkg/kubectl"
)

// RunCheck executes the cluster check logic
func RunCheck(opts *Options) error {
	// Resolve home directory in the file path
	expandedPath := os.ExpandEnv(opts.ProtectedFilePath)

	file, err := os.Open(expandedPath)
	if err != nil {
		fmt.Println("⚠️  Unable to open protected clusters file")
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	// Read the protected clusters file
	protectedClusters, err := ReadProtectedPatterns(file)
	if err != nil {
		fmt.Println("⚠️  Unable to read protected clusters file")
		return err
	}

	// Get the current Kubernetes cluster
	currentCluster, err := kubectl.GetCurrentCluster()
	if err != nil {
		fmt.Println("❌ Failed to retrieve the current cluster")
		return err
	}

	// Check if the cluster is protected
	for _, pattern := range protectedClusters {
		match, err := matchClusterPattern(currentCluster, pattern)
		if err != nil {
			fmt.Printf("⚠️  Invalid pattern: %q\n", pattern)
			return err
		}
		if match {
			fmt.Println("🚨 WARNING: You are connected to a PROTECTED cluster", currentCluster)
			if opts.SilenceOnProtected {
				return nil
			}
			return fmt.Errorf("cluster (%s) is protected", currentCluster)
		}
	}

	// If no match was found
	fmt.Printf("✅ Current cluster (%s) is not protected.\n", currentCluster)
	return nil
}

// matchClusterPattern checks if the current cluster matches a given pattern
func matchClusterPattern(cluster, pattern string) (bool, error) {
	// Direct match
	if cluster == pattern {
		return true, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.MatchString(cluster), nil
}
