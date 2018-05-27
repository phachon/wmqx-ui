package remotes

import (
	"fmt"
	"wmqx-ui/app/utils"
	"encoding/json"
	"errors"
)

var (
	messageAddPath = "/message/add"
	messageUpdatePath = "/message/update"
	messageDeletePath = "/message/delete"
	messageStatusPath = "/message/status"
	messageListPath = "/message/list"
	getMessageByNamePath = "/message/getMessageByName"
	getConsumersByNamePath = "/message/getConsumersByName"
)

func NewMessageByNode(node map[string]string) *Message {
	return NewMessage(node["manager_uri"], node["token_header_name"], node["token"])
}

func NewMessage(managerUri string, tokenHeader string, token string) *Message {
	return &Message{
		ManagerUri: managerUri,
		TokenHeaderName: tokenHeader,
		Token: token,
	}
}

type Message struct {
	ManagerUri string
	TokenHeaderName string
	Token string
}

func (m *Message) GetMessages() (messages []map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageListPath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	body, code, err := utils.Request.HttpGet(url, nil, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return messages, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return messages, errors.New(fmt.Sprintf(v["message"].(string)))
	}
	for _, items := range v["data"].([]interface{}) {
		messages = append(messages, items.(map[string]interface{}))
	}
	return messages, nil
}

func (m *Message) AddMessage(message map[string]string) (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageAddPath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	body, code, err := utils.Request.HttpPost(url, message, headerValue)
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

func (m *Message) UpdateMessage(message map[string]string) (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageUpdatePath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	body, code, err := utils.Request.HttpPost(url, message, headerValue)
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

func (m *Message) DeleteMessage(name string) (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageDeletePath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}
	queryValue := map[string]string{
		"name": name,
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
		return errors.New(fmt.Sprintf(v["message"].(string)))
	}

	return nil
}

func (m *Message) GetMessageByName(name string) (message map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, getMessageByNamePath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}
	queryValue := map[string]string{
		"name": name,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return message, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return message, errors.New(fmt.Sprintf(v["message"].(string)))
	}

	return v["data"].(map[string]interface{}), nil
}

func (m *Message) GetConsumersByName(name string) (consumers []map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, getConsumersByNamePath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}
	queryValue := map[string]string{
		"name": name,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return consumers, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return consumers, errors.New(fmt.Sprintf(v["message"].(string)))
	}
	for _, items := range v["data"].([]interface{}) {
		consumers = append(consumers, items.(map[string]interface{}))
	}

	return
}