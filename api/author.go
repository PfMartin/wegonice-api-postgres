package api

import (
	"database/sql"
	"net/http"

	db "github.com/PfMartin/wegonice-api/db/sqlc"
	"github.com/PfMartin/wegonice-api/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// TODO: Create custom validators for Website, Instagram and YouTube
type createAuthorRequest struct {
	AuthorName string `json:"author_name" binding:"required"`
	Website    string `json:"website"`
	Instagram  string `json:"instagram"`
	Youtube    string `json:"youtube"`
}

func (server *Server) createAuthor(ctx *gin.Context) {
	var req createAuthorRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAuthorParams{
		AuthorName:  req.AuthorName,
		Website:     req.Website,
		Instagram:   req.Instagram,
		Youtube:     req.Youtube,
		UserCreated: authPayload.Email,
	}

	author, err := server.store.CreateAuthor(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

type getAuthorRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAuthor(ctx *gin.Context) {
	var req getAuthorRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	author, err := server.store.GetAuthor(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

type deleteAuthorRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteAuthor(ctx *gin.Context) {
	var req deleteAuthorRequest
	if err := ctx.BindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	author, err := server.store.DeleteAuthorById(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

type updateAuthorUriRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateAuthorBodyRequest struct {
	AuthorName string `json:"author_name" binding:"required"`
	Website    string `json:"website"`
	Instagram  string `json:"instagram"`
	Youtube    string `json:"youtube"`
}

// TODO: Use ctx.AbortWithError instead of ctx.JSON
func (server *Server) updateAuthor(ctx *gin.Context) {
	var reqUri updateAuthorUriRequest
	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var reqBody updateAuthorBodyRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAuthorByIdParams{
		ID:         reqUri.ID,
		AuthorName: reqBody.AuthorName,
		Website:    reqBody.Website,
		Instagram:  reqBody.Instagram,
		Youtube:    reqBody.Youtube,
	}

	// TODO: Add unique constraint for name column in authors table
	author, err := server.store.UpdateAuthorById(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, author)
}

type listAuthorsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=50"`
}

func (server *Server) listAuthors(ctx *gin.Context) {
	var req listAuthorsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAuthorsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	authors, err := server.store.ListAuthors(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, authors)
}
