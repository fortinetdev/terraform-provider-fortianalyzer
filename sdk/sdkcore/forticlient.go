package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/terraform-providers/terraform-provider-fortianalyzer/sdk/auth"
	"github.com/terraform-providers/terraform-provider-fortianalyzer/sdk/config"
	"github.com/terraform-providers/terraform-provider-fortianalyzer/sdk/request"
)

// FortiSDKClient describes the global FAZ plugin client instance
type FortiSDKClient struct {
	Config  config.Config
	Session string
	Retries int
}

func escapeURLString(v string) string { // doesn't support "<>()"'#"
	return v
	// return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) *FortiSDKClient {
	c := &FortiSDKClient{}

	c.Session = ""

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = auth.Hostname

	if auth.Session != "" {
		c.Session = auth.Session
	} else {
		login(auth, c)
	}

	return c
}

func login(auth *auth.Auth, c *FortiSDKClient) {
	data := make(map[string]interface{})
	data["method"] = "exec"
	data["params"] = make([]map[string]interface{}, 0)

	paramItem := make(map[string]interface{})
	paramItem["url"] = "sys/login/user"
	v1 := make([]map[string]interface{}, 0)

	paramItemData := make(map[string]interface{})
	paramItemData["user"] = c.Config.Auth.User
	paramItemData["passwd"] = c.Config.Auth.Passwd
	v1 = append(v1, paramItemData)
	paramItem["data"] = v1

	v2 := make([]map[string]interface{}, 0)
	v2 = append(v2, paramItem)
	data["params"] = v2

	locJSON, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)

	req := c.NewRequest("POST", "/jsonrpc", nil, bytes)
	err = req.Send()
	if err != nil || req.HTTPResponse == nil {
		err = fmt.Errorf("cannot send request %v", err)
		return
	}

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	req.HTTPResponse.Body.Close()
	if err != nil || body == nil {
		err = fmt.Errorf("cannot get response body %v", err)
		return
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	if result != nil && result["result"] != nil {
		v := result["result"].([]interface{})
		if v[0] != nil {
			v1 := v[0].(map[string]interface{})
			if v1["status"] != nil {
				v2 := v1["status"].(map[string]interface{})

				if vv, ok := v2["code"].(float64); ok {
					tmp := int(vv)
					if tmp == 0 {
						c.Session = fmt.Sprintf("%v", result["session"])
					}
				}

			}
		}
	}

	return
}

// NewRequest creates the request to FAZ for the client
// and return it to the client
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data)
}
