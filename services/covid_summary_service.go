package services

import (
	"runtime"
	"sync"

	"github.com/Captainistz/lmwn-intern-2025/models"
	"github.com/Captainistz/lmwn-intern-2025/repository"
)

type summaryProcessor struct {
	workers    int
	casesChan  chan []models.Case
	resultChan chan *models.CaseSummary
	wg         *sync.WaitGroup
}

func newSummaryProcessor(workers int) *summaryProcessor {
	return &summaryProcessor{
		workers:    workers,
		casesChan:  make(chan []models.Case, workers),
		resultChan: make(chan *models.CaseSummary, workers),
		wg:         &sync.WaitGroup{},
	}
}

func (p *summaryProcessor) process(covidCases []models.Case) models.CaseSummary {
	p.startWorkers()
	p.splitCases(covidCases)
	return p.aggregateResults()
}

func (p *summaryProcessor) startWorkers() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.processCases()
	}
}

func (p *summaryProcessor) splitCases(covidCases []models.Case) {
	chunkSize := (len(covidCases) + p.workers - 1) / p.workers
	for i := 0; i < len(covidCases); i += chunkSize {
		end := min(i+chunkSize, len(covidCases))
		p.casesChan <- covidCases[i:end]
	}
	close(p.casesChan)
}

func (p *summaryProcessor) processCases() {
	defer p.wg.Done()
	result := &models.CaseSummary{
		Province: make(map[string]int),
	}
	for cases := range p.casesChan {
		for _, covidCase := range cases {
			updateProvinceCounts(&result.Province, covidCase)
			updateAgeGroups(&result.AgeGroup, covidCase)
		}
	}
	p.resultChan <- result
}

func (p *summaryProcessor) aggregateResults() models.CaseSummary {
	summary := models.CaseSummary{
		Province: make(map[string]int),
	}
	for i := 0; i < p.workers; i++ {
		result := <-p.resultChan
		mergeSummaries(&summary, result)
	}
	p.wg.Wait()
	return summary
}

func updateProvinceCounts(provinceMap *map[string]int, covidCase models.Case) {
	province := "N/A"
	if covidCase.Province != nil {
		province = *covidCase.Province
	}
	(*provinceMap)[province]++
}

func updateAgeGroups(ageGroup *models.CovidAgeGroup, covidCase models.Case) {
	switch {
	case covidCase.Age == nil || *covidCase.Age < 0:
		ageGroup.Age_NA++
	case *covidCase.Age <= 30:
		ageGroup.Age0_30++
	case *covidCase.Age <= 60:
		ageGroup.Age31_60++
	default:
		ageGroup.Age_61p++
	}
}

func mergeSummaries(target *models.CaseSummary, source *models.CaseSummary) {
	for key, count := range source.Province {
		target.Province[key] += count
	}
	target.AgeGroup.Age_NA += source.AgeGroup.Age_NA
	target.AgeGroup.Age0_30 += source.AgeGroup.Age0_30
	target.AgeGroup.Age31_60 += source.AgeGroup.Age31_60
	target.AgeGroup.Age_61p += source.AgeGroup.Age_61p
}

func GetCases() ([]models.Case, error) {
	covidCaseRepository, err := repository.NewCovidCaseRepository()
	if err != nil {
		return nil, err
	}
	return covidCaseRepository.GetCases()
}

func GetCovidSummary(covidCases []models.Case) (models.CaseSummary, error) {
	processor := newSummaryProcessor(runtime.NumCPU())
	return processor.process(covidCases), nil
}
