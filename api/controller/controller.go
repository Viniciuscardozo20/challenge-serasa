package controller

import (
	"challenge-serasa/api/controller/cryptoModule"
	"challenge-serasa/api/database"
	"challenge-serasa/api/mainframe"
)

type Controller struct {
	coll          database.Collection
	mainframeHost string
	passphrase    string
}

func NewController(coll database.Collection, host, passphrase string) *Controller {
	return &Controller{
		coll:          coll,
		mainframeHost: host,
		passphrase:    passphrase,
	}
}

func (c *Controller) UpdateNegativations() error {
	negativations, err := mainframe.GetNegativations(c.mainframeHost)
	if err != nil {
		return err
	}
	encryptedData, err := c.encryptNegativations(negativations)
	if err != nil {
		return err
	}
	err = c.coll.SaveDocuments(encryptedData)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetNegativationByCustomer(customerId string) ([]mainframe.Negativation, error) {
	negativations, err := c.coll.GetDocuments(customerId, "customerDocument")
	if err != nil {
		return nil, err
	}
	decryptedData, err := c.encryptNegativations(negativations)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}

func (c *Controller) encryptNegativations(data []mainframe.Negativation) ([]mainframe.Negativation, error) {
	var encryptedNegativations []mainframe.Negativation
	for _, negativation := range data {
		companyDocument, err := cryptoModule.Encrypt([]byte(negativation.CompanyDocument), c.passphrase)
		if err != nil {
			return nil, err
		}
		companyName, err := cryptoModule.Encrypt([]byte(negativation.CompanyName), c.passphrase)
		if err != nil {
			return nil, err
		}
		customerDocument, err := cryptoModule.Encrypt([]byte(negativation.CustomerDocument), c.passphrase)
		if err != nil {
			return nil, err
		}
		contract, err := cryptoModule.Encrypt([]byte(negativation.Contract), c.passphrase)
		if err != nil {
			return nil, err
		}
		encryptedNegativation := mainframe.GenerateNegativation(companyDocument, companyName, customerDocument, negativation.Value, contract, negativation.DebtDate, negativation.InclusionDate)
		encryptedNegativations = append(encryptedNegativations, *encryptedNegativation)
	}
	return encryptedNegativations, nil
}

func (c *Controller) decryptNegativations(data []mainframe.Negativation) ([]mainframe.Negativation, error) {
	var decryptedNegativations []mainframe.Negativation
	for _, negativation := range data {
		companyDocument, err := cryptoModule.Decrypt(negativation.CompanyDocument, c.passphrase)
		if err != nil {
			return nil, err
		}
		companyName, err := cryptoModule.Decrypt(negativation.CompanyName, c.passphrase)
		if err != nil {
			return nil, err
		}
		customerDocument, err := cryptoModule.Decrypt(negativation.CustomerDocument, c.passphrase)
		if err != nil {
			return nil, err
		}
		contract, err := cryptoModule.Decrypt(negativation.Contract, c.passphrase)
		if err != nil {
			return nil, err
		}
		decryptedNegativation := mainframe.GenerateNegativation(companyDocument, companyName, customerDocument, negativation.Value, contract, negativation.DebtDate, negativation.InclusionDate)
		decryptedNegativations = append(decryptedNegativations, *decryptedNegativation)
	}
	return decryptedNegativations, nil
}
