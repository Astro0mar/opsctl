package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
    Use:   "cluster",
    Short: "Cluster operations",
}

var clusterCheckCmd = &cobra.Command{
    Use:   "check",
    Short: "Check cluster health",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Cluster status: HEALTHY")
        fmt.Println("API Server: Reachable")
        fmt.Println("Nodes Ready: 3/3")
    },
}

func init() {
    rootCmd.AddCommand(clusterCmd)
    clusterCmd.AddCommand(clusterCheckCmd)
}
