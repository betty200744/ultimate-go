package client

type Client struct {
	userId int64
	osType int64
}

func NewClient(userId, osType int64) *Client {
	return &Client{
		userId: userId,
		osType: osType,
	}
}

func (c *Client) UserId() int64 {
	return c.userId
}
func (c *Client) OsType() int64 {
	return c.osType
}
