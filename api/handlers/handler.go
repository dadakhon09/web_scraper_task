package handlers

import (
	"github.com/dadakhon09/web_scraper_task/api/models"
	"github.com/dadakhon09/web_scraper_task/config"
	"github.com/dadakhon09/web_scraper_task/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handlerV1 struct {
	log logger.Logger
	cfg *config.Config
}

type HandlerV1Options struct {
	Log logger.Logger
	Cfg *config.Config
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		log: options.Log,
		cfg: options.Cfg,
	}
}

const (
	ErrorCodeInvalidURL        = "INVALID_URL"
	ErrorCodeInvalidJSON       = "INVALID_JSON"
	ErrorCodeInternal          = "INTERNAL_SERVER_ERROR"
	ErrorCodeUnauthorized      = "UNAUTHORIZED"
	ErrorCodeAlreadyExists     = "ALREADY_EXISTS"
	ErrorCodeNotFound          = "NOT_FOUND"
	ErrorCodeInvalidCode       = "INVALID_CODE"
	ErrorBadRequest            = "BAD_REQUEST"
	ErrorCodeForbidden         = "FORBIDDEN"
	ErrorCodeNotApproved       = "NOT_APPROVED"
	ErrorCodePasswordsNotEqual = "PASSWORDS_NOT_EQUAL"
	ErrServiceUnavailable      = "SERVICE_UNAVAILABLE"
	ErrInvalidArgument         = "INVALID_ARGUMENT"
)

func (h *handlerV1) handleError(c *gin.Context, err error, message string) bool {
	st, ok := status.FromError(err)

	switch st.Code() {
	case codes.Internal:
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternal,
			Message: st.Message(),
		})
		h.log.Error(message+", internal server error", logger.Error(err))

		return true

	case codes.Unavailable:
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrServiceUnavailable,
			Message: st.Message(),
		})
		h.log.Error(message+", service unavailable", logger.Error(err))

		return true

	case codes.AlreadyExists:
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeAlreadyExists,
			Message: st.Message(),
		})
		h.log.Error(message+", already exists", logger.Error(err))

		return true

	case codes.InvalidArgument:
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrInvalidArgument,
			Message: st.Message(),
		})
		h.log.Error(message+", invalid field", logger.Error(err))

		return true

	case codes.NotFound:
		c.JSON(http.StatusNotFound, models.ResponseError{
			Code:    ErrorCodeNotFound,
			Message: st.Message(),
		})
		h.log.Error(message+", not found", logger.Error(err))

		return true

	default:
		if err != nil || !ok {
			c.JSON(http.StatusInternalServerError, models.ResponseError{
				Code:    ErrorCodeInternal,
				Message: st.Message(),
			})
			h.log.Error(message+", unknown error", logger.Error(err))
			return true
		}

	}

	return false
}

func handleInternalWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{
			Code:    ErrorCodeInternal,
			Message: "Internal Server Error",
		})
		l.Error(message, logger.Error(err))
		return true
	}

	return false
}

func handleBadRequestErrWithMessage(c *gin.Context, l logger.Logger, err error, message string) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{
			Code:    ErrorCodeInvalidJSON,
			Message: "Invalid Json",
		})
		l.Error(message, logger.Error(err))
		return true
	}
	return false
}
