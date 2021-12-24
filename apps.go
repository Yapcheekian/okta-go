package okta

const (
	APPS_API = "/api/v1/apps"
)

type AppsService struct {
	c *Client
}

func NewAppsService(client *Client) *AppsService {
	return &AppsService{c: client}
}
