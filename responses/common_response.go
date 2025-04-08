package responses

import "github.com/gin-gonic/gin"

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}


type ResponseBuilder struct {
	response CommonResponse
}

func NewResponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{response: CommonResponse{}}
}

func (b *ResponseBuilder) WithCode(code int) *ResponseBuilder {
	b.response.Code = code
	return b
}

func (b *ResponseBuilder) WithMessage(message string) *ResponseBuilder {
	b.response.Message = message
	return b
}

func (b *ResponseBuilder) WithData(data interface{}) *ResponseBuilder {
	b.response.Data = data
	return b
}

func (b *ResponseBuilder) WithErrors(errors interface{}) *ResponseBuilder {
	b.response.Errors = errors
	return b
}

func (b *ResponseBuilder) Build() CommonResponse {
	return b.response
}

func (b *ResponseBuilder) RespondWithJSON(ctx *gin.Context) {
	ctx.JSON(b.response.Code, b.response)
}
