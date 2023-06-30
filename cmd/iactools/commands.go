package iactools

import (
	"github.com/ArseniSkobelev/iactools/pkg/iactools"
	"github.com/spf13/cobra"
)

// second lvl
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Creates and fills out files required to deploy infrastructure or services",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// third lvl
var kubernetesCreateCommand = &cobra.Command{
	Use:     "kubernetes",
	Aliases: []string{"kube"},
	Short:   "Manages Kubernetes infrastructure",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		k8sResourceType, _ := iactools.ShowSelectPrompt(
			"Select Kubernetes resource type",
			[]string{"Node", "Cluster"},
			true,
		)

		_ = iactools.CreateKubernetes(k8sResourceType)
	},
}

var kubernetesNodeCreateCommand = &cobra.Command{
	Use:     "node",
	Aliases: []string{"n"},
	Short:   "Create Kubernetes node",
	Args:    cobra.MaximumNArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		_ = iactools.CreateKubernetes("Node")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(kubernetesCreateCommand)
	kubernetesCreateCommand.AddCommand(kubernetesNodeCreateCommand)
}
