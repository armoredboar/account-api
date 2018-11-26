package controllers

import (
	"net/http"
	"strings"

	"github.com/armoredboar/account-api/pkg/uuid"

	"github.com/armoredboar/account-api/internal/server/contracts"
	"github.com/armoredboar/account-api/internal/server/models"
	"github.com/armoredboar/account-api/internal/validation"

	"github.com/gin-gonic/gin"
)

func Authorize(c *gin.Context) {

	responseType := strings.TrimSpace(c.Query("response_type"))
	clientID := strings.TrimSpace(c.Query("client_id"))
	redirectURI := strings.TrimSpace(c.Query("redirect_uri"))
	scope := strings.TrimSpace(c.Query("scope"))
	state := strings.TrimSpace(c.Query("state"))

	parameters := responseType + clientID + redirectURI + scope + state

	c.JSON(http.StatusOK, parameters)
}

func AccessToken(c *gin.Context) {

	var response contracts.AccessTokenResponse

	var accessToken models.AccessToken
	if err := c.Bind(&accessToken); err != nil {
		response.Report.AddError(validation.CouldNotBeParsed, "post body")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validates the received parameters.
	if errors := accessToken.Validate(); len(errors) > 0 {
		response.Report.Errors = errors
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// TODO: Check grant_type
	// TODO: Validate ClientID and ClientSecret
	// TODO: Validate username ans password

	response.AccessToken, _ = uuid.CreateUUID()
	response.ExpiresIn = 300
	response.RefreshToken = "refresh_token"
	response.TokenType = "Bearer"

	response.Success = true

	c.JSON(http.StatusOK, response)
}
