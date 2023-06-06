package user_auth_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	userauth "github.com/fit-smart/api/cmd/fit-smart/internal/user_auth"
	"github.com/fit-smart/api/internal/auth"
	"github.com/fit-smart/api/internal/auth/authfakes"
	"github.com/fit-smart/api/internal/auth/encryption/encryptionfakes"
	"github.com/fit-smart/api/internal/auth/google_oauth"
	"github.com/fit-smart/api/internal/auth/google_oauth/google_oauthfakes"
	"github.com/fit-smart/api/internal/auth/jwt/jwtfakes"
	"github.com/fit-smart/api/internal/logging/loggingfakes"
	"github.com/fit-smart/api/internal/mail/mailfakes"

	authgen "github.com/fit-smart/api/internal/gen/auth"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImVtYWlsIjoiaGhAZXhhbXBsZS5jb20iLCJyb2xlIjoic2ZpdF91c2VyIiwidXNlcklEIjoiMSJ9LCJleHAiOjE2NjI5MjU0NzQsImlzcyI6InNtYXJ0LWZpdC1hcGktZjc3YjVkYmEtNTgwYi00NDA0LWExODctMGY3OWU5MGEwMmJiIn0.1-pkQtFvmoeee5ZsEFtNmv6UIV5x32NonYmqyonOpy8"

var _ = Describe("User Auth Service", func() {
	var (
		fakeStore       *authfakes.FakeUserStore
		fakeGoogleOAuth *google_oauthfakes.FakeGoogleRegistry
		fakeEncryption  *encryptionfakes.FakeEncryption
		fakeJWT         *jwtfakes.FakeJWTClient
		fakeMail        *mailfakes.FakeService
		fakeLogger      *loggingfakes.FakeDatabaseLogger
		service         *userauth.Service
	)

	BeforeEach(func() {
		fakeStore = new(authfakes.FakeUserStore)
		fakeGoogleOAuth = new(google_oauthfakes.FakeGoogleRegistry)
		fakeEncryption = new(encryptionfakes.FakeEncryption)
		fakeJWT = new(jwtfakes.FakeJWTClient)
		fakeMail = new(mailfakes.FakeService)
		fakeLogger = new(loggingfakes.FakeDatabaseLogger)
		service = userauth.NewService(fakeStore, fakeGoogleOAuth, fakeEncryption, fakeJWT, fakeMail, fakeLogger)
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
						ID:       1,
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
						ID:       1,
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

	Describe("LoginGoogle", func() {
		var (
			ctx      context.Context
			err      error
			response *authgen.LoginResponse
		)

		JustBeforeEach(func() {
			ctx = context.Background()
			payload := &authgen.LoginGooglePayload{
				Token:        "token",
				RefreshToken: "refresh",
			}
			response, err = service.LoginGoogle(ctx, payload)
		})
		When("server gets user info from Google", func() {
			BeforeEach(func() {
				u := googleoauth.GoogleUserInfo{
					Email:         "test@example",
					FirstName:     "Ivan",
					LastName:      "Ivanov",
					PictureURL:    "https://example.com/example.jpeg",
					VerifiedEmail: true,
				}
				fakeGoogleOAuth.GetUserInfoReturns(u, nil)
			})

			var hasMoreThan25Chars = func(l int) bool { return l > 25 }
			var u *auth.User
			Context("and user with this email exists", func() {
				BeforeEach(func() {
					u = &auth.User{
						ID:    1,
						Email: "tony@example.com",
						Role:  auth.RoleUser,
					}
					fakeStore.GetUserByEmailReturns(u, nil)
					Expect(err).ToNot(HaveOccurred())
				})

				Context("and token generation succeeds", func() {
					BeforeEach(func() {
						fakeJWT.GenerateJWTReturns("ey$2ab123456iueigfrbrsggbrwhjhftfythtffh", nil)
					})
					It("should respond with success", func() {
						Expect(response.Email).To(Equal(u.Email))
						Expect(response.Token).To(ContainSubstring("ey"))
						Expect(len(response.Token)).To(Satisfy(hasMoreThan25Chars))
					})
				})

				Context("but token generation fails", func() {
					BeforeEach(func() {
						fakeJWT.GenerateJWTReturns("", errors.New("who's the boss here"))
					})
					It("should return an error", func() {
						Expect(err).To(HaveOccurred())
					})
				})

			})

			Context("and user with this email does NOT exist", func() {

				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, &auth.ErrUserNotFound{})
				})

				Context("and adding new user succeeds", func() {

					BeforeEach(func() {
						u := &auth.User{
							ID:       1,
							Email:    "tony@example.com",
							Password: "$2a$04$0kN0QJ7tWqGEADeYSHQfWO6d9BEwqcKyFu4advGBI33sYS/bhJSUC",
							Role:     auth.RoleUser,
						}
						fakeStore.RegisterReturns(u, nil)
					})
					Context("and token generation succeeds", func() {
						BeforeEach(func() {
							fakeJWT.GenerateJWTReturns("ey$2ab123456kghjkghjkhgjshjkshgjkhfhgdjkghfjkghfjh", nil)
						})
						It("should respond with success", func() {
							Expect(response.Email).To(Equal(u.Email))
							Expect(response.Token).To(ContainSubstring("ey"))
							Expect(len(response.Token)).To(Satisfy(hasMoreThan25Chars))
						})
					})

					Context("but token generation fails", func() {
						BeforeEach(func() {
							fakeJWT.GenerateJWTReturns("", errors.New("who's the boss here"))
						})
						It("should return an error", func() {
							Expect(err).To(HaveOccurred())
						})
					})

				})

				Context("but adding new user succeeds", func() {
					BeforeEach(func() {
						fakeStore.RegisterReturns(nil, errors.New("go away"))
					})
					It("should return an error", func() {
						Expect(err).To(HaveOccurred())
					})
				})
			})

			Context("but retrieving the user fails", func() {
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, errors.New("try again"))
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("error retrieving user with this email"))
				})
			})
		})

		When("server does NOT get user info from Google", func() {
			BeforeEach(func() {
				fakeGoogleOAuth.GetUserInfoReturns(googleoauth.GoogleUserInfo{}, errors.New("not funny"))
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("error getting user info"))
			})
		})

		When("user info from Google is NOT verified", func() {
			BeforeEach(func() {
				u := googleoauth.GoogleUserInfo{
					VerifiedEmail: false,
				}
				fakeGoogleOAuth.GetUserInfoReturns(u, nil)
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("Google account is not verified"))
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
						ID:        1,
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
					ID:        1,
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
