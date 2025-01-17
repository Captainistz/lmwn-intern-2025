package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Captainistz/lmwn-intern-2025/models"
)

type CovidCaseRepository struct {
	repo string
}

func NewCovidCaseRepository() (*CovidCaseRepository, error) {
	repo, ok := os.LookupEnv("COVID_CASES_API_URL")
	if !ok {
		return nil, fmt.Errorf("ENV:COVID_CASES_API_URL not found")
	}
	return &CovidCaseRepository{
		repo: repo,
	}, nil
}

func GetCasesFromFile(fileName string) ([]models.Case, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return readBody(file)
}

func (ctx *CovidCaseRepository) GetCases() ([]models.Case, error) {
	res, err := http.Get(ctx.repo)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	return readBody(res.Body)
}

func readBody(r io.Reader) ([]models.Case, error) {
	body, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var cases models.CovidCases
	if err := json.Unmarshal(body, &cases); err != nil {
		return nil, err
	}

	return cases.Data, nil
}
