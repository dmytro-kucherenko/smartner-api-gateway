package authorizer

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/meta"
	"github.com/google/uuid"
)

const TokenKey string = "authorization" // Get from template env

func respond(principalID, methodArn, effect string, context map[string]any) (events.APIGatewayCustomAuthorizerResponse, error) {
	return events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: principalID,
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{methodArn},
				},
			},
		},
		Context: context,
	}, nil
}

func deny(principalID, methodArn, message string) (events.APIGatewayCustomAuthorizerResponse, error) {
	return respond(principalID, methodArn, "Deny", map[string]any{
		"error": message,
	})
}

func allow(principalID, methodArn string, session meta.Session) (events.APIGatewayCustomAuthorizerResponse, error) {
	data := meta.SetSessionHeader(session)

	return respond(principalID, methodArn, "Allow", data)
}

func Handle(_ context.Context, event events.APIGatewayCustomAuthorizerRequest) (response events.APIGatewayCustomAuthorizerResponse, err error) {
	cookies, err := http.ParseCookie(event.AuthorizationToken)
	if err != nil {
		return deny("user", event.MethodArn, "cookies are invalid")
	}

	i := slices.IndexFunc(cookies, func(cookie *http.Cookie) bool { return strings.EqualFold(cookie.Name, TokenKey) })
	if i < 0 {
		return deny("user", event.MethodArn, fmt.Sprintf("%v cookie is not found", TokenKey))
	}

	cookie := cookies[i]
	token := cookie.Value

	// TODO: Handle auth - get session in response where handle token value and expiry, rate limiting
	id, err := uuid.Parse(token)
	if err != nil {
		return deny("user", event.MethodArn, fmt.Sprintf("%v cookie is invalid", TokenKey))
	}

	session := meta.Session{
		UserID: id,
	}

	return allow("user", event.MethodArn, session)
}
