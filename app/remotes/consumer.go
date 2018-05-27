package remotes

import (
	"fmt"
	"wmqx-ui/app/utils"
	"encoding/json"
	"errors"
)

var (
	consumerAddPath = "/consumer/add"
	consumerUpdatePath = "/consumer/update"
	consumerDeletePath = "/consumer/delete"
	consumerStatusPath = "/consumer/status"
	getConsumerByIdPath = "/consumer/getConsumerById"
)

func NewConsumerByNode(node map[string]string, messageName string) *Consumer {
	return NewConsumer(node["manager_uri"], node["token_header_name"], node["token"], messageName)
}

func NewConsumer(managerUri string, tokenHeader string, token string, messageName string) *Consumer {
	return &Consumer{
		ManagerUri: managerUri,
		TokenHeaderName: tokenHeader,
		Token: token,
		MessageName: messageName,
	}
}

type Consumer struct {
	ManagerUri string
	TokenHeaderName string
	Token string
	MessageName string
}

func (c *Consumer) AddConsumer(consumer map[string]string) (err error) {

	url := fmt.Sprintf("%s%s", c.ManagerUri, consumerAddPath)

	headerValue := map[string]string{
		c.TokenHeaderName: c.Token,
	}
	consumer["name"] = c.MessageName

	body, code, err := utils.Request.HttpPost(url, consumer, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return errors.New(fmt.Sprintf(v["message"].(string)))
	}

	return nil
}

func (c *Consumer) UpdateConsumerByConsumerId(consumerId string, consumer map[string]string) (err error) {

	url := fmt.Sprintf("%s%s", c.ManagerUri, consumerUpdatePath)

	headerValue := map[string]string{
		c.TokenHeaderName: c.Token,
	}

	consumer["name"] = c.MessageName
	consumer["consumer_id"] = consumerId

	body, code, err := utils.Request.HttpPost(url, consumer, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return errors.New(fmt.Sprintf(v["consumer"].(string)))
	}

	return nil
}

func (c *Consumer) DeleteConsumerByConsumerId(consumerId string) (err error) {

	url := fmt.Sprintf("%s%s", c.ManagerUri, consumerDeletePath)

	headerValue := map[string]string{
		c.TokenHeaderName: c.Token,
	}
	queryValue := map[string]string{
		"name": c.MessageName,
		"consumer_id": consumerId,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return errors.New(fmt.Sprintf(v["consumer"].(string)))
	}

	return nil
}

func (c *Consumer) GetStatusByConsumerId(consumerId string) (consumerStatus map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", c.ManagerUri, consumerStatusPath)

	headerValue := map[string]string{
		c.TokenHeaderName: c.Token,
	}
	queryValue := map[string]string{
		"name": c.MessageName,
		"consumer_id": consumerId,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return consumerStatus, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return consumerStatus, errors.New(fmt.Sprintf(v["consumer"].(string)))
	}
	consumerStatus = v["data"].(map[string]interface{})

	return
}

func (c *Consumer) GetConsumerByConsumerId(consumerId string) (consumer map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", c.ManagerUri, getConsumerByIdPath)

	headerValue := map[string]string{
		c.TokenHeaderName: c.Token,
	}
	queryValue := map[string]string{
		"name": c.MessageName,
		"consumer_id": consumerId,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return consumer, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return consumer, errors.New(fmt.Sprintf(v["consumer"].(string)))
	}

	return v["data"].(map[string]interface{}), nil
}

