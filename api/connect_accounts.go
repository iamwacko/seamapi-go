package seamapi

import "time"

type ConnectedAccount struct {
	ConnectedAccountID string    `json:"connected_account_id"`
	CreatedAt          time.Time `json:"created_at"`
	UserIdentifier     struct {
		Email string `json:"email"`
	} `json:"user_identifier"`
	AccountType string        `json:"account_type"`
	Errors      []interface{} `json:"errors"`
} 
