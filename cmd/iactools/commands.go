package iactools

import (
	"errors"
	"fmt"
	"github.com/ArseniSkobelev/iactools/pkg/iactools"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// second lvl
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Creates and fills out files required to deploy infrastructure or services",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create command called")
	},
}

// third lvl
var kubernetesCmd = &cobra.Command{
	Use:     "kubernetes",
	Aliases: []string{"kube"},
	Short:   "Manages Kubernetes infrastructure",
	Args:    cobra.MaximumNArgs(6),
	Run: func(cmd *cobra.Command, args []string) {
		providerPrompt := promptui.Select{
			Label:        "Select CSP",
			Items:        []string{"Proxmox", "Azure", "AWS", "GCP"},
			HideSelected: true,
		}

		_, cloudProvider, err := providerPrompt.Run()

		if err != nil {
			fmt.Println("Something went wrong!")
			os.Exit(1)
		}

		kubernetesData := iactools.KubernetesNode{}
		cspData := iactools.CspData{}

		kubernetesData.CspData = cspData

		cspData.Provider = cloudProvider

		switch cloudProvider {
		case "Proxmox":
			providerAuthenticationPrompt := promptui.Select{
				Label:        "Select Proxmox authentication method",
				Items:        []string{"Username and password", "API Token"},
				HideSelected: true,
			}

			_, cloudProviderAuthentication, err := providerAuthenticationPrompt.Run()

			if err != nil {
				fmt.Println("Something went wrong!")
				os.Exit(1)
			}

			switch cloudProviderAuthentication {
			case "Username and password":
				creds := iactools.ProxmoxUsernameCredentials{}

				kubernetesData.CspData.Credentials = &creds

				proxmoxUserNamePrompt := promptui.Prompt{
					Label: "Proxmox username",
				}
				proxmoxUserName, err := proxmoxUserNamePrompt.Run()

				if err != nil {
					fmt.Println("Something went wrong!")
					os.Exit(1)
				}

				proxmoxPasswordPrompt := promptui.Prompt{
					Label: "Proxmox password",
					Mask:  '*',
				}
				proxmoxPassword, err := proxmoxPasswordPrompt.Run()

				if err != nil {
					fmt.Println("Something went wrong!")
					os.Exit(1)
				}

				validateBaseUri := func(input string) error {
					if !strings.Contains(input, "/api2/json") {
						return errors.New("Check the URI structure!")
					}
					return nil
				}

				baseUriPrompt := promptui.Prompt{
					Label: "Proxmox API base uri (must be formatted as following: https://<proxmox_host_ip>:8006/api2" +
						"/json)",
					Validate: validateBaseUri,
				}

				baseUri, err := baseUriPrompt.Run()

				if err != nil {
					fmt.Println("Something went wrong!")
					os.Exit(1)
				}

				creds.Username = proxmoxUserName
				creds.Password = proxmoxPassword
				creds.BaseUri = baseUri
			case "API Token":
				cspData.Credentials = iactools.ProxmoxTokenCredentials{}
			}
		case "Azure":
			fmt.Println("Yet to be implemented..")
			os.Exit(0)
		case "AWS":
			fmt.Println("Yet to be implemented..")
			os.Exit(0)
		case "GCP":
			fmt.Println("Yet to be implemented..")
			os.Exit(0)
		}

		vmData, err := iactools.GetVirtualMachineData()

		fmt.Println(vmData.Cores)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(kubernetesCmd)
}
