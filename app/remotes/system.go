package remotes

import (
	"fmt"
	"wmqx-ui/app/utils"
	"encoding/json"
	"errors"
)

var (
	systemReload = "/system/reload"
)

func NewSystemByNode(node map[string]string) *System {
	return NewSystem(node["manager_uri"], node["token_header_name"], node["token"])
}

func NewSystem(managerUri string, tokenHeader string, token string) *System {
	return &System{
		ManagerUri: managerUri,
		TokenHeaderName: tokenHeader,
		Token: token,
	}
}

type System struct {
	ManagerUri string
	TokenHeaderName string
	Token string
}

func (m *System) ReloadSystem() (err error) {

	url := fmt.Sprintf("%s%s", m.ManagerUri, systemReload)

	headerValue := map[string]string{
		m.TokenHeaderName: m.Token,
	}

	body, code, err := utils.Request.HttpGet(url, nil, headerValue)
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