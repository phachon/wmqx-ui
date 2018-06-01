package remotes

import (
	"fmt"
	"wmqx-ui/app/utils"
	"encoding/json"
	"errors"
)

var (
	logSearchPath = "/log/search"
	logListPath = "/log/list"
)

func NewLogByNode(node map[string]string) *Log {
	return NewLog(node["manager_uri"], node["token_header_name"], node["token"])
}

func NewLog(managerUri string, tokenHeader string, token string) *Log {
	return &Log{
		ManagerUri: managerUri,
		TokenHeaderName: tokenHeader,
		Token: token,
	}
}

type Log struct {
	ManagerUri string
	TokenHeaderName string
	Token string
}

func (m *Log) Search(number string, level string, keyword string) (logs []map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, logSearchPath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	queryValue := map[string]string{
		"number": number,
		"level": level,
		"keyword": keyword,
	}

	body, code, err := utils.Request.HttpGet(url, queryValue, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return logs, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return logs, errors.New(fmt.Sprintf(v["log"].(string)))
	}

	for _, items := range v["data"].([]interface{}) {
		if items == nil {
			continue
		}
		v1 := map[string]interface{}{}
		if json.Unmarshal([]byte(items.(string)), &v1) != nil {
			continue
		}
		logs = append(logs, v1)
	}

	return logs,nil
}

func (m *Log) List() (logs []string, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, logListPath)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	body, code, err := utils.Request.HttpGet(url, nil, headerValue)
	if err != nil {
		return
	}
	if len(body) == 0 {
		return logs, errors.New(fmt.Sprintf("request wmqx failed, httpStatus: %d", code))
	}
	v := map[string]interface{}{}
	if json.Unmarshal(body, &v) != nil {
		return
	}
	if v["code"].(float64) == 0 {
		return logs, errors.New(fmt.Sprintf(v["log"].(string)))
	}

	for _, items := range v["data"].([]interface{}) {
		if items == nil {
			continue
		}
		logs = append(logs, items.(string))
	}

	return logs,nil
}