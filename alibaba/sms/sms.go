package sms

import (
	"fmt"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type Client struct {
	client       *dysmsapi20170525.Client
	signName     string
	templateCode string
}

func CreateClient(accessKeyId, accessKeySecret, signName, templateCode string) (*Client, error) {
	config := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: &accessKeyId,
		// 您的AccessKey Secret
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	result, err := dysmsapi20170525.NewClient(config)
	return &Client{client: result, signName: signName, templateCode: templateCode}, err
}

func (c *Client) SendSmsCode(phoneNumber, code string) error {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      &c.signName,
		TemplateCode:  &c.templateCode,
		PhoneNumbers:  tea.String(phoneNumber),
		TemplateParam: tea.String(fmt.Sprintf("{\"code\":\"%s\"}", code)),
	}
	
	result, err := c.client.SendSms(sendSmsRequest)
	if err != nil {
		return err
	}

	if *result.Body.Code != "OK" {
		return fmt.Errorf("send verify code")
	}

	return nil
}
