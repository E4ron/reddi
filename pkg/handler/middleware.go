package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reddi/models"
)

const headerAuthorization = "Authorization"

func (h *Handler) apiMiddleware(c *gin.Context) {
	h.setUserIdentity(c)
}

func (h *Handler) setUserIdentity(c *gin.Context) {
	session := c.Request.Header.Get(headerAuthorization)
	if session == "" {
		sendUnauthorizedError(c, errors.New("header authorization is empty"))
		return
	}
	account, err := h.services.Session.GetAccount(session)
	if err != nil {
		sendUnauthorizedError(c, err)
	}
	c.Set(gin.AuthUserKey, account)
}

func getAuthAccount(c *gin.Context) *models.Account {
	account, ok := c.Get(gin.AuthUserKey)
	if !ok {
		sendInternalServerErrorGetAutUser(c)
		return nil
	}
	return account.(*models.Account)
}
