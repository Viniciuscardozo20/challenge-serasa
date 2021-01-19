package cryptoModule

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestEncypt(t *testing.T) {
	g := NewGomegaWithT(t)
	dataEncrypted, err := Encrypt([]byte("Test Encrypt data"), "password")
	g.Expect(err).Should(BeNil())
	g.Expect(string(dataEncrypted)).ShouldNot(BeNil())
}

func TestDecrypt(t *testing.T) {
	g := NewGomegaWithT(t)
	dataEncrypted, err := Encrypt([]byte("Test Encrypt data"), "password")
	g.Expect(err).Should(BeNil())
	g.Expect(string(dataEncrypted)).ShouldNot(BeNil())
	dataDecrypted, err := Decrypt(dataEncrypted, "password")
	g.Expect(err).Should(BeNil())
	g.Expect(string(dataDecrypted)).ShouldNot(BeNil())
	g.Expect(dataDecrypted).Should(BeEquivalentTo("Test Encrypt data"))
}
