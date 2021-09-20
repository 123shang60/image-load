package register

type NodeInfo struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Port string `json:"port"`
}

type RegistResult struct {
	Code int32  `json:"code"`
	Data string `json:"data"`
}
