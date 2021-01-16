package mainframe

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Negativation struct {
	CompanyDocument  string    `json:"companyDocument"`
	CompanyName      string    `json:"companyName"`
	CustomerDocument string    `json:"customerDocument"`
	Value            float64   `json:"value"`
	Contract         string    `json:"contract"`
	DebtDate         time.Time `json:"debtDate"`
	InclusionDate    time.Time `json:"inclusionDate"`
}

type Negativations struct {
	Negativations []Negativations `json:"negativations"`
}

func GenerateNegativation(companyDocument, companyName, customerDocument string, value float64, contract string, debtDate, inclusionDate time.Time) *Negativation {
	return &Negativation{
		CompanyDocument:  companyDocument,
		CompanyName:      companyName,
		CustomerDocument: customerDocument,
		Value:            value,
		Contract:         contract,
		DebtDate:         debtDate,
		InclusionDate:    inclusionDate,
	}
}

func GetNegativations(host string) ([]Negativation, error) {
	res, err := http.Get(host)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect mainframe service")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}
	var negativations []Negativation
	err = json.Unmarshal(body, &negativations)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal")
	}
	return negativations, nil
}
