package routes

import (
	"net/http"

	"github.com/Walon-Foundation/go-gin-doc/cmd/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	gonanoid"github.com/matoous/go-nanoid/v2"
)



// RegisterUser godoc
// @Summary     Register a new user
// @Description Create a new user account using name, username, and password
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       body  body      object{name=string,username=string,password=string}  true  "Register payload"
// @Success     201   {object}  utils.SuccessResponseSchema  "User created successfully"
// @Failure     400   {object}  utils.ErrorResponseSchema    "Invalid request body"
// @Failure     409   {object}  utils.ErrorResponseSchema    "User already exists"
// @Failure     500   {object}  utils.ErrorResponseSchema    "Internal server error"
// @Router      /auth/signup [post]
func(r *route) RegisterUser(c *gin.Context){
	//handler logic
	type user struct {
		Name string `json:"name" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}
	
	var registerRequest user
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request body")
		return
	}
	
	ctx := c.Request.Context()
	
	row := r.Models.User.GetUser(ctx,registerRequest.Name, registerRequest.Username)
	var tempName string
	if err := row.Scan(&tempName); err != pgx.ErrNoRows && len(tempName) != 0 {
		utils.ErrorResponse(
			c,
			http.StatusConflict,
			"User with email and username already exist",
		)
		
		return
	}
	
	passwordHash,err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password),10)
	if err != nil {
		utils.ErrorResponse(
			c, 
			http.StatusInternalServerError,
			"Internal server error",
		)
		return
	}
	
	id, _ := gonanoid.New(20)
	
	row = r.Models.User.InsertUser(ctx, id,registerRequest.Name, registerRequest.Username, string(passwordHash))
	var tempId string
	if err := row.Scan(&tempId); err != nil && len(tempId) == 0 {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"Failed to save user",
		)
		return
	}
	
	utils.SuccessResponse(
		c,
		http.StatusCreated,
		"User created",
		tempId,
	)
}




// LoginUser godoc
// @Summary     Login user
// @Description Authenticate a user using username and password
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Param       body  body      object{username=string,password=string}  true  "Login payload"
// @Success     200   {object}  utils.SuccessResponseSchema  "Login successful"
// @Failure     400   {object}  utils.ErrorResponseSchema    "Invalid request body"
// @Failure     401   {object}  utils.ErrorResponseSchema    "Invalid credentials"
// @Failure     500   {object}  utils.ErrorResponseSchema    "Internal server error"
// @Router      /auth/login [post]
func (r *route) LoginUser(c *gin.Context){
	//handler logic
	type user struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8"`
	}
	
	var loginrequest user
	
	if err := c.ShouldBindJSON(&loginrequest); err != nil {
		utils.ErrorResponse(
			c, 
			http.StatusBadRequest,
			"invalid request body",
		)
		return
	}
	
	ctx := c.Request.Context()
	
	row := r.Models.User.GetUserByUsername(ctx, loginrequest.Username)
	var tempName,tempPassword, tempId string
	if err := row.Scan(&tempId,&tempName,&tempPassword); err == pgx.ErrNoRows && len(tempName) == 0 {
		utils.ErrorResponse(
			c,
			http.StatusUnauthorized,
			"invalid credentials",
		)
		
		return
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(tempPassword), []byte(loginrequest.Password)); err != nil {
		utils.ErrorResponse(
			c,
			http.StatusUnauthorized,
			"invalid login credentials",
		)
		return 
	}
	
	//Todo: create a jwt token and send to the user
	token, err := utils.CreateToken(tempId)
	if err != nil {
		utils.ErrorResponse(
			c,
			http.StatusInternalServerError,
			"failed to created token",
		)
		return
	}
	
	utils.SuccessResponse(
		c,
		http.StatusOK,
		"Login successfull",
		token,
	)
}