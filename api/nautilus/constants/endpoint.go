package constant

type EndPointURL string

const (
	// ユーザ管理サーバ
	UsersServiceEndPoint EndPointURL = "http://localhost:8880/users/v1"

	// 通知管理APIサーバ
	NotificationsServiceEndPoint EndPointURL = "http://notifications-api:8080/notifications/v1"

	// 配送管理APIサービス
	DeliveriesServiceEndPoint EndPointURL = "http://localhost:8883/deliveries/v1"
)
