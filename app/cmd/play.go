package cmd

import (
	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground,but runing at our application context",
	Run:   runplay,
}

func runplay(cmd *cobra.Command, args []string) {

}
