package remotes

import (
	"fmt"
	"wmqx-ui/app/utils"
	"encoding/json"
	"errors"
)

var (
	logIndex = "/log/index"
	logSearch = "/log/search"
	logDownload = "/log/download"
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

func (m *Log) Index() (logs []map[string]interface{}, err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, logIndex)

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
		v1 := map[string]interface{}{}
		if json.Unmarshal([]byte(items.(string)), &v1) != nil {
			continue
		}
		logs = append(logs, v1)
	}

	return logs,nil
}