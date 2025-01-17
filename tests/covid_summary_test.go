package tests

import (
	"testing"

	"github.com/Captainistz/lmwn-intern-2025/models"
	"github.com/Captainistz/lmwn-intern-2025/repository"
	"github.com/Captainistz/lmwn-intern-2025/services"

	"github.com/stretchr/testify/assert"
)

func TestGetCovidSummary(t *testing.T) {
	covidCases, err := repository.GetCasesFromFile("covid_cases.json")
	assert.NoError(t, err)

	expectedSummary := models.CaseSummary{
		Province: map[string]int{
			"Amnatcharoen":      204,
			"Angthong":          203,
			"Ayutthaya":         225,
			"Bangkok":           213,
			"Betong":            195,
			"Buriram":           200,
			"Chachoengsao":      199,
			"Chainat":           203,
			"Chaiyaphum":        201,
			"Chanthaburi":       228,
			"Chiangmai":         223,
			"Chiangrai":         200,
			"Chonburi":          204,
			"Chumphon":          205,
			"Kalasin":           227,
			"Kamphaengphet":     235,
			"Kanchanaburi":      201,
			"Khonkaen":          196,
			"Krabi":             233,
			"Lampang":           197,
			"Lamphun":           174,
			"Loei":              200,
			"Lopburi":           201,
			"Maehongson":        185,
			"Mahasarakham":      211,
			"Mukdahan":          191,
			"Nakhonnayok":       192,
			"Nakhonpathom":      226,
			"Nakhonphanom":      225,
			"Nakhonratchasima":  198,
			"Nakhonsawan":       218,
			"Nakhonsithammarat": 172,
			"Nan":               227,
			"Narathiwat":        165,
			"Nongbualamphu":     201,
			"Nongkhai":          179,
			"Nonthaburi":        200,
			"Pathumthani":       169,
			"Pattani":           203,
			"Phangnga":          200,
			"Phatthalung":       223,
			"Phayao":            190,
			"Phetchabun":        209,
			"Phetchaburi":       200,
			"Phichit":           200,
			"Phitsanulok":       211,
			"Phrae":             205,
			"Phuket":            183,
			"Prachinburi":       200,
			"Prachuapkhirikhan": 207,
			"Ranong":            201,
			"Ratchaburi":        188,
			"Rayong":            195,
			"Roiet":             222,
			"Sakaeo":            216,
			"Sakonnakhon":       192,
			"Samutprakan":       181,
			"Samutsakhon":       227,
			"Samutsongkhram":    211,
			"Saraburi":          201,
			"Satun":             218,
			"Singburi":          219,
			"Sisaket":           208,
			"Songkhla":          227,
			"Sukhothai":         222,
			"Suphanburi":        221,
			"Suratthani":        197,
			"Surin":             196,
			"Tak":               207,
			"Trang":             231,
			"Trat":              191,
			"Ubonratchathani":   227,
			"Udonthani":         182,
			"Uthaithani":        204,
			"Uttaradit":         225,
			"Yala":              203,
			"Yasothon":          218,
			"N/A":               4213,
		},
		AgeGroup: models.CovidAgeGroup{
			Age0_30:  5175,
			Age31_60: 5414,
			Age_61p:  5266,
			Age_NA:   4145,
		},
	}

	summary, err := services.GetCovidSummary(covidCases)
	assert.NoError(t, err)
	assert.Equal(t, expectedSummary, summary)
}

func TestCovidSummaryWithNullAndNegativeAge(t *testing.T) {
	summary, err := services.GetCovidSummary(
		[]models.Case{
			{Age: nil},
			{Age: intPtr(-1)},
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, 2, summary.AgeGroup.Age_NA)
}

func TestOffByOneAge(t *testing.T) {
	summary, err := services.GetCovidSummary(
		[]models.Case{
			{Age: intPtr(29)},
			{Age: intPtr(30)},
			{Age: intPtr(60)},
			{Age: intPtr(61)},
		},
	)
	assert.NoError(t, err)
	expectedAgeGroup := models.CovidAgeGroup{
		Age0_30:  2,
		Age31_60: 1,
		Age_61p:  1,
	}
	assert.Equal(t, expectedAgeGroup, summary.AgeGroup)
}

func TestProvinceNull(t *testing.T) {
	covidCase := models.Case{Province: nil}
	summary, err := services.GetCovidSummary([]models.Case{covidCase})
	assert.NoError(t, err)
	assert.Equal(t, 1, summary.Province["N/A"])
}

func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
