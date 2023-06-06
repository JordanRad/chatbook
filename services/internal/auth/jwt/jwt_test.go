package jwt_test

import (
	"github.com/JordanRad/chatbook/services/internal/auth/jwt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Jwt", func() {

	var service jwt.JWTClient

	const (
		USER_ID = "1aaaaaaaaaa2b"
		EMAIL   = "test@example.com"
	)
	const (
		VALID_HASH = "$2a$10$G0D/Q2I8KLTlhpy58fCsGeVGlSOUhnK/TOlMJnLjT367vEN62mpzS"
		TEST_PASS  = "123456"
	)

	BeforeEach(func() {
		service = new(jwt.JWTService)
	})

	Describe("Generate JWT", func() {
		var (
			result string
			err    error
		)
		When("generating succeeds", func() {
			JustBeforeEach(func() {
				result, err = service.GenerateJWT(USER_ID, EMAIL, false)
			})
			It("should return hashed string", func() {
				Expect(result).To(ContainSubstring("ey"))
			})
			It("should NOT return an error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

})
