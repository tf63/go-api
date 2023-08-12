package entity

// StatusError型の宣言
type StatusError int

// Errorを実装
func (e StatusError) Error() string {
	return "error"
}

// 構造体を列挙
const (
	STATUS_SERVICE_UNAVAILABLE StatusError = 503 // 503
	STATUS_NOT_FOUND           StatusError = 404 // 404
)
