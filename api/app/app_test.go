package app

import (
	"bytes"
	"challenge-serasa/api/database"
	. "challenge-serasa/api/helper_tests/h_database"
	"challenge-serasa/api/helper_tests/h_mainframe"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	. "github.com/onsi/gomega"
)

const testCollection = "dummy-collection"

type token struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

func TestCollection(t *testing.T) {
	g := NewGomegaWithT(t)
	server := h_mainframe.MockMainframeServer(g)
	config := FakeAppConfig(server.URL)
	app, err := LoadApp(config)
	g.Expect(err).Should(BeNil())
	app.Run()
	defer app.Close()

	t.Run("validate integration", func(t *testing.T) {
		url := fmt.Sprintf("http://localhost:%d/v1/integration", config.Port)
		res, err := http.Post(url, "application/json", nil)
		g.Expect(err).Should(BeNil())
		g.Expect(res.StatusCode).Should(BeEquivalentTo(http.StatusOK))
	})

	t.Run("validate login", func(t *testing.T) {
		urlI := fmt.Sprintf("http://localhost:%d/v1/integration", config.Port)
		res, err := http.Post(urlI, "application/json", nil)
		g.Expect(err).Should(BeNil())
		g.Expect(res.StatusCode).Should(BeEquivalentTo(http.StatusOK))
		var jsonStr = []byte(`{"customerDocument":"62824334010"}`)
		urlL := fmt.Sprintf("http://localhost:%d/v1/login", config.Port)
		req, err := http.NewRequest("POST", urlL, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		clientResp, err := client.Do(req)
		g.Expect(err).Should(BeNil())
		g.Expect(clientResp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
		body, _ := ioutil.ReadAll(clientResp.Body)
		g.Expect(string(body)).ShouldNot(BeEquivalentTo(""))
	})

	t.Run("validate get negativations", func(t *testing.T) {
		urlI := fmt.Sprintf("http://localhost:%d/v1/integration", config.Port)
		res, err := http.Post(urlI, "application/json", nil)
		g.Expect(err).Should(BeNil())
		g.Expect(res.StatusCode).Should(BeEquivalentTo(http.StatusOK))
		var jsonStr = []byte(`{"customerDocument":"62824334010"}`)
		urlL := fmt.Sprintf("http://localhost:%d/v1/login", config.Port)
		req, err := http.NewRequest("POST", urlL, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		clientResp, err := client.Do(req)
		g.Expect(err).Should(BeNil())
		g.Expect(clientResp.StatusCode).Should(BeEquivalentTo(http.StatusOK))
		body, _ := ioutil.ReadAll(clientResp.Body)
		g.Expect(string(body)).ShouldNot(BeEquivalentTo(""))
		var token token
		err = json.Unmarshal(body, &token)
		g.Expect(err).Should(BeNil())
		urlN := fmt.Sprintf("http://localhost:%d/v1/negativations/62824334010", config.Port)
		req, err = http.NewRequest("GET", urlN, nil)
		req.Header.Set("Token", token.Token)

		client = &http.Client{}
		clientResp, err = client.Do(req)
		g.Expect(err).Should(BeNil())
		g.Expect(clientResp.StatusCode).Should(BeEquivalentTo(http.StatusOK))

	})

}

func FakeAppConfig(url string) Config {
	return Config{
		Passphrase:   "secretpassphrase",
		Key:          "secretkey",
		MainframeUrl: url,
		Port:         8082,
		Database: Database{
			Config: database.Config{
				Host:     DBHostTest,
				Port:     DBPortTest,
				User:     DBUserTest,
				Password: DBPassTest,
				Database: DBNameTest,
			},
			NegativationCollection: "dummy-collection",
		},
	}
}
