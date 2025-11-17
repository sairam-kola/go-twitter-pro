package twitter

// TerminateAllConnectionsData is the data returned from terminating all connections
type TerminateAllConnectionsData struct {
	Attempted bool `json:"attempted"`
}

// TerminateAllConnectionsResponse is the response to terminating all connections
type TerminateAllConnectionsResponse struct {
	Data      *TerminateAllConnectionsData `json:"data"`
	RateLimit *RateLimit
}
