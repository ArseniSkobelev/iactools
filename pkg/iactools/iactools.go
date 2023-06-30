package iactools

import (
	"errors"
	"fmt"
	"github.com/aymerick/raymond"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"strings"
)

func ShowSelectPrompt(label string, choices []string, hideSelected bool) (input string, err error) {
	prompt := promptui.Select{
		Label:        label,
		Items:        choices,
		HideSelected: hideSelected,
	}

	_, input, err = prompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	return input, nil
}

func ShowInputPrompt(label string, hideSelected bool) (input string, err error) {
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: hideSelected,
	}

	input, err = prompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	return input, nil
}

func ShowSecureInputPrompt(label string, hideSelected bool, mask rune) (input string, err error) {
	prompt := promptui.Prompt{
		Label:       label,
		HideEntered: hideSelected,
		Mask:        mask,
	}

	input, err = prompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	return input, nil
}

func GetCSP() (input string, err error) {
	input, _ = ShowSelectPrompt(
		"Select CSP",
		[]string{"Proxmox", "Azure", "AWS", "GCP"},
		true,
	)

	return input, nil
}

func GetCspCreds(cloudProvider string) (creds interface{}, err error) {
	cspCreds := CspData{}
	cspCreds.Provider = cloudProvider

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
			creds := ProxmoxUsernameCredentials{}

			proxmoxUserNamePrompt := promptui.Prompt{
				Label:       "Proxmox username",
				HideEntered: true,
			}
			proxmoxUserName, err := proxmoxUserNamePrompt.Run()

			if err != nil {
				fmt.Println("Something went wrong!")
				os.Exit(1)
			}

			proxmoxPasswordPrompt := promptui.Prompt{
				Label:       "Proxmox password",
				Mask:        '*',
				HideEntered: true,
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
				Validate:    validateBaseUri,
				HideEntered: true,
			}

			baseUri, err := baseUriPrompt.Run()

			if err != nil {
				fmt.Println("Something went wrong!")
				os.Exit(1)
			}

			creds.Username = proxmoxUserName
			creds.Password = proxmoxPassword
			creds.BaseUri = baseUri

			cspCreds.Credentials = creds
		case "API Token":
			fmt.Println("Yet to be implemented.. GetCspCreds -> proxmox -> api token")
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

	creds = cspCreds
	return creds, nil
}

func CreateKubernetes(k8sResourceType string) (err error) {
	// get CSP and get credentials
	csp, _ := GetCSP()
	_, _ = GetCspCreds(csp)

	if k8sResourceType == "Node" {
		switch csp {
		case "Proxmox":
			// handle create k8s node
			vmData, _ := GetVirtualMachineData()

			proxmoxTemplate, _ := ShowInputPrompt("Enter Proxmox template name", true)
			storageName, _ := ShowInputPrompt("Enter Proxmox storage name", true)
			targetNode, _ := ShowInputPrompt("Enter Proxmox target node", true)

			sshUserName, _ := ShowInputPrompt("Enter new VM username", true)
			sshPassword, _ := ShowSecureInputPrompt("Enter new VM password", true, '*')

			networkBridge, _ := ShowInputPrompt("Enter network bridge name", true)

			ctx := map[string]string{
				"hostname":          vmData.Hostname,
				"description":       "VM created by iactools",
				"targetNode":        targetNode,
				"template":          proxmoxTemplate,
				"memory":            vmData.Memory,
				"cores":             vmData.Cores,
				"sshUser":           sshUserName,
				"password":          sshPassword,
				"ip":                vmData.Ip,
				"gateway":           vmData.Gateway,
				"sshKeys":           vmData.PublicSshKey,
				"storageName":       storageName,
				"diskType":          "virtio",
				"storageAmount":     vmData.Storage,
				"networkBridgeType": networkBridge,
			}

			result, err := raymond.Render(ProxmoxVirtualMachineTemplate, ctx)
			if err != nil {
				panic("Please report a bug :)")
			}

			outputPath, _ := ShowInputPrompt("Please enter a path to store Terraform files ("+
				"Example: /c/users/user/home/iactools-terraform)", true)

			_ = createDirectory(outputPath)

			f, err := os.Create(fmt.Sprintf("%s/main.tf", outputPath))

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			_, err2 := f.WriteString(result)

			if err2 != nil {
				log.Fatal(err2)
			}

			fmt.Println("done")
		}
	}

	return nil
}

func createDirectory(path string) error {
	// Check if the directory already exists
	_, err := os.Stat(path)
	if err == nil {
		// Directory already exists, nothing to do
		return nil
	}

	// Create the directory recursively
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	return nil
}
