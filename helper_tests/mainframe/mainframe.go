package mainframe

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/gomega"
)

func MockMainframeServer(g *GomegaWithT) *httptest.Server {
	// Mock http server
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(FakeDataNavigations()))
			} else {
				w.WriteHeader(http.StatusNotFound)
			}
		}),
	)
	return ts
}

func FakeDataNavigations() []byte {
	return []byte(`  
		[
			{
				"companyDocument": "59291534000167",
				"companyName": "ABC S.A.",
				"customerDocument": "51537476467",
				"value": 1235.23,
				"contract": "bc063153-fb9e-4334-9a6c-0d069a42065b",
				"debtDate": "2015-11-13T20:32:51-03:00",
				"inclusionDate": "2020-11-13T20:32:51-03:00"
			},
			{
				"companyDocument": "77723018000146",
				"companyName": "123 S.A.",
				"customerDocument": "51537476467",
				"value": 400.00,
				"contract": "5f206825-3cfe-412f-8302-cc1b24a179b0",
				"debtDate": "2015-10-12T20:32:51-03:00",
				"inclusionDate": "2020-10-12T20:32:51-03:00"
			}
		]
					
`)

}
