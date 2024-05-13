package types

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
type VmInventory struct {
	Vms []VmInventoryInfo `json:"data"`
}
type VmInventoryInfo struct {
	Name string `json:"name"`
	Id   int    `json:"vmid"`
}
type LxcInventory struct {
	Lxc []LxcInventoryInfo `json:"data"`
}
type LxcInventoryInfo struct {
	Name string `json:"name"`
	Id   any    `json:"vmid"`
}

// Lxc Structs
type LxcInformation struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Vmid   any    `json:"vmid"`
	Cpus   int    `json:"cpus"`
	Memory int    `json:"maxmem"`
}
type LxcsInfo struct {
	Data []LxcInformation `json:"data"`
}
type LxcInfo struct {
	Data LxcInformation `json:"data"`
}
