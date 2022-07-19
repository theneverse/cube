package oss

import (
	"fmt"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Client struct {
	*oss.Client
	bucket      *oss.Bucket
	pubEndPoint string
}

func NewClient(endPoint, accessKeyId, accessKeySecret, bucketName, pubEndPoint string) (*Client, error) {
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}
	return &Client{Client: client, bucket: bucket, pubEndPoint: pubEndPoint}, err
}

func (c *Client) PutObj(name string, dir string, reader io.Reader) (string, error) {
	path := fmt.Sprintf("%s/%s", dir, name)
	if err := c.bucket.PutObject(path, reader); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", c.pubEndPoint, path), nil
}

func (c *Client) CreateOssBucket(bucketName string, options ...oss.Option) error {
	return c.CreateBucket(bucketName, options...)
}
func (c *Client) BucketExist(bucketName string) (bool, error) {
	exist, err := c.IsBucketExist(bucketName)
	if err != nil {
		return false, err
	}
	return exist, nil
}
