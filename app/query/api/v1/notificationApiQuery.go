package v1

type NotificationBody struct {
	SendUserId      int64             `json:"id"`
	SendNickName    string            `json:"sendNickName"`
	ReceiveUserId   int64             `json:"receiveUid"`
	ReceiveNickname string            `json:"receiveNickname"`
	Subject         string            `json:"subject"`
	Content         string            `json:"content"`
	ExtData         map[string]string `json:"extData"`
	ExpiresAt       int64             `json:"expiresAt"`
}

type NotificationPayload struct {
	NotificationType int              `json:"notificationType"`
	IsForce          bool             `json:"isForce"`
	Body             NotificationBody `json:"body"`
}
