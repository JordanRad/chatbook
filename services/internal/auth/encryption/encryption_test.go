package encryption_test

import (
	"github.com/JordanRad/chatbook/services/internal/auth/encryption"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encryption", func() {

	var encrypter encryption.Encryption

	const (
		VALID_HASH = "$2a$10$G0D/Q2I8KLTlhpy58fCsGeVGlSOUhnK/TOlMJnLjT367vEN62mpzS"
		TEST_PASS  = "123456"
	)

	BeforeEach(func() {
		encrypter = new(encryption.Encrypter)
	})

	Describe("Encrypt Password", func() {
		var (
			result string
			err    error
		)
		When("ecryption succeeds", func() {
			JustBeforeEach(func() {
				result, err = encrypter.EncryptPassword(TEST_PASS)
			})
			It("should return hashed string", func() {
				Expect(result).To(ContainSubstring("$2a"))
			})
			It("should NOT return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Describe("Check Password", func() {
		var ok bool

		When("the password is correct", func() {
			JustBeforeEach(func() {
				ok = encrypter.CheckPassword(VALID_HASH, TEST_PASS)
			})
			It("should return true", func() {
				Expect(ok).To(BeTrue())
			})
		})

		When("the password is NOT correct", func() {
			JustBeforeEach(func() {
				ok = encrypter.CheckPassword(VALID_HASH, "alabala0")
			})
			It("should return false", func() {
				Expect(ok).To(BeFalse())
			})
		})

	})
})
