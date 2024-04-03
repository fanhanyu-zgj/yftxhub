package cmd

import (
	"yftxhub/pkg/cache"
	"yftxhub/pkg/console"

	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache Management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "Clear cache",
	Run:   runcCacheClear,
}

func init() {
	CmdCache.AddCommand(CmdCacheClear)
}

func runcCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}
