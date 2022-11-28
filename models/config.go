package models

type ConfigJsonFormat struct {
	Route string   `json:"route"`
	Auth  bool     `json:"auth"`
	Nodes []string `json:"nodes"`
}
