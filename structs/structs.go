package structs

// Group Structs
type Groups struct {
	Groups []Group
}
type Group struct {
	Name string
	Vms  []Vm
}
type Vm struct {
	Name string
	Id   int
}
type Vms struct {
	Vms []Vm
}

// Vms Structs
type Information struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Vmid   int    `json:"vmid"`
	Cpus   int    `json:"cpus"`
	Memory int    `json:"maxmem"`
}
type VmsInfo struct {
	Data []Information `json:"data"`
}
type VmInfo struct {
	Data Information `json:"data"`
}

// Nodes Structs
type Node struct {
	Memory  map[string]int `json:"memory"`
	Loadavg []string       `json:"loadavg"`
	Uptime  int            `json:"uptime"`
	Cpu     Cpu            `json:"cpuinfo"`
}
type Cpu struct {
	Cpus int `json:"cpus"`
}
type NodeInfo struct {
	Node Node `json:"data"`
}

// Inventory Structs
type Inventory struct {
	Vms []InventoryInfo `json:"data"`
}
type InventoryInfo struct {
	Name string `json:"name"`
	Id   int    `json:"vmid"`
}

// Configuration Structs
type Security struct {
	User    string
	Realm   string
	Tokenid string
	Token   string
}
type Server struct {
	Node     string
	Ip       string
	Security Security
}
type Config struct {
	Config []Server
}
