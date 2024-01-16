package report

// Metrics 指标数据
type Metrics struct {
	Impression          int64   `json:"impression"`
	Click               int64   `json:"click"`
	Ctr                 float64 `json:"ctr"`
	Cost                float64 `json:"cost"`
	Trueview            float64 `json:"trueview"`
	DownloadStart       int64   `json:"downloadStart"`
	DownloadFinish      int64   `json:"downloadFinish"`
	Install             int64   `json:"install"`
	AvgCpm              float64 `json:"avgCpm"`
	AvgCpc              float64 `json:"avgCpc"`
	AvgCpv              float64 `json:"avgCpv"`
	ConversionCount     int64   `json:"conversionCount"`
	ConversionCost      float64 `json:"conversionCost"`
	DeepConversionCount float64 `json:"deepConversionCount"`
	DeepConversionCost  float64 `json:"deepConversionCost"`
	Roi                 float64 `json:"roi"`
	DownloadFinishCost  float64 `json:"downloadFinishCost"`
	DownloadFinishRate  float64 `json:"downloadFinishRate"`
	DownloadStartCost   float64 `json:"downloadStartCost"`
	DownloadStartRate   float64 `json:"downloadStartRate"`
	InstallCost         float64 `json:"installCost"`
	InstallRate         float64 `json:"installRate"`
	ActiveCnt           int64   `json:"activeCnt"`
	ActiveCost          float64 `json:"activeCost"`
	ActiveRate          float64 `json:"activeRate"`
	RegisterCnt         int64   `json:"registerCnt"`
	RegisterCost        float64 `json:"registerCost"`
	RegisterRate        float64 `json:"registerRate"`
	RetentionCnt        int64   `json:"retentionCnt"`
	RetentionCost       float64 `json:"retentionCost"`
	PaymentCnt          int64   `json:"paymentCnt"`
	PaymentCost         float64 `json:"paymentCost"`
	VisitCnt            int64   `json:"visitCnt"`
	VisitCost           float64 `json:"visitCost"`
}
