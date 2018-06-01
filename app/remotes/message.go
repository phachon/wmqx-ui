package remotes

import (
	"fmt"
	"wmqx-ui/app/utils"
	"encoding/json"
	"errors"
	"strings"
)

var (
	messageAddPath = "/message/add"
	messageUpdatePath = "/message/update"
	messageDeletePath = "/message/delete"
	messageStatusPath = "/message/status"
	messageListPath = "/message/list"
	getMessageByNamePath = "/message/getMessageByName"
	getConsumersByNamePath = "/message/getConsumersByName"
	messageReload = "/message/reload"
)

func NewMessageByNode(node map[string]string) *Message {
	return NewMessage(node["manager_uri"], node["token_header_name"], node["token"], node["publish_uri"])
}

func NewMessage(managerUri string, tokenHeader string, token string, publishUri string) *Message {
	return &Message{
		ManagerUri: managerUri,
		PublishUri: publishUri,
		TokenHeaderName: tokenHeader,
		Token: token,
	}
}

type Message struct {
	ManagerUri string
	PublishUri string
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

func (m *Message) ReloadMessage(name string) (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageReload)

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

func (m *Message) GetConsumersStatus(name string) (consumerStatus []map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, messageStatusPath)

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
		return consumerStatus, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return consumerStatus, errors.New(fmt.Sprintf(v["message"].(string)))
	}
	for _, items := range v["data"].([]interface{}) {
		if items == nil {
			continue
		}
		consumerStatus = append(consumerStatus, items.(map[string]interface{}))
	}
	return
}

func (m *Message) Publish(name string, method string, data string, routeKey string) (err error) {

	message, err := m.GetMessageByName(name)
	if err != nil {
		return
	}
	headerValue := map[string]string{}
	if message["is_need_token"].(bool) {
		headerValue["WMQX_MESSAGE_TOKEN"] = message["token"].(string)
	}
	if routeKey != "" {
		headerValue["WMQX_MESSAGE_ROUTEKEY"] = routeKey
	}
	url := fmt.Sprintf("%s/publish/%s", m.PublishUri, name)

	queryValue := utils.Request.ParseString(data)
	code := 0
	if strings.ToLower(method) == "get" {
		_, code, err = utils.Request.HttpGet(url, queryValue, headerValue)
	}else {
		_, code, err = utils.Request.HttpPost(url, queryValue, headerValue)
	}
	if err != nil {
		return
	}
	if code != 200 {
		return errors.New(fmt.Sprintf("request status: %d", code))
	}
	return nil
}