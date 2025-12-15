package ingest

import "os/exec"

// ================== GIT OPERATIONS ==================
func CloneRepo(repoURL, targetDir string) error {
	// ========== SHALLOW CLONE FOR EFFICIENCY ==========
	// Using --depth 1 to clone only the latest commit
	cmd := exec.Command("git", "clone", "--depth", "1", repoURL, targetDir)
	return cmd.Run()
}
