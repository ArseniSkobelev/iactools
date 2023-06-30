package iactools

import (
	"fmt"
	"github.com/ArseniSkobelev/iactools/pkg/iactools"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var rootCmd = &cobra.Command{
	Use:     "iactools",
	Version: version,
	Short:   "iactools - Streamline manual IaC",
	Long:    `iactools is a CLI tool to streamline manual editing and provisioning of IaC files`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			// handle 0 args
			command, _ := iactools.ShowSelectPrompt(
				"Select command",
				[]string{"Create"},
				true,
			)

			resourceType, _ := iactools.ShowSelectPrompt(
				"Select resource type",
				[]string{"Kubernetes"},
				true,
			)

			if resourceType == "Kubernetes" {
				if command == "Create" {
					// handle k8s resource type
					k8sResourceType, _ := iactools.ShowSelectPrompt(
						"Select Kubernetes resource",
						[]string{"Node", "Cluster"},
						true,
					)

					_ = iactools.CreateKubernetes(k8sResourceType)
				}
			}
			os.Exit(0)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Exception: '%s'", err)
		os.Exit(1)
	}
}
