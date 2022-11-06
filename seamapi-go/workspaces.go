package seamapi

import "time"

type Workspace struct {
	WorkspaceID        string `json:"workspace_id"`
	Name               string `json:"name"`
	ConnectPartnerName string `json:"connect_partner_name"`
	IsSandbox          bool   `json:"is_sandbox"`
}
