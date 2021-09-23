package response

type Response struct {
	Status  int         `json:"status"`
	Message []string    `json:"message"`
	Data    interface{} `json:"data"`
}
type IResponse interface {
	ResponseSuccess(statuscode int, message []string, data interface{}) *Response
	ResponseError(statuscode int, message []string, data interface{}) *Response
}

func InitResponse() IResponse {
	return &Response{}
}

func (r *Response) ResponseSuccess(statuscode int, message []string, data interface{}) *Response {
	return &Response{
		Status:  statuscode,
		Message: message,
		Data:    data,
	}
}

func (r *Response) ResponseError(statuscode int, message []string, data interface{}) *Response {

	return &Response{
		Status:  statuscode,
		Message: message,
		Data:    data,
	}
}
