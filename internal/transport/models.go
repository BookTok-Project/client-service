package transport

type subscribeRequest struct {
	SubscriberID int64 `json:"subscriber_id"`
	SubscribeeID int64 `json:"subscribee_id"`
}

type errorResponse struct {
	Error string `json:"error"`
}
