package authorize

import (
	"context"
	"time"

	"github.com/SKF/proto/common"

	"google.golang.org/grpc"

	authorize_grpcapi "github.com/SKF/proto/authorize"
)

type AuthorizeClient interface { // nolint: golint
	Dial(host, port string, opts ...grpc.DialOption) error
	DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) error
	Close()

	DeepPing() error
	DeepPingWithContext(ctx context.Context) error

	IsAuthorized(userID, action string, resource *common.Origin) (bool, error)
	IsAuthorizedWithContext(ctx context.Context, userID, action string, resource *common.Origin) (bool, error)

	AddAuthorizationResource(resource common.Origin) error
	AddAuthorizationResourceWithContext(ctx context.Context, resource common.Origin) error

	RemoveAuthorizationResource(resource common.Origin) error
	RemoveAuthorizationResourceWithContext(ctx context.Context, resource common.Origin) error

	GetAuthorizationResourcesByType(resourceType string) (resources []common.Origin, err error)
	GetAuthorizationResourcesByTypeWithContext(ctx context.Context, resourceType string) (resources []common.Origin, err error)

	GetResourcesByOriginAndType(originID string, resourceType string) (resources []common.Origin, err error)
	GetResourcesByOriginAndTypeWithContext(ctx context.Context, originID string, resourceType string) (resources []common.Origin, err error)

	GetUserIDsWithAccessToResource(originID string) (resources []string, err error)
	GetUserIDsWithAccessToResourceWithContext(ctx context.Context, originID string) (resources []string, err error)

	AddAuthorizationResourceRelation(resource common.Origin, parent common.Origin) error
	AddAuthorizationResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	RemoveAuthorizationResourceRelation(resource common.Origin, parent common.Origin) error
	RemoveAuthorizationResourceRelationWithContext(ctx context.Context, resource common.Origin, parent common.Origin) error

	GetAuthorizationResourceRelations(resource common.Origin) (resources []common.Origin, err error)
	GetAuthorizationResourceRelationsWithContext(ctx context.Context, resource common.Origin) (resources []common.Origin, err error)

	AddUserPermission(userID, role string, resource common.Origin) error
	AddUserPermissionWithContext(ctx context.Context, userID, role string, resource common.Origin) error

	RemoveUserPermission(userID, role string, resource common.Origin) error
	RemoveUserPermissionWithContext(ctx context.Context, userID, role string, resource common.Origin) error
}

type client struct {
	conn *grpc.ClientConn
	api  authorize_grpcapi.AuthorizeClient
}

func CreateClient() AuthorizeClient {
	return &client{}
}

// Dial creates a client connection to the given host with background context and no timeout
func (c *client) Dial(host, port string, opts ...grpc.DialOption) error {
	return c.DialWithContext(context.Background(), host, port, opts...)
}

// DialWithContext creates a client connection to the given host with context (for timeout and transaction id)
func (c *client) DialWithContext(ctx context.Context, host, port string, opts ...grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, host+":"+port, opts...)
	if err != nil {
		return
	}

	c.conn = conn
	c.api = authorize_grpcapi.NewAuthorizeClient(conn)
	return
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) DeepPing() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.DeepPingWithContext(ctx)
}

func (c *client) DeepPingWithContext(ctx context.Context) error {
	_, err := c.api.DeepPing(ctx, &common.Void{})
	return err
}
