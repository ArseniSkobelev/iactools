package iactools

type KubernetesNode struct {
	IsMaster bool
	CspData  CspData
}

type CspData struct {
	Provider    string
	Credentials interface{}
}

type ProxmoxUsernameCredentials struct {
	BaseUri  string
	Username string
	Password string
}

type ProxmoxTokenCredentials struct {
	BaseUri string
	TokenId string
	Secret  string
}

type VirtualMachine struct {
	Memory       string
	Storage      string
	Hostname     string
	Ip           string
	Gateway      string
	Cores        string
	PublicSshKey string
}
