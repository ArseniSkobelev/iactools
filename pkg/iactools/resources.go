package iactools

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"strconv"
)

func GetVirtualMachineData() (vmData VirtualMachine, err error) {
	vmData = VirtualMachine{}

	validateNumber := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Please enter a number of cores!")
		}
		return nil
	}

	cpuCoresPrompt := promptui.Prompt{
		Label:     "Enter amount of cores desired",
		Validate:  validateNumber,
		Default:   "2",
		AllowEdit: true,
	}

	cpuCores, err := cpuCoresPrompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	amountOfMemoryPrompt := promptui.Prompt{
		Label:     "Enter amount of memory (RAM) (in MB)",
		Validate:  validateNumber,
		Default:   "2048",
		AllowEdit: true,
	}

	memory, err := amountOfMemoryPrompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	storagePrompt := promptui.Prompt{
		Label:     "Enter amount of storage (disk) (in GB)",
		Validate:  validateNumber,
		Default:   "32",
		AllowEdit: true,
	}

	storage, err := storagePrompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	hostnamePrompt := promptui.Prompt{
		Label:     "Enter the hostname of the machine",
		Default:   "virtual-machine",
		AllowEdit: true,
	}

	hostname, err := hostnamePrompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	ipPrompt := promptui.Prompt{
		Label: "Enter the IP of the machine",
	}

	ip, err := ipPrompt.Run()

	if err != nil {
		fmt.Println("Something went wrong!")
		os.Exit(1)
	}

	vmData.Cores = cpuCores
	vmData.Memory = memory
	vmData.Hostname = hostname
	vmData.Storage = storage
	vmData.Ip = ip

	return vmData, nil
}
