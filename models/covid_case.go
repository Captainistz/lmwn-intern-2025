package models

type CovidCases struct {
	Data []Case `json:"Data"`
}

type Case struct {
	ConfirmDate    *string `json:"ConfirmDate"`
	No             *int    `json:"No"`
	Age            *int    `json:"Age"`
	Gender         *string `json:"Gender"`
	GenderEn       *string `json:"GenderEn"`
	Nation         *string `json:"Nation"`
	NationEn       *string `json:"NationEn"`
	Province       *string `json:"Province"`
	ProvinceId     *int    `json:"ProvinceId"`
	District       *string `json:"District"`
	ProvinceEn     *string `json:"ProvinceEn"`
	StatQuarantine *int    `json:"StatQuarantine"`
}

type CovidAgeGroup struct {
	Age0_30  int `json:"0-30"`
	Age31_60 int `json:"31-60"`
	Age_61p  int `json:"61+"`
	Age_NA   int `json:"N/A"`
}

type CaseSummary struct {
	Province map[string]int `json:"Province"`
	AgeGroup CovidAgeGroup  `json:"AgeGroup"`
}
