package iactools

const ProxmoxVirtualMachineTemplate = "resource \"proxmox_vm_qemu\" \"virtual_machines\" {\n  name" +
	"             = \"{{ hostname }}\"\n  qemu_os          = \"other\"\n  desc             = \"{{ description }}\"\n" +
	"  target_node      = \"{{ targetNode }}\"\n  os_type          = \"cloud-init\"\n  full_clone       = true\n  clone            = \"{{ template }}\"\n  memory           = {{ memory }}\n  cores            = {{ cores }}\n  ssh_user         = \"{{ sshUser }}\"\n  ciuser           = \"{{ sshUser }}\"\n  ipconfig0        = \"ip={{ ip }}/8,gw={{ gateway }}\"\n  cipassword       = \"{{ password }}\"\n  automatic_reboot = true\n  sshkeys          = \"{{ sshKeys }}\"\n\n  disk {\n    storage = \"{{ storageName }}\"\n    type    = \"{{ diskType }}\"\n    size    = {{ storageAmount }}\n  }\n\n  network {\n    bridge   = \"{{ networkBridgeType }}\"\n    model    = \"virtio\"\n    mtu      = 0\n    queues   = 0\n    rate     = 0\n    firewall = false\n  }\n}\n"
