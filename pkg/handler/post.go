package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reddi/models"
	"strconv"
)

func (h *Handler) GetPostById(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c, errors.New("invalid item_id"))
		return
	}

	post, err := h.services.Post.GetById(id)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, post)
}
func (h *Handler) GetList(c *gin.Context) {
	page := c.Param("page")
	limit := c.Param("limit")
	if page == "" || limit == "" {
		sendBadRequestError(c, errors.New("invalid item_id"))
	}
	intPage, err := strconv.Atoi(page)
	if err != nil {
		sendBadRequestError(c, err)
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		sendBadRequestError(c, err)
	}
	if intPage <= 0 || intLimit <= 0 {
		sendBadRequestError(c, err)
	}

	output, err := h.services.Post.GetList(intPage, intLimit)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, output)

}
func (h *Handler) Create(c *gin.Context) {
	var input models.InputPost

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}

	output, err := h.services.Post.Create(&input)
	if err != nil {
		sendInternalServerError(c, err)
		return
	}
	c.JSON(http.StatusOK, output)
}
func (h *Handler) Update(c *gin.Context) {
	var input models.InputUpdatePost

	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c, errors.New("invalid item_id"))
		return
	}
	input.Id = id

	if err := c.BindJSON(&input); err != nil {
		sendBadRequestError(c, err)
		return
	}
	if err := h.services.Post.Update(&input); err != nil {
		sendInternalServerError(c, err)
		return
	}

	c.Status(http.StatusOK)

}
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("item_id")
	if id == "" {
		sendBadRequestError(c, errors.New("invalid item_id"))
		return
	}
	if err := h.services.Post.Delete(id); err != nil {
		sendInternalServerError(c, err)
		return
	}
	c.Status(http.StatusOK)

}
