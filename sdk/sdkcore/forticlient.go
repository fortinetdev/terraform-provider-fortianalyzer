package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/terraform-providers/terraform-provider-fortianalyzer/sdk/auth"
	"github.com/terraform-providers/terraform-provider-fortianalyzer/sdk/config"
	"github.com/terraform-providers/terraform-provider-fortianalyzer/sdk/request"
)

// FortiSDKClient describes the global FAZ plugin client instance
type FortiSDKClient struct {
	Config  config.Config
	Session string
	Token   string
	Retries int
}

func escapeURLString(v string) string { // doesn't support "<>()"'#"
	return v
	// return strings.Replace(url.QueryEscape(v), "+", "%20", -1)
}

// NewClient initializes a new global plugin client
// It returns the created client object
func NewClient(auth *auth.Auth, client *http.Client) (c *FortiSDKClient, err error) {
	c = &FortiSDKClient{}

	c.Session = ""

	c.Config.Auth = auth
	c.Config.HTTPCon = client
	c.Config.FwTarget = auth.Hostname
	if auth.Token != "" {
		c.Token = auth.Token
	}

	// Only login here if not specify cleansession.
	if !auth.CleanSession {
		err = c.login()
	}

	return
}

func (c *FortiSDKClient) login() (err error) {
	auth := c.Config.Auth

	if auth.Session != "" {
		c.Session = auth.Session
	} else {
		c.Session, err = c.loginSession()
	}

	if err != nil {
		return
	}

	if auth.LogSession == true && c.Session != "" {
		saveSession(c.Session)
	}
	return
}

func saveSession(session string) {
	f, err := os.Create("presession.txt")
	if err != nil {
		return
	}
	_, err = f.WriteString(session + "\n")
	if err != nil {
		f.Close()
		return
	}

	err = f.Close()
	if err != nil {
		return
	}
}

func (c *FortiSDKClient) loginSession() (session string, err error) {
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
	// log.Printf("FAZ reading login response: %s", string(body))
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
						session = fmt.Sprintf("%v", result["session"])
					}
				}
			}
		}
	} else {
		err = fmt.Errorf("Response body is nil or do not contain result.")
	}
	return
}

func (c *FortiSDKClient) logoutSession(session string) (err error) {
	if session == "" {
		return
	}

	data := make(map[string]interface{})
	data["method"] = "exec"
	data["params"] = make([]map[string]interface{}, 0)
	data["session"] = session

	paramItem := make(map[string]interface{})
	paramItem["url"] = "/sys/logout"
	params := make([]map[string]interface{}, 0)
	params = append(params, paramItem)
	data["params"] = params

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
					if tmp == 0 || tmp == -11 {
						c.Session = ""
					} else {
						err = fmt.Errorf("Response code could not be recognized: %v", tmp)
					}
				}
			}
		}
	} else {
		err = fmt.Errorf("Response body is nil or do not contain result.")
	}
	return
}

// NewRequest creates the request to FAZ for the client
// and return it to the client
func (c *FortiSDKClient) NewRequest(method string, path string, params interface{}, data *bytes.Buffer) *request.Request {
	return request.New(c.Config, method, path, params, data)
}
