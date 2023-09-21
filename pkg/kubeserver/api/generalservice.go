package api

type GeneralService struct {
	ObjectTypeMeta
	Params    map[string]interface{} `json:"params,omitempty"`
	ReleaseId string                 `json:"releaseId"`
	ServiceId string                 `json:"serviceId,omitempty"`
	GeneralServiceStatus
}

type GeneralServiceStatus struct {
	Phase        string             `json:"phase,omitempty"`
	Reason       string             `json:"reason,omitempty"`
	ExternalInfo GeneralServiceInfo `json:"externalInfo"`
	TryTimes     int32              `json:"tryTimes"`
}

type GeneralServiceInfo struct {
	OnecloudExternalInfoBase
	PrimaryKey string `json:"primaryKey"`
	ExternalId string `json:"externalId,omitempty"`
}
