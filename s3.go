package s3

import (
	wof_s3 "github.com/whosonfirst/go-whosonfirst-aws/s3"
	"github.com/whosonfirst/go-whosonfirst-cache"
)

type S3Cache struct {
	cache.Cache
	conn *wof_s3.S3Connection
}

func NewS3Cache(dsn string) (cache.Cache, error) {

	cfg, err := wof_s3.NewS3ConfigFromString(dsn)

	if err != nil {
		return nil, err
	}

	conn, err := wof_s3.NewS3Connection(&config)

	if err != nil {
		return nil, err
	}

	c := S3Cache{
		conn: conn,
	}

	return &c, nil
}

func (c *S3Cache) Get(key string) (io.ReadCloser, error) {

	return c.conn.Get(key)
}

func (c *S3Cache) Set(key string, fh io.ReadCloser) (io.ReadCloser, error) {

	// PLEASE MAKE ME BETTER - presumably with some
	// confusing io.WahWah stuff (20180615/thisisaaronland)

	err := c.conn.Put(key, fh)

	if err != nil {
		return nil, err
	}

	return c.Get(key)
}

func (c *S3Cache) Unset(key string) error {

	return c.conn.Delete(key)
}
