package response

import (
	"acl-casbin/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Builder struct {
	c          *gin.Context
	code       int
	success    bool
	message    string
	messageID  string
	traceID    string
	data       interface{}
	errors     interface{}
	pagination *PaginationMeta
	extra      map[string]interface{}
}

func New(c *gin.Context) *Builder {
	traceID := c.GetString("trace_id")
	if traceID == "" {
		traceID = uuid.NewString()
	}
	return &Builder{
		c:       c,
		traceID: traceID,
		code:    http.StatusOK,
		success: true,
	}
}

func (b *Builder) Status(code int) *Builder {
	b.code = code
	b.success = code < 400
	return b
}

func (b *Builder) Message(msg string) *Builder {
	b.message = msg
	return b
}

func (b *Builder) MessageID(msgID string) *Builder {
	b.messageID = msgID
	return b
}

func (b *Builder) Data(data interface{}) *Builder {
	b.data = data
	return b
}

func (b *Builder) Errors(err error) *Builder {
	b.errors = err
	b.success = false
	if b.code == http.StatusOK {
		b.code = http.StatusBadRequest
	}
	if validationError := utils.GetValidationErrors(err); validationError != nil {
		b.errors = validationError
	} else {
		b.errors = []map[string]interface{}{
			{
				"message": err.Error(),
			},
		}
	}
	return b
}

func (b *Builder) Pagination(p *PaginationMeta) *Builder {
	b.pagination = p
	return b
}

func (b *Builder) Extra(extra map[string]interface{}) *Builder {
	b.extra = extra
	return b
}

type Meta struct {
	Success    bool                   `json:"success"`
	Code       int                    `json:"code"`
	Message    string                 `json:"message"`
	MessageID  string                 `json:"message_id,omitempty"` // For i18n
	TraceID    string                 `json:"trace_id,omitempty"`
	Pagination *PaginationMeta        `json:"pagination,omitempty"`
	Extra      map[string]interface{} `json:"extra,omitempty"` // For additional metadata
}

type PaginationMeta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type Response struct {
	Meta   Meta        `json:"meta"`
	Data   interface{} `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}

func (b *Builder) Dispatch() {
	meta := Meta{
		Success:    b.success,
		Code:       b.code,
		Message:    b.message,
		MessageID:  b.messageID,
		TraceID:    b.traceID,
		Pagination: b.pagination,
		Extra:      b.extra,
	}

	resp := Response{
		Meta:   meta,
		Data:   b.data,
		Errors: b.errors,
	}

	b.c.JSON(b.code, resp)
}
