package helper

import "encoding/json"

// Response struct
type HostingResponse struct {
	Hostname []string `json:",omitempty"`
	Code     string   `json:",omitempty"`
}

func (r *HostingResponse) GetJSON() string {
	reqJSON, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(reqJSON)
}
