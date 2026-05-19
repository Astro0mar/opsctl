package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var podsCmd = &cobra.Command{
    Use:   "pods",
    Short: "Pod operations",
}

var failedPodsCmd = &cobra.Command{
    Use:   "failed",
    Short: "List failed pods",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("NAMESPACE     POD                  STATUS")
        fmt.Println("default       nginx-xyz            CrashLoopBackOff")
        fmt.Println("monitoring    grafana-0            Error")
    },
}

func init() {
    rootCmd.AddCommand(podsCmd)
    podsCmd.AddCommand(failedPodsCmd)
}
