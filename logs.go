package okta

const (
	LOGS_API = "/api/v1/logs"
)

type LogsService struct {
	c *Client
}

func NewLogsService(client *Client) *LogsService {
	return &LogsService{c: client}
}
