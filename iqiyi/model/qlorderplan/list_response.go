package qlorderplan

import "donson.com.cn/draining/internal/pkg/iqiyi/model"

type ListResponse struct {
	PageInfo model.PageInfo `json:"pageInfo"`
	List     []List         `json:"list"`
}

type List struct {
	ID                       int64         `json:"id"`
	AdvertiserID             int64         `json:"advertiserId"`
	OrderGroupID             int64         `json:"orderGroupId"`
	PromotionGoal            string        `json:"promotionGoal"`
	DeliveryMedia            string        `json:"deliveryMedia"`
	WebsiteURL               string        `json:"websiteUrl"`
	WebsiteAppName           interface{}   `json:"websiteAppName"`
	AppleID                  interface{}   `json:"appleId"`
	IosAppName               interface{}   `json:"iosAppName"`
	ApkURL                   string        `json:"apkUrl"`
	AndroidApkName           string        `json:"androidApkName"`
	AndroidAppName           string        `json:"androidAppName"`
	AndroidAppDetailURL      string        `json:"androidAppDetailUrl"`
	WakeUpType               interface{}   `json:"wakeUpType"`
	DeepIOSAppName           interface{}   `json:"deepIOSAppName"`
	DeeplinkIOS              interface{}   `json:"deeplinkIOS"`
	IosSchema                interface{}   `json:"iosSchema"`
	IosDeepLinkDetailURL     interface{}   `json:"iosDeepLinkDetailUrl"`
	DeepAndroidAppName       interface{}   `json:"deepAndroidAppName"`
	ApkName                  interface{}   `json:"apkName"`
	Deeplink                 interface{}   `json:"deeplink"`
	AndroidDeepLinkDetailURL interface{}   `json:"androidDeepLinkDetailUrl"`
	AdTypeID                 interface{}   `json:"adTypeId"`
	Status                   string        `json:"status"`
	DisplayStatus            string        `json:"displayStatus"`
	TargetList               []TargetList  `json:"targetList"`
	OptimizedType            string        `json:"optimizedType"`
	SettleType               string        `json:"settleType"`
	Price                    float64       `json:"price"`
	DayBudget                int           `json:"dayBudget"`
	ConversionID             int64         `json:"conversionId"`
	ConversionTrackingType   string        `json:"conversionTrackingType"`
	ConversionTarget         string        `json:"conversionTarget"`
	ConversionTargetPrice    float64       `json:"conversionTargetPrice"`
	ConversionOptimize       string        `json:"conversionOptimize"`
	BiddingStrategy          string        `json:"biddingStrategy"`
	DeepConversionTargetID   int           `json:"deepConversionTargetId"`
	DeepConversionPrice      interface{}   `json:"deepConversionPrice"`
	DeepConversionTarget     string        `json:"deepConversionTarget"`
	CreativeDisplayType      interface{}   `json:"creativeDisplayType"`
	AdCategoryID             int           `json:"adCategoryId"`
	OrderPlanName            string        `json:"orderPlanName"`
	AutoExpansionType        int           `json:"autoExpansionType"`
	NoAutoExpansionTargets   []int         `json:"noAutoExpansionTargets"`
	AppDownloadType          string        `json:"appDownloadType"`
	AdTagList                []interface{} `json:"adTagList"`
	SourceType               string        `json:"sourceType"`
	TargetRoi                interface{}   `json:"targetRoi"`
	CreateTime               string        `json:"createTime"`
	UpdateTime               string        `json:"updateTime"`
	RetentionRatio           float64       `json:"retentionRatio"`
}

type ValueList struct {
	ID         int         `json:"id"`
	Name       string      `json:"name"`
	ExtendInfo interface{} `json:"extendInfo"`
}

type TargetList struct {
	Type      int         `json:"type"`
	Name      string      `json:"name"`
	ValueList []ValueList `json:"valueList"`
	IsAnti    bool        `json:"isAnti"`
}
