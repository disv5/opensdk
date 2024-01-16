package enum

// 广告状态
type PromotionStatus string

const (
	// 不限
	NOT_DELETED PromotionStatus = "NOT_DELETED"
	// 不限（包含已删除）
	ALL                      PromotionStatus = "ALL"
	PROMOTION_STATUS_DELETED PromotionStatus = "PROMOTION_STATUS_DELETED"
)
