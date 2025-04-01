package authorizer

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/schema/modules/common"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/server"
	"github.com/dmytro-kucherenko/smartner-utils-package/pkg/types"
	"github.com/google/uuid"
)

func Handle(_ context.Context, event events.APIGatewayCustomAuthorizerRequest) (response events.APIGatewayCustomAuthorizerResponse, err error) {
	// TODO: Handle auth, whitelisted routes, rate limiting

	var id types.ID
	key := "451f4f07-5140-456f-9ffc-4751a808f45f"

	id, err = uuid.Parse(key)
	if err != nil {
		return
	}

	ctx, err := common.EncodeStruct(server.Session{
		UserID: id,
	})

	if err != nil {
		return
	}

	response = events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: "user",
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{"*"},
				},
			},
		},
		Context: ctx,
	}

	return
}
