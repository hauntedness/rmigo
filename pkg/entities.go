package pkg

type Item struct {
	ID     string    `json:"id"`
	Desc   string    `json:"desc"`
	Parent *Item     `json:"parent"`
	Dir    Directory `json:"dir"`
}

type Story struct {
	Item
	Acceptance string `json:"acceptance"`
	Attachment string `json:"attachment"`
	Owner      string `json:"owner"`
	Status     string `json:"status"`
}

type Defect struct {
	Item
	Acceptance string `json:"acceptance"`
	Attachment string `json:"attachment"`
	Owner      string `json:"owner"`
	Status     string `json:"status"`
}

type Directory *Node

type Sprint struct {
	*Node
}

type Project struct {
	*Node
}
