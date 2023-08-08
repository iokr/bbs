package orm

type Option func(cli *Client)

func WithPool(pool *Pool) Option {
	return func(cli *Client) {
		cli.pool = pool
	}
}
