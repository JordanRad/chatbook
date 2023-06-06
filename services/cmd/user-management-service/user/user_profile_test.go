package user_profile_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/fit-smart/api/internal/auth"
	"github.com/fit-smart/api/internal/auth/authfakes"
	"github.com/fit-smart/api/internal/auth/encryption/encryptionfakes"
	"github.com/fit-smart/api/internal/cloud_storage/cloud_storagefakes"
	usergen "github.com/fit-smart/api/internal/gen/user"
	"github.com/fit-smart/api/internal/logging/loggingfakes"
	"github.com/fit-smart/api/internal/mail/mailfakes"
)

const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7ImVtYWlsIjoiaGhAZXhhbXBsZS5jb20iLCJyb2xlIjoic2ZpdF91c2VyIiwidXNlcklEIjoiMSJ9LCJleHAiOjE2NjI5MjU0NzQsImlzcyI6InNtYXJ0LWZpdC1hcGktZjc3YjVkYmEtNTgwYi00NDA0LWExODctMGY3OWU5MGEwMmJiIn0.1-pkQtFvmoeee5ZsEFtNmv6UIV5x32NonYmqyonOpy8"

var _ = Describe("User Profile Service", func() {
	var (
		fakeStore        *authfakes.FakeUserStore
		fakeCloudStorage *cloud_storagefakes.FakeStorage
		fakeEncryption   *encryptionfakes.FakeEncryption
		fakeMail         *mailfakes.FakeService
		fakeLogger       *loggingfakes.FakeDatabaseLogger
		service          *user_profile.Service
	)

	BeforeEach(func() {
		fakeStore = new(authfakes.FakeUserStore)
		fakeCloudStorage = new(cloud_storagefakes.FakeStorage)
		fakeEncryption = new(encryptionfakes.FakeEncryption)
		fakeMail = new(mailfakes.FakeService)
		fakeLogger = new(loggingfakes.FakeDatabaseLogger)
		service = user_profile.NewService(fakeStore, fakeCloudStorage, fakeEncryption, fakeMail, fakeLogger)
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
					Settings:          &usergen.BodyMetricsSettings{Length: "cm", Distance: "km", Weight: "kg"},
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
					Settings:          &usergen.BodyMetricsSettings{Length: "cm", Distance: "km", Weight: "kg"},
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

	Describe("GetUserProfile", func() {
		var (
			ctx      context.Context
			err      error
			response *usergen.UserProfileResponse
		)

		Context("auth details are correct", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, auth.ContextKeyUser, &auth.User{
					ID:    1,
					Email: "t@example.com",
					Role:  "sfit_user",
				})
				response, err = service.GetUserProfile(ctx)
			})

			When("user retrieval succeeds", func() {
				var validUser *auth.User
				BeforeEach(func() {
					validUser = &auth.User{
						ID:    1,
						Email: "t@example.com",
						Role:  "sfit_user",
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
				response, err = service.GetUserProfile(ctx)
			})
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("UpdateSettings", func() {
		var (
			ctx      context.Context
			response *usergen.OperationStatusResponse
			err      error
		)
		Context("there is a valid user in context", func() {
			var p *usergen.UpdateSettingsPayload
			JustBeforeEach(func() {
				ctx = context.Background()

				ctx = context.WithValue(ctx, auth.ContextKeyUser, &auth.User{
					ID:    1,
					Email: "t@example.com",
					Role:  "sfit_user",
				})

				response, err = service.UpdateSettings(ctx, p)
			})
			When("and the payload is correct", func() {
				When("and the settings update succeeds", func() {
					BeforeEach(func() {
						p = &usergen.UpdateSettingsPayload{
							Settings: &usergen.BodyMetricsSettings{
								Length:   "test-cm",
								Weight:   "test-kg",
								Distance: "test-distance",
							}}

						fakeStore.UpdateUserSettingsReturns(nil)
						Expect(err).ToNot(HaveOccurred())
					})
					It("should return a message for success", func() {
						Expect(response).ToNot(BeNil())
						Expect(response.Message).To(ContainSubstring("settings have been updated successfully"))
					})
				})

				When("but the settings update fails", func() {
					BeforeEach(func() {
						p = &usergen.UpdateSettingsPayload{
							Settings: &usergen.BodyMetricsSettings{
								Length:   "test-cm",
								Weight:   "test-kg",
								Distance: "test-distance",
							}}

						fakeStore.UpdateUserSettingsReturns(errors.New("just try a bit later, dude"))

					})
					It("should return a message for failure", func() {
						Expect(err).To(HaveOccurred())
					})
				})

			})

			When("and the payload is NOT correct", func() {
				BeforeEach(func() {
					p = &usergen.UpdateSettingsPayload{
						Settings: &usergen.BodyMetricsSettings{
							Distance: "test-distance",
						}}

					fakeStore.UpdateUserSettingsReturns(nil)

				})
				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})

		})
		Context("context does NOT contain any user", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, auth.ContextKeyUser, "")
				payload := &usergen.UpdateSettingsPayload{
					Settings: &usergen.BodyMetricsSettings{},
				}
				response, err = service.UpdateSettings(ctx, payload)
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
					ID:    1,
					Email: "t@example.com",
					Role:  "sfit_user",
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

	Describe("Update password", func() {
		var (
			ctx      context.Context
			response *usergen.OperationStatusResponse
			err      error
		)
		Context("there is a valid user in context", func() {
			var p *usergen.UpdatePasswordPayload
			JustBeforeEach(func() {
				ctx = context.Background()

				ctx = context.WithValue(ctx, auth.ContextKeyUser, &auth.User{
					ID:       1,
					Email:    "t@example.com",
					Role:     "sfit_user",
					Password: "$2a$04$xjB2iYishesrSJhEOoIj5.1axnC963oWIFNWy293uMAoXEyIDKZMS",
				})

				response, err = service.UpdatePassword(ctx, p)
			})
			Context("and password check succeeds", func() {
				BeforeEach(func() {
					fakeEncryption.CheckPasswordReturns(true)
				})
				Context("and password encryption succeeds", func() {
					BeforeEach(func() {
						fakeEncryption.EncryptPasswordReturns("$2ab123456", nil)
					})

					When("the payload is correct and", func() {
						When("the password update succeeds", func() {
							BeforeEach(func() {
								p = &usergen.UpdatePasswordPayload{
									OldPassword: "tonychase",
									NewPassword: "tony2chase",
								}

								fakeStore.UpdatePasswordReturns(nil)
								Expect(err).ToNot(HaveOccurred())
							})
							It("should return a message for success", func() {
								Expect(response).ToNot(BeNil())
								Expect(response.Message).To(ContainSubstring("password has been updated successfully"))
							})
						})

						When("the password update fails", func() {
							BeforeEach(func() {
								p = &usergen.UpdatePasswordPayload{
									OldPassword: "tonychase",
									NewPassword: "tony2chase",
								}

								fakeStore.UpdatePasswordReturns(errors.New("try l8r bruh"))
							})
							It("should return a message for failure", func() {
								Expect(err).To(HaveOccurred())
							})
						})

					})

					When("the payload is NOT correct", func() {
						When("the old password is NOT correct", func() {
							BeforeEach(func() {
								p = &usergen.UpdatePasswordPayload{
									OldPassword: "tonychasee",
									NewPassword: "tony2chase",
								}

								fakeStore.UpdatePasswordReturns(errors.New("financial mistake"))
							})
							It("should return an error", func() {
								Expect(err).To(HaveOccurred())
							})
						})

						When("the the old password is NOT present in payload", func() {
							BeforeEach(func() {
								p = &usergen.UpdatePasswordPayload{
									NewPassword: "tony2chase",
								}

								fakeStore.UpdatePasswordReturns(errors.New("don't blame me"))
							})
							It("should return an error", func() {
								Expect(err).To(HaveOccurred())
							})
						})
					})

				})
				Context("but password encryption fails", func() {
					BeforeEach(func() {
						fakeEncryption.EncryptPasswordReturns("", errors.New("who's the boss here"))
					})
					It("should return an error", func() {
						Expect(err).To(HaveOccurred())
					})
				})
			})

			Context("but password check fails", func() {
				BeforeEach(func() {
					fakeEncryption.CheckPasswordReturns(false)
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
				payload := &usergen.UpdatePasswordPayload{
					OldPassword: "",
					NewPassword: "",
				}
				response, err = service.UpdatePassword(ctx, payload)
			})
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete user", func() {
		var (
			ctx      context.Context
			response *usergen.OperationStatusResponse
			err      error
		)
		Context("there is a valid user in context", func() {
			JustBeforeEach(func() {
				ctx = context.Background()

				ctx = context.WithValue(ctx, auth.ContextKeyUser, &auth.User{
					ID:       1,
					Email:    "t@example.com",
					Role:     "sfit_user",
					Password: "$2a$04$xjB2iYishesrSJhEOoIj5.1axnC963oWIFNWy293uMAoXEyIDKZMS",
				})

				response, err = service.DeleteUser(ctx)
			})

			When("the user is successfully deleted", func() {
				BeforeEach(func() {
					fakeStore.DeleteUserReturns(nil)
					Expect(err).ToNot(HaveOccurred())
				})
				It("should return a message for success", func() {
					Expect(response).ToNot(BeNil())
					Expect(response.Message).To(ContainSubstring("User has been deleted successfully"))
				})
			})
			When("the user is NOT successfully deleted", func() {
				BeforeEach(func() {
					fakeStore.DeleteUserReturns(errors.New("nope"))
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

				response, err = service.DeleteUser(ctx)
			})
			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Reset password", func() {
		var (
			ctx context.Context
			r   *usergen.OperationStatusResponse
			err error
		)
		Context("the payload is valid and", func() {
			ctx = context.Background()
			JustBeforeEach(func() {
				p := &usergen.ResetPasswordPayload{
					Email:    "t@example.com",
					Password: "tony2chase",
				}
				r, err = service.ResetPassword(ctx, p)
			})
			Context("such user exists", func() {
				var u *auth.User
				BeforeEach(func() {
					u = &auth.User{
						ID:       1,
						Email:    "t@example",
						Role:     "sfit_user",
						Password: "$2a$04$xjB2iYishesrSJhEOoIj5.1axnC963oWIFNWy293uMAoXEyIDKZMS",
					}
					fakeStore.GetUserByEmailReturns(u, nil)
				})

				When("the password encryption succeeds", func() {
					BeforeEach(func() {
						fakeEncryption.EncryptPasswordReturns("$2a$04$ver/JJzJExg8T/5OxhpJneQycQx43OWgccrXoofxnKEEktzOQxwra", nil)
					})

					When("and the password reset succeeds", func() {
						BeforeEach(func() {
							fakeStore.UpdatePasswordReturns(nil)
							Expect(err).ToNot(HaveOccurred())
						})
						It("should return a message for success", func() {
							Expect(r).ToNot(BeNil())
							Expect(r.Message).To(ContainSubstring("password has been updated successfully"))
						})
					})

					When("but the password reset fails", func() {
						BeforeEach(func() {
							fakeStore.UpdatePasswordReturns(errors.New("try l8r bruh"))
						})
						It("should have an empty response and throw an error", func() {
							Expect(r).To(BeNil())
							Expect(err).To(HaveOccurred())
						})
					})
				})
				When("the password encryption fails", func() {
					BeforeEach(func() {
						fakeEncryption.EncryptPasswordReturns("", errors.New("this aint gon work"))
					})
					It("should have an empty response and throw an error", func() {
						Expect(r).To(BeNil())
						Expect(err).To(HaveOccurred())
					})
				})
			})
			Context("user does NOT exist", func() {
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, errors.New("nope"))
				})
				It("should have an empty response and throw an error", func() {
					Expect(r).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})

		Context("the payload is NOT valid and", func() {
			ctx = context.Background()
			JustBeforeEach(func() {
				p := &usergen.ResetPasswordPayload{
					Email: "t@example.com",
				}
				r, err = service.ResetPassword(ctx, p)
			})
			When("the password encryption fails", func() {
				BeforeEach(func() {
					fakeEncryption.EncryptPasswordReturns("", errors.New("this aint gon work"))
				})
				It("should have an empty response and throw an error", func() {
					Expect(r).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("Recover password", func() {
		var (
			ctx context.Context
			r   *usergen.OperationStatusResponse
			err error
		)
		Context("the payload is valid and", func() {
			ctx = context.Background()
			JustBeforeEach(func() {
				p := &usergen.RecoverPasswordPayload{
					Email: "t@example.com",
				}
				r, err = service.RecoverPassword(ctx, p)
			})
			Context("such user exists", func() {
				var u *auth.User
				BeforeEach(func() {
					u = &auth.User{
						ID:       1,
						Email:    "t@example",
						Role:     "sfit_user",
						Password: "$2a$04$xjB2iYishesrSJhEOoIj5.1axnC963oWIFNWy293uMAoXEyIDKZMS",
					}
					fakeStore.GetUserByEmailReturns(u, nil)
				})

				When("and the email is generated and sent successfully", func() {
					BeforeEach(func() {
						fakeMail.SendEmailReturns(nil)
					})
					It("should return a message for success", func() {
						Expect(r).ToNot(BeNil())
						Expect(r.Message).To(ContainSubstring("recovery email successfully sent"))
					})
					It("should NOT return an error", func() {
						Expect(err).ToNot(HaveOccurred())
					})
				})
				When("but the email generation fails", func() {
					BeforeEach(func() {
						fakeMail.SendEmailReturns(errors.New("this aint gon work pal"))
					})
					It("should return an error", func() {
						Expect(err).To(HaveOccurred())
					})
				})
			})
			Context("user does NOT exist", func() {
				BeforeEach(func() {
					fakeStore.GetUserByEmailReturns(nil, errors.New("nope"))
				})
				It("should have an empty response and throw an error", func() {
					Expect(r).To(BeNil())
					Expect(err).To(HaveOccurred())
				})
			})
		})
		Context("the payload is NOT valid and", func() {
			ctx = context.Background()
			JustBeforeEach(func() {
				p := &usergen.RecoverPasswordPayload{
					Email: "",
				}
				r, err = service.RecoverPassword(ctx, p)
			})
			Context("the email generation fails", func() {
				BeforeEach(func() {
					fakeMail.SendEmailReturns(errors.New("this aint gon work pal"))
					Expect(err).To(HaveOccurred())
				})
			})
		})
	})

	Describe("UploadPicture", func() {
		var (
			ctx      context.Context
			err      error
			response *usergen.OperationStatusResponse
		)

		When("auth details are correct", func() {
			JustBeforeEach(func() {
				ctx = context.Background()
				ctx = context.WithValue(ctx, auth.ContextKeyUser, &auth.User{
					ID:    1,
					Email: "t@example.com",
					Role:  "sfit_user",
				})
				p := &usergen.UploadPicturePayload{
					Name:    "example.jpg",
					Content: []byte{1, 23, 4, 56},
					Format:  "jpg",
				}
				response, err = service.UploadPicture(ctx, p)
			})

			When("picture upload succeeds", func() {
				BeforeEach(func() {
					fakeCloudStorage.UploadPictureReturns(nil)
				})

				When("and picture URL is saved to the database successfully", func() {
					BeforeEach(func() {
						fakeStore.UpdateProfilePictureReturns(nil)
					})
					It("should respond with success", func() {
						Expect(response).ToNot(BeNil())
					})
					It("should NOT return an error", func() {
						Expect(err).ToNot(HaveOccurred())
					})
				})

				When("and picture URL is NOT saved", func() {
					BeforeEach(func() {
						fakeStore.UpdateProfilePictureReturns(errors.New("your mistake, Tony"))
					})
					It("should return an error", func() {
						Expect(err).To(HaveOccurred())
					})
				})
			})

			When("picture upload fails", func() {
				BeforeEach(func() {
					fakeCloudStorage.UploadPictureReturns(errors.New("sorry, you're too ugly"))
				})

				It("should return an error", func() {
					Expect(err).To(HaveOccurred())
				})
			})

		})

		When("auth details are NOT correct", func() {
			JustBeforeEach(func() {
				ctx = context.Background()

				p := &usergen.UploadPicturePayload{
					Name:    "example.jpg",
					Content: []byte{1, 23, 4, 56},
					Format:  "jpg",
				}
				response, err = service.UploadPicture(ctx, p)
			})

			It("should return an error", func() {
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
