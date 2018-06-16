package s3

import (
       "github.com/whosonfirst/go-whosonfirst-cache"
       wof_s3 "github.com/whosonfirst/go-whosonfirst-aws/s3"
)

type S3Cache struct {
     cache.Cache
     conn *wof_s3.S3Connection
}

func NewS3Cache() (cache.Cache, error){

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

     	// FIX ME TO RETURN AN io.ReadCloser

	return c.conn.Put(key, fh)
}

func (c *S3Cache) Unset(key string) error {

     return c.conn.Delete(key)
}
