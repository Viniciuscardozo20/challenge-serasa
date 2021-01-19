package controller

import (
	"challenge-serasa/api/database"
	"challenge-serasa/api/mainframe"
)

type Controller struct {
	coll          database.Collection
	mainframeHost string
}

func NewController(coll database.Collection, host string) *Controller {
	return &Controller{
		coll:          coll,
		mainframeHost: host,
	}
}

func (c *Controller) UpdateNegativations() error {
	nav, err := mainframe.GetNegativations(c.mainframeHost)
	if err != nil {
		return err
	}
	err = c.coll.SaveDocuments(nav)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetNegativationByCustomer(customerId string) (*[]mainframe.Negativation, error) {
	negativations, err := c.coll.GetDocuments(customerId, "customerDocument")
	if err != nil {
		return nil, err
	}
	return negativations, nil
}
