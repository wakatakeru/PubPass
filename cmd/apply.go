package cmd

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var ApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply config to authorized_keys",
	Run: func(cmd *cobra.Command, args []string) {
		glog.Info("[TODO] Apply Command")
	},
}

// If `--all`, we applies all users configurations without user specification.
// If no flags and no username, we shows error.
// If `pubpass apply username`, we apply configuration to a user.
