package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "opsctl",
    Short: "Simple DevOps CLI written in Go",
    Long:  "opsctl is a beginner-friendly DevOps CLI for Kubernetes diagnostics.",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
