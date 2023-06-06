package userauth_test

import (
	"context"
	"errors"
	"log"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/JordanRad/chatbook/services/cmd/user-management-service/userauth"
	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/auth/authfakes"
	"github.com/JordanRad/chatbook/services/internal/auth/encryption/encryptionfakes"
	"github.com/JordanRad/chatbook/services/internal/auth/jwt/jwtfakes"
	authgen "github.com/JordanRad/chatbook/services/internal/gen/auth"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImVtYWlsIjoiaGhAZXhhbXBsZS5jb20iLCJyb2xlIjoic2ZpdF91c2VyIiwidXNlcklEIjoiMSJ9LCJleHAiOjE2NjI5MjU0NzQsImlzcyI6InNtYXJ0LWZpdC1hcGktZjc3YjVkYmEtNTgwYi00NDA0LWExODctMGY3OWU5MGEwMmJiIn0.1-pkQtFvmoeee5ZsEFtNmv6UIV5x32NonYmqyonOpy8"

var _ = Describe("User Auth Service", func() {
	var (
		fakeStore      *authfakes.FakeUserStore
		fakeEncryption *encryptionfakes.FakeEncryption
		fakeJWT        *jwtfakes.FakeJWTClient

		service *userauth.Service
	)

	BeforeEach(func() {
		fakeStore = new(authfakes.FakeUserStore)
		fakeEncryption = new(encryptionfakes.FakeEncryption)
		fakeJWT = new(jwtfakes.FakeJWTClient)
		// Initialize loger
		logger := log.New(os.Stdout, "", log.LstdFlags)
		service = userauth.NewService(fakeStore, fakeEncryption, fakeJWT, logger)
	})

	Describe("Login", func() {
		var (
			ctx      context.Context
			err      error
			response *authgen.LoginResponse
		)

		Context("password is correct", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				payload := &authgen.LoginPayload{
					Email:    "tony@example.com",
					Password: "tony2chase",
				}
				response, err = service.Login(ctx, payload)
			})
			Context("user successfully logged in and service responds with a valid JWT", func() {
				var validUser *auth.User
				BeforeEach(func() {
					validUser = &auth.User{
						ID:       "1",
						Email:    "tony@example.com",
						Password: "$2a$04$0kN0QJ7tWqGEADeYSHQfWO6d9BEwqcKyFu4advGBI33sYS/bhJSUC",
					}
					fakeStore.GetUserByEmailReturns(validUser, nil)
				})

				Context("and password check succeeds", func() {
					BeforeEach(func() {
						fakeEncryption.CheckPasswordReturns(true)
					})

					Context("and token generating succeeds", func() {
						BeforeEach(func() {
							fakeJWT.GenerateJWTReturns("eysafetokenalabala123456789102", nil)
						})
						var hasMoreThan25Chars = func(l int) bool { return l > 25 }

						It("should respond with success", func() {
							Expect(response.Email).To(Equal(validUser.Email))
							Expect(response.Token).To(ContainSubstring("ey"))
							Expect(len(response.Token)).To(Satisfy(hasMoreThan25Chars))
						})
						It("should NOT return errors", func() {
							Expect(err).To(Not(HaveOccurred()))
						})
					})

				})

				Context("but password check fails", func() {
					BeforeEach(func() {
						fakeEncryption.CheckPasswordReturns(false)
					})

					Context("and token generating fails", func() {
						BeforeEach(func() {
							fakeJWT.GenerateJWTReturns("", errors.New("bye bro"))
						})
						It("should return an error", func() {
							Expect(err).To(HaveOccurred())
						})
					})
				})

			})

			Context("login is NOT successful due to wrong email, service returns invalid credentials error", func() {
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, errors.New("invalid credentials"))
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Context("password is NOT correct", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				payload := &authgen.LoginPayload{
					Email:    "tony@example.com",
					Password: "tonychase123",
				}
				response, err = service.Login(ctx, payload)
			})

			When("user login is NOT successful and response returns nil", func() {
				var invalidUser *auth.User
				BeforeEach(func() {
					invalidUser = &auth.User{
						ID:       "1",
						Email:    "tony@example.com",
						Password: "$2a$04$0kN0QJ7tWqGEADeYSHQfWO6d9BEwqcKyFu4advGBI33sYS/bhJSUC",
					}
					fakeStore.GetUserByEmailReturns(invalidUser, nil)
				})

				It("should return an error", func() {
					Expect(response).To(BeNil())
				})
			})
		})
	})

	Describe("RefreshToken", func() {
		var (
			ctx context.Context
			r   *authgen.LoginResponse
			err error
		)

		When("payload is valid", func() {
			ctx = context.Background()
			JustBeforeEach(func() {
				p := &authgen.RefreshTokenPayload{
					Email:        "ivo@example.com",
					RefreshToken: token,
				}
				r, err = service.RefreshToken(ctx, p)
			})

			Context("and such user exists", func() {
				var u *auth.User
				BeforeEach(func() {
					u = &auth.User{
						ID:        "1",
						Email:     "ivo@example",
						FirstName: "Ivo",
						LastName:  "Andonov",
					}
					fakeStore.GetUserByEmailReturns(u, nil)

				})

				Context("and token validation succeeds", func() {
					BeforeEach(func() {
						fakeJWT.ValidateJWTReturns(true, nil)
					})

					Context("and token generating succeeds", func() {
						BeforeEach(func() {
							fakeJWT.GenerateJWTReturns("eysafetokenalabala123456789102", nil)
						})
						var hasMoreThan25Chars = func(l int) bool { return l > 25 }

						It("should respond with success", func() {
							Expect(r.Email).To(Equal(u.Email))
							Expect(r.Token).To(ContainSubstring("ey"))
							Expect(len(r.Token)).To(Satisfy(hasMoreThan25Chars))
						})
						It("should NOT return errors", func() {
							Expect(err).To(Not(HaveOccurred()))
						})
					})
					Context("but token generating fails", func() {
						BeforeEach(func() {
							fakeJWT.GenerateJWTReturns("", errors.New("try again next year"))
						})

						It("should return an error", func() {
							Expect(err).To(HaveOccurred())
						})

					})
				})

				Context("but token validation fails", func() {
					BeforeEach(func() {
						fakeJWT.ValidateJWTReturns(false, errors.New("try again later"))
					})

					It("should return an error", func() {
						Expect(err).To(HaveOccurred())
					})
				})

			})

			Context("and such user does NOT exist", func() {
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, errors.New("**** off"))
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})

		When("refresh token is NOT valid", func() {
			ctx = context.Background()
			JustBeforeEach(func() {
				p := &authgen.RefreshTokenPayload{
					Email:        "ivo@example.com",
					RefreshToken: "defnotvalid",
				}
				r, err = service.RefreshToken(ctx, p)
			})
			var u *auth.User
			BeforeEach(func() {
				u = &auth.User{
					ID:        "1",
					Email:     "ivo@example",
					FirstName: "Ivo",
					LastName:  "Andonov",
				}
				fakeStore.GetUserByEmailReturns(u, errors.New("e"))
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})

		When("email is NOT valid", func() {
			ctx = context.Background()
			JustBeforeEach(func() {
				p := &authgen.RefreshTokenPayload{
					Email:        "ivo@example.com",
					RefreshToken: token,
				}
				r, err = service.RefreshToken(ctx, p)
			})
			BeforeEach(func() {
				fakeStore.GetUserByEmailReturns(nil, errors.New("bug off"))
			})
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
