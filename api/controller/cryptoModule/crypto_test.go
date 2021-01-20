package cryptoModule

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestEncypt(t *testing.T) {
	g := NewGomegaWithT(t)
	dataEncrypted, err := Encrypt([]byte("Test Encrypt data"), "secretpassphrase")
	g.Expect(err).Should(BeNil())
	g.Expect(string(dataEncrypted)).ShouldNot(BeNil())
	g.Expect(dataEncrypted).Should(BeEquivalentTo("df666f21c96de26f9eec7171e5e2a4d9b44e99bbd49068d44d4dd63f20b9eb65a6"))
}

func TestDecrypt(t *testing.T) {
	g := NewGomegaWithT(t)
	dataEncrypted, err := Encrypt([]byte("Test Encrypt data"), "secretpassphrase")
	g.Expect(err).Should(BeNil())
	g.Expect(string(dataEncrypted)).ShouldNot(BeNil())
	dataDecrypted, err := Decrypt(dataEncrypted, "secretpassphrase")
	g.Expect(err).Should(BeNil())
	g.Expect(string(dataDecrypted)).ShouldNot(BeNil())
	g.Expect(dataDecrypted).Should(BeEquivalentTo("Test Encrypt data"))
}
