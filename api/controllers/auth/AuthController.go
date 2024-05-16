package authcontrollers

import (
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
	"user-services/api/config"
	masterentities "user-services/api/entities/master"
	"user-services/api/exceptions"
	"user-services/api/helper"
	jsonchecker "user-services/api/helper/json/json-checker"
	"user-services/api/payloads"
	redisservices "user-services/api/services/redis"
	userservices "user-services/api/services/user"
	"user-services/api/utils/validation"
)

type AuthController interface {
	Login(writer http.ResponseWriter, request *http.Request)
	Register(writer http.ResponseWriter, request *http.Request)
	AuthLogin(writer http.ResponseWriter, request *http.Request)
	//ForgotPassword(writer http.ResponseWriter, request *http.Request)
	//ResetPassword(writer http.ResponseWriter, request *http.Request)
	//ChangePassword(writer http.ResponseWriter, request *http.Request)
	//GenerateOTP(writer http.ResponseWriter, request *http.Request)
	//VerifyOTP(writer http.ResponseWriter, request *http.Request)
	//Logout(writer http.ResponseWriter, request *http.Request)
}

type AuthControllerImpl struct {
	UserService  userservices.UserService
	AuthService  userservices.AuthService
	RedisService redisservices.RedisService
}

func NewAuthController(
	authService userservices.AuthService,
	userService userservices.UserService,
	redisService redisservices.RedisService,
) AuthController {
	return &AuthControllerImpl{
		UserService:  userService,
		AuthService:  authService,
		RedisService: redisService,
	}
}

// AuthLogin Get All Bining List Via Header
//
//	@Summary		Login With User
//	@Description	Login With User
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		payloads.LoginRequestPayloads	true	"Insert Header Request"
//	@Success		200		{object}	payloads.Respons
//
// @Router	/auth/loginAuth [post]
func (controller *AuthControllerImpl) AuthLogin(writer http.ResponseWriter, request *http.Request) {
	loginRequest := payloads.LoginRequestPayloads{}
	err := jsonchecker.ReadFromRequestBody(request, &loginRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}

	err = validation.ValidationForm(writer, request, &loginRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}
	loginReq, errs := controller.AuthService.LoginAuth(loginRequest)

	if errs != nil {
		exceptions.NewBadRequestException(writer, request, errs)
		return
	}
	expTime := time.Now().Add(time.Minute * 1)
	claims := config.JWTClaim{
		UserName: loginReq.UserName,
		UserRole: loginReq.UserRoleId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "devin",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, erronToken := tokenAlgo.SignedString(config.JWT_KEY)
	if erronToken != nil {
		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Message:    "tokenExpired",
			Data:       nil,
			Err:        erronToken,
		})
	}
	//payloads.ResponseToken(writer, loginReq, "Login Successfully", http.StatusOK)
	//generatedToken(loginReq)
	payloads.HandleSuccess(writer, convertReqToRes(loginReq, token), "Register Success", http.StatusCreated)

}
func convertReqToRes(req masterentities.UserEntities, token string) payloads.LoginResponses {
	return payloads.LoginResponses{
		UserName: req.UserName,
		UserRole: req.UserRoleId,
		Token:    token,
	}
}
func (controller *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request) {
	loginRequest := payloads.LoginRequest{}
	err := jsonchecker.ReadFromRequestBody(request, &loginRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}

	err = validation.ValidationForm(writer, request, &loginRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}
	loginReq, errs := controller.AuthService.Login(loginRequest)

	if errs != nil {
		exceptions.NewBadRequestException(writer, request, errs)
		return
	}

	//payloads.ResponseToken(writer, loginReq, "Login Successfully", http.StatusOK)
	//generatedToken(loginReq)
	payloads.HandleSuccess(writer, loginReq, "Register Success", http.StatusCreated)

}

// Register Get All Bining List Via Header
//
//	@Summary		Register User
//	@Description	REST API User
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		payloads.RegisterRequest	true	"Insert Register Request"
//	@Success		200		{object}	payloads.Respons
//
// @Router	/auth/register [post]
func (controller *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request) {
	registerRequest := payloads.RegisterRequest{}
	role := chi.URLParam(request, "role")

	err := jsonchecker.ReadFromRequestBody(request, &registerRequest)
	if err != nil {
		exceptions.NewEntityException(writer, request, err)
		return
	}

	err = validation.ValidationForm(writer, request, &registerRequest)
	if err != nil {
		exceptions.NewBadRequestException(writer, request, err)
		return
	}
	if role == "admin" {
		registerRequest.UserRoleId = 1
	} else if role == "hash-micro" {
		registerRequest.UserRoleId = 2
	} else {
		panic("")
		return
	}
	created, err := controller.AuthService.Register(registerRequest)
	if err != nil {
		helper.ReturnError(writer, request, err)
		return
	}
	payloads.HandleSuccess(writer, created, "Register Success", http.StatusCreated)

}

//	@Summary		Change Password User
//	@Description	REST API User
//	@Accept			json
//	@Produce		json
//	@Tags			Auth Controller
//	@Security		BearerAuth
//	@Param			reqBody	body		payloads.ChangePasswordInput	true	"Form Request"
//	@Success		200		{object}	payloads.Respons
//	@Router			/auth/password/change [put]
//func (controller *AuthControllerImpl) ChangePassword(writer http.ResponseWriter, request *http.Request) {
//	changePasswordRequest := payloads.ChangePasswordInput{}
//	err := jsonchecker.ReadFromRequestBody(request, &changePasswordRequest)
//	if err != nil {
//		exceptions.NewEntityException(writer, request, err)
//		return
//	}
//	claims, _ := securities.ExtractAuthToken(request)
//
//	session, err := controller.AuthService.UpdatePassword(claims, changePasswordRequest)
//	if err != nil {
//		helper.ReturnError(writer, request, err)
//		return
//	}
//	payloads.HandleSuccess(writer, session, "Change Password Success", http.StatusOK)
//
//}

//	@Summary		Generate OTP User
//	@Description	REST API OTP User
//	@Accept			json
//	@Produce		json
//	@Tags			Auth Controller
//	@Security		BearerAuth
//	@Success		200	{object}	payloads.Respons

//	@Router	/auth/generate [post]
//func (controller *AuthControllerImpl) GenerateOTP(writer http.ResponseWriter, request *http.Request) {
//
//	claims, _ := securities.ExtractAuthToken(request)
//
//	fileName, err := controller.AuthService.GenerateOTP(claims.UserID)
//
//	if err != nil {
//		helper.ReturnError(writer, request, err)
//		return
//	}
//
//	http.ServeFile(writer, request, fileName)
//}

//	@Summary		Verify OTP User
//	@Description	REST API Verify OTP User
//	@Accept			json
//	@Produce		json
//	@Tags			Auth Controller
//	@Param			reqBody	body		entities.OTPInput	true	"Form Request"
//	@Success		200		{object}	payloads.Respons

//	@Router	/auth/verify [post]
//func (controller *AuthControllerImpl) VerifyOTP(writer http.ResponseWriter, request *http.Request) {
//	otpInput := masterentities.OTPInput{}
//	err := jsonchecker.ReadFromRequestBody(request, &otpInput)
//	if err != nil {
//		exceptions.NewEntityException(writer, request, err)
//		return
//	}
//
//	response, err := controller.AuthService.UpdateUserOTP(otpInput, request.RemoteAddr)
//	if err != nil {
//		helper.ReturnError(writer, request, err)
//		return
//	}
//	payloads.ResponseToken(writer, response, "Login Successfully", http.StatusOK)
//}

//func (controller *AuthControllerImpl) ValidateOTP(writer http.ResponseWriter, request *http.Request, userReq payloads.LoginRequest) {
//	otpInput := masterentities.OTPInput{}
//	err := jsonchecker.ReadFromRequestBody(request, &otpInput)
//
//	if err != nil {
//		exceptions.NewEntityException(writer, request, err)
//		return
//	}
//	getUser, err := controller.UserService.GetUser(userReq.Username)
//
//	if err != nil {
//		helper.ReturnError(writer, request, err)
//		return
//	}
//
//	valid := totp.Validate(otpInput.Token, getUser.OtpSecret)
//	if !valid {
//		exceptions.NewBadRequestException(writer, request, &exceptions.BaseErrorResponse{
//			Err: errors.New("OTP is not valid"),
//		})
//		return
//	}
//
//	payloads.HandleSuccess(writer, valid, "OTP Valid", http.StatusOK)
//}

//	@Summary		Forgot Password User
//	@Description	REST API Password User
//	@Accept			json
//	@Produce		json
//	@Tags			Auth Controller
//	@Param			reqBody	body		payloads.ForgotPasswordInput	true	"Form Request"
//	@Success		200		{object}	payloads.Respons

//	@Router	/auth/forgot-password [post]
//func (controller *AuthControllerImpl) ForgotPassword(writer http.ResponseWriter, request *http.Request) {
//	forgotPasswordReq := payloads.ForgotPasswordInput{}
//	err := jsonchecker.ReadFromRequestBody(request, &forgotPasswordReq)
//
//	if err != nil {
//		exceptions.NewEntityException(writer, request, err)
//	}
//
//	err = validation.ValidationForm(writer, request, &forgotPasswordReq)
//	if err != nil {
//		exceptions.NewBadRequestException(writer, request, err)
//		return
//	}
//	message := "You will receive a reset email if user with that email exist"
//
//	// Generate Verification Code
//	resetToken := randstr.String(100)
//
//	passwordResetToken := utils.Encode(resetToken)
//
//	_, err = controller.AuthService.UpdatePasswordTokenByEmail(payloads.UpdateEmailTokenRequest{
//		Email:              forgotPasswordReq.Email,
//		PasswordResetToken: utils.StringPtr(passwordResetToken),
//		PasswordResetAt:    utils.TimePtr(time.Now().Add(time.Minute * 15)),
//	})
//
//	if err != nil {
//		exceptions.NewAppException(writer, request, err)
//		return
//	}
//	// Send Email
//	emailData := email.EmailData{
//		URL:     config.EnvConfigs.ClientOrigin + "/auth/password/reset/" + resetToken,
//		Subject: "Your password reset token (valid for 15 min)",
//	}
//	sendEmail, errors := email.SendEmail(forgotPasswordReq.Email, &emailData, "ResetPassword.html")
//
//	if errors != nil {
//		exceptions.NewAppException(writer, request, &exceptions.BaseErrorResponse{
//			Err: errors,
//		})
//		return
//	}
//	payloads.HandleSuccess(writer, sendEmail, message, http.StatusOK)
//}

//	@Summary		Forgot Password User
//	@Description	REST API Password User
//	@Accept			json
//	@Produce		json
//	@Tags			Auth Controller
//	@Param			reqBody		body		payloads.ResetPasswordInput	true	"Form Request"
//	@Param			reset_token	path		string						true	"Reset Token"
//	@Success		200			{object}	payloads.Respons
//	@Router			/auth/password/reset/{reset_token} [patch]
//func (controller *AuthControllerImpl) ResetPassword(writer http.ResponseWriter, request *http.Request) {
//	forgotPasswordReq := payloads.ResetPasswordInput{}
//	err := jsonchecker.ReadFromRequestBody(request, &forgotPasswordReq)
//	if err != nil {
//		exceptions.NewEntityException(writer, request, err)
//		return
//	}
//	resetToken := chi.URLParam(request, "reset_token")
//
//	_, err = controller.AuthService.ResetPassword(resetToken, forgotPasswordReq)
//
//	if err != nil {
//		helper.ReturnError(writer, request, err)
//		return
//	}
//
//	payloads.HandleSuccess(writer, true, "Password reset Successfully", http.StatusOK)
//
//}

//	@Summary		Logout User
//	@Description	REST API Logout User
//	@Accept			json
//	@Produce		json
//	@Tags			Auth Controller
//	@Security		BearerAuth
//	@Success		200	{object}	payloads.Respons
//	@Router			/auth/logout [post]
//func (controller *AuthControllerImpl) Logout(writer http.ResponseWriter, request *http.Request) {
//
//	claims, _ := securities.ExtractAuthToken(request)
//
//	controller.RedisService.DeleteCredential(claims.UserID)
//
//	payloads.HandleSuccess(writer, true, "Logout Success", http.StatusOK)
//}
