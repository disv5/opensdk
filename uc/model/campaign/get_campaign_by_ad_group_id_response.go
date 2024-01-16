package campaign

type GetCampaignByAdGroupIdResponse struct {
	AdGroupCampaigns []AdGroupCampaign `json:"adGroupCampaigns"`
}

type Objectives struct {
	TargetURL       string      `json:"targetUrl"`
	SchemeURL       string      `json:"schemeUrl"`
	SiteID          int64       `json:"siteId"`
	ConvertType     int         `json:"convertType"`
	AdConvertID     int64       `json:"adConvertId"`
	DeepConvertType int         `json:"deepConvertType"`
	AdmFixedUlk     string      `json:"admFixedUlk"`
	GroupTargetUrl  []TargetURL `json:"groupTargetUrl"`
}

type TargetURL struct {
	TargetURL string `json:"targetUrl"`
}

type Targetings struct {
	AudienceTargeting  string        `json:"audienceTargeting"`
	AllRegion          string        `json:"allRegion"`
	RegionPeople       string        `json:"regionPeople"`
	Gender             string        `json:"gender"`
	Age                string        `json:"age"`
	Platform           string        `json:"platform"`
	NetworkEnv         string        `json:"networkEnv"`
	Interest           []interface{} `json:"interest"`
	App                []interface{} `json:"app"`
	IntelliTargeting   int64         `json:"intelliTargeting"`
	ConvertFilter      int64         `json:"convertFilter"`
	TargetingPackageID int64         `json:"targetingPackageId"`
	AutoExpand         int           `json:"autoExpand"`
	AutoTargeting      int           `json:"autoTargeting"`
	ManualTargeting    int64         `json:"manualTargeting"`
	Region             []interface{} `json:"region"`
	Appcategory        []interface{} `json:"appcategory"`
}
type Schedule struct {
	StartDate int64  `json:"startDate"`
	EndDate   int64  `json:"endDate"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
	Sunday    string `json:"sunday"`
}
type Bids struct {
	Bid                 string `json:"bid"`
	OptBid              string `json:"optBid"`
	BidStage            int64  `json:"bidStage"`
	DeepBid             string `json:"deepBid"`
	DeepConvertStrategy int64  `json:"deepConvertStrategy"`
	SkipFirstStage      int64  `json:"skipFirstStage"`
	RoiTarget           string `json:"roiTarget"`
	RoiWindow           int64  `json:"roiWindow"`
}
type CampaignType struct {
	Index                 int           `json:"index"`
	Type                  int           `json:"type"`
	Paused                bool          `json:"paused"`
	State                 int           `json:"state"`
	ID                    int64         `json:"id"`
	AdGroupID             int64         `json:"adGroupId"`
	Name                  string        `json:"name"`
	OptTarget             int           `json:"optTarget"`
	Delivery              int           `json:"delivery"`
	Objectives            []Objectives  `json:"objectives"`
	TrackArgs             string        `json:"trackArgs"`
	Targetings            []Targetings  `json:"targetings"`
	Budget                string        `json:"budget"`
	Schedule              Schedule      `json:"schedule"`
	ChargeType            int           `json:"chargeType"`
	Bids                  []Bids        `json:"bids"`
	RefuseReason          string        `json:"refuseReason"`
	AnxtStatus            int           `json:"anxtStatus"`
	EnableAnxt            bool          `json:"enableAnxt"`
	DeliveryMode          int           `json:"deliveryMode"`
	ComponentFlag         int           `json:"componentFlag"`
	NegativeWords         []interface{} `json:"negativeWords"`
	ExactNegativeWords    []interface{} `json:"exactNegativeWords"`
	ExactRatio            string        `json:"exactRatio"`
	PhraseRatio           string        `json:"phraseRatio"`
	AbroadRatio           string        `json:"abroadRatio"`
	AutoTargetingWinfoBid string        `json:"autoTargetingWinfoBid,omitempty"`
	AutoTargetingWinfoURL string        `json:"autoTargetingWinfoUrl,omitempty"`
}
type AdGroupCampaign struct {
	AdGroupID     int64          `json:"adGroupId"`
	CampaignTypes []CampaignType `json:"campaignTypes"`
}
