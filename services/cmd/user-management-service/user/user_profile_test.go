package user_test

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/JordanRad/chatbook/services/cmd/user-management-service/user"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/JordanRad/chatbook/services/internal/auth"
	"github.com/JordanRad/chatbook/services/internal/auth/authfakes"
	"github.com/JordanRad/chatbook/services/internal/auth/encryption/encryptionfakes"
	usergen "github.com/JordanRad/chatbook/services/internal/gen/user"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImVtYWlsIjoiaGhAZXhhbXBsZS5jb20iLCJyb2xlIjoic2ZpdF91c2VyIiwidXNlcklEIjoiMSJ9LCJleHAiOjE2NjI5MjU0NzQsImlzcyI6InNtYXJ0LWZpdC1hcGktZjc3YjVkYmEtNTgwYi00NDA0LWExODctMGY3OWU5MGEwMmJiIn0.1-pkQtFvmoeee5ZsEFtNmv6UIV5x32NonYmqyonOpy8"

var _ = Describe("User Profile Service", func() {
	var (
		fakeStore      *authfakes.FakeUserStore
		fakeEncryption *encryptionfakes.FakeEncryption

		service *user.Service
	)

	BeforeEach(func() {
		fakeStore = new(authfakes.FakeUserStore)
		fakeEncryption = new(encryptionfakes.FakeEncryption)

		// Initialize loger
		logger := log.New(os.Stdout, "", log.LstdFlags)

		service = user.NewService(fakeStore, fakeEncryption, logger)
	})

	Describe("Register", func() {
		var (
			ctx      context.Context
			err      error
			response *usergen.RegisterResponse
		)

		Context("matching passwords are provided", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				payload := &usergen.RegisterPayload{
					FirstName:         "Tony",
					LastName:          "Chase",
					Email:             "t@example.com",
					Password:          "tony2chase",
					ConfirmedPassword: "tony2chase",
				}
				response, err = service.Register(ctx, payload)
			})
			Context("and password encryption succeeds", func() {
				BeforeEach(func() {
					fakeEncryption.EncryptPasswordReturns("$2ab123456", nil)
				})

			})

			Context("but password encryption fails", func() {
				BeforeEach(func() {
					fakeEncryption.EncryptPasswordReturns("", errors.New("who's the boss here"))
				})

			})

			When("user is registered successfully", func() {
				BeforeEach(func() {
					u := &auth.User{}
					fakeStore.RegisterReturns(u, nil)
				})
				It("should return succesful register message", func() {
					Expect(response.Message).To(Equal("User has been registered successfully"))
				})
				It("should NOT return errors", func() {
					Expect(err).To(Not(HaveOccurred()))
				})
			})

			When("user is NOT successfully created", func() {
				BeforeEach(func() {
					fakeStore.RegisterReturns(nil, errors.New("that's too sad"))
				})
				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Context("NON-matching passwords are provided", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				payload := &usergen.RegisterPayload{
					FirstName:         "Tony",
					LastName:          "Chase",
					Email:             "tony@example.com",
					Password:          "tony2chase",
					ConfirmedPassword: "tonychase",
				}
				response, err = service.Register(ctx, payload)
			})

			When("user is NOT successfully created", func() {
				It("should return error: passwords do not match!", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("GetProfile", func() {
		var (
			ctx      context.Context
			err      error
			response *usergen.UserProfileResponse
		)

		Context("auth details are correct", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, auth.ContextKeyUser, &auth.User{
					ID:    "123xyz",
					Email: "t@example.com",
				})
				response, err = service.GetProfile(ctx)
			})

			When("user retrieval succeeds", func() {
				var validUser *auth.User
				BeforeEach(func() {
					validUser = &auth.User{
						ID:    "123xyz",
						Email: "t@example.com",
					}
					fakeStore.GetUserByEmailReturns(validUser, nil)
				})

				It("should respond with success", func() {
					Expect(response).ToNot(BeNil())
					Expect(response.ID).To(Equal(validUser.ID))
					Expect(response.Email).To(Equal(validUser.Email))
				})
				It("should NOT return an error", func() {
					Expect(err).ToNot(HaveOccurred())
				})
			})

			When("user retrieval fails", func() {
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, errors.New("I found a hacker!"))
				})
				It("should  return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})

		})

		Context("context does NOT contain any user", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, auth.ContextKeyUser, "")
				response, err = service.GetProfile(ctx)
			})
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update names", func() {
		var (
			ctx      context.Context
			response *usergen.OperationStatusResponse
			err      error
		)
		Context("there is a valid user in context", func() {
			var p *usergen.UpdateProfileNamesPayload
			JustBeforeEach(func() {
				ctx = context.Background()

				ctx = context.WithValue(ctx, auth.ContextKeyUser, &auth.User{
					ID:    "12",
					Email: "t@example.com",
				})

				response, err = service.UpdateProfileNames(ctx, p)
			})
			When("the payload is correct and", func() {
				When("the profile names update succeeds", func() {
					BeforeEach(func() {
						p = &usergen.UpdateProfileNamesPayload{
							FirstName: "Michael",
							LastName:  "Jackson",
						}

						fakeStore.UpdateProfileNamesReturns(nil)
						Expect(err).ToNot(HaveOccurred())
					})
					It("should return a message for success", func() {
						Expect(response).ToNot(BeNil())
						Expect(response.Message).To(ContainSubstring("names have been updated successfully"))
					})
				})

				When("the profile names update fails", func() {
					BeforeEach(func() {
						p = &usergen.UpdateProfileNamesPayload{
							FirstName: "Michael",
							LastName:  "Jackson",
						}

						fakeStore.UpdateProfileNamesReturns(errors.New("try l8r bruh"))
					})
					It("should return a message for success", func() {
						Expect(err).To(HaveOccurred())
					})
				})

			})

			When("the payload is NOT correct (full)", func() {
				BeforeEach(func() {
					p = &usergen.UpdateProfileNamesPayload{
						FirstName: "Michael",
					}

					fakeStore.UpdateProfileNamesReturns(nil)
				})
				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("context does NOT contain a user", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, auth.ContextKeyUser, "")
				payload := &usergen.UpdateProfileNamesPayload{
					FirstName: "",
					LastName:  "",
				}
				response, err = service.UpdateProfileNames(ctx, payload)
			})
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
