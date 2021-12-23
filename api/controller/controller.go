package controller

import (
	"challenge-serasa/api/controller/auth"
	"challenge-serasa/api/controller/cryptoModule"
	"challenge-serasa/api/database"
	"challenge-serasa/api/mainframe"
)

type Controller struct {
	coll          database.Collection
	mainframeHost string
	passphrase    string
	secretkey     string
}

func NewController(coll database.Collection, host, passphrase, secretkey string) *Controller {
	return &Controller{
		coll:          coll,
		mainframeHost: host,
		passphrase:    passphrase,
		secretkey:     secretkey,
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

func (c *Controller) GetNegativationByCustomer(customerDocument string) ([]mainframe.Negativation, error) {
	encryptedCustomer, err := cryptoModule.Encrypt([]byte(customerDocument), c.passphrase)
	if err != nil {
		return nil, err
	}
	negativations, err := c.coll.GetDocuments(encryptedCustomer, "customerDocument")
	if err != nil {
		return nil, err
	}
	decryptedData, err := c.decryptNegativations(negativations)
	if err != nil {
		return nil, err
	}
	return decryptedData, nil
}

func (c *Controller) Login(customerDocument string) (string, error) {
	encryptedCustomer, err := cryptoModule.Encrypt([]byte(customerDocument), c.passphrase)
	if err != nil {
		return "", err
	}
	n, err := c.coll.GetDocuments(encryptedCustomer, "customerDocument")
	if err != nil || len(n) == 0 {
		return "", err
	}
	token, err := auth.CreateToken(customerDocument, c.secretkey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *Controller) TokenValid(dataToken string) error {
	token, err := auth.VerifyToken(dataToken)
	if err != nil {
		return err
	}
	if !token.Valid {
		return err
	}
	return nil
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
