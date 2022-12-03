package handlers

import (
	"net/http"
	"strconv"

	"http-server/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)



// CreateAuthor godoc
// @Summary     Create author
// @Description create a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthorModel true "author body"
// @Success     201    {object} models.JSONResponse{data=models.Author}
// @Failure     400    {object} models.JSONErrorResponce
// @Failure     500    {object} models.JSONErrorResponce
// @Router      /v1/author [post]
func (h handler) CreateAuthor(c *gin.Context) {
	var body models.CreateAuthorModel

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		return
	}
	//ToDo - validation should be here

	id := uuid.New()
	err := h.Stg.AddAuthor(id.String(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	author, err := h.Stg.GetAuthorById(id.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Author | Created",
		Data:    author,
	})
}

// GetAuthorById godoc
// @Summary     get author by id
// @Description get an author by id
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Author ID"
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     400 {object} models.JSONErrorResponce
// @Router      /v1/author/{id} [get]
func (h handler) GetAuthorById(c *gin.Context) {

	idStr := c.Param("id")

	author, err := h.Stg.GetAuthorById(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    author,
	})

}

// getAuthorList godoc
// @Summary     List authors
// @Description get authors
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset query    string false "0"
// @Param       limit  query    string false "10"
// @Param       search query    string false "smth"
// @Failure     400    {object} models.JSONErrorResponce
// @Success     200    {object} models.JSONResponse{data=[]models.Author}
// @Router      /v1/author [get]
func (h handler) GetAuthorList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", "0")
	limitStr := c.DefaultQuery("limit", "10")
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	authorList, err := h.Stg.GetAuthorList(offset, limit, searchStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Ok",
		Data:    authorList,
	})
}

// UpdateAuthor...
// UpdateAuthor godoc
// @Summary     Update author
// @Description update a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.UpdateAuthorModel true "author body"
// @Success     200    {object} models.JSONResponse{data=[]models.Author}
// @Failure     400    {object} models.JSONErrorResponce
// @Router      /v1/author [put]
func (h handler) UpdateAuthor(c *gin.Context) {

	var body models.UpdateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{Error: err.Error()})
		return
	}
	err := h.Stg.UpdateAuthor(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	author, err := h.Stg.GetAuthorById(body.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "ok",
		Data:    author,
	})
}

// DeleteAuthor...
// DeleteAuthor godoc
// @Summary     delete author by id
// @Description delete an author by id
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Author ID"
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     400 {object} models.JSONErrorResponce
// @Router      /v1/author/{id} [delete]
func (h handler) DeleteAuthor(c *gin.Context) {

	idStr := c.Param("id")

	author, err := h.Stg.GetAuthorById(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	err = h.Stg.RemoveAuthor(author.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponce{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "Ok",
		Data:    author,
	})

}
