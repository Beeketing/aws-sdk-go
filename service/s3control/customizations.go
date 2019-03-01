package s3control

import (
	"github.com/Beeketing/aws-sdk-go/aws/client"
	"github.com/Beeketing/aws-sdk-go/aws/request"
	"github.com/Beeketing/aws-sdk-go/internal/s3err"
	"github.com/Beeketing/aws-sdk-go/private/protocol"
)

type accountIDGetter interface {
	getAccountId() string
}

func init() {
	initClient = defaultInitClientFn
}

func defaultInitClientFn(c *client.Client) {
	c.Handlers.UnmarshalError.PushBackNamed(s3err.RequestFailureWrapperHandler())
}

func buildPrefixHostHandler(fieldName, value string) request.NamedHandler {
	return request.NamedHandler{
		Name: "awssdk.s3control.prefixhost",
		Fn: func(r *request.Request) {
			paramErrs := request.ErrInvalidParams{Context: r.Operation.Name}
			if !protocol.ValidHostLabel(value) {
				paramErrs.Add(request.NewErrParamFormat(fieldName, "[a-zA-Z0-9-]{1,63}", value))
				r.Error = paramErrs
				return
			}
			r.HTTPRequest.URL.Host = value + "." + r.HTTPRequest.URL.Host
		},
	}
}
func buildRemoveHeaderHandler(key string) request.NamedHandler {
	return request.NamedHandler{
		Name: "awssdk.s3control.removeHeader",
		Fn: func(r *request.Request) {
			r.HTTPRequest.Header.Del(key)
		},
	}
}
