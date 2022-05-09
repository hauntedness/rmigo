package pkg

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

var (
	ErrInvalidCookie       error = errors.New("cookie is expired")
	ErrInvalidUserPassword error = errors.New("user password is wrong")
)

type RTCClient struct {
	http.Client
	header  map[string]string
	cookies map[string]*http.Cookie
}

func NewRTCClient() *RTCClient {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	return &RTCClient{
		Client: http.Client{
			Jar: jar,
		},
		cookies: Conf.Cookies,
		header:  make(map[string]string),
	}
}

//TODO
func (c *RTCClient) Login(user, password string) error {
	//var j_userid, j_password string
	u := &url.URL{}
	u.Query().Add("j_userid", user)
	u.Query().Add("j_password", password)
	res := c.Request(http.MethodPost, "", strings.NewReader(u.Query().Encode()))
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusFound:
		return nil
	default:
		b, _ := ioutil.ReadAll(res.Body)
		log.Println(string(b))
		return errors.New("login failed")
	}
}

//TODO
func (c *RTCClient) ListProjects() error {

	return nil
}

//TODO
func (c *RTCClient) ListSprints(project_id string) error {
	return nil
}

//TODO
func (c *RTCClient) GetStoryDetail(sprint_id string) error {
	return nil
}

//TODO
func (c *RTCClient) GetDefectDetail(sprint_id string) error {
	return nil
}

//TODO
func (c *RTCClient) SetStoryDetail(story_id string) error {
	return nil
}

//TODO
func (c *RTCClient) setDefectDetail(defect_id string) error {
	return nil
}

// todo
func (c *RTCClient) Request(method, url string, body io.Reader) *http.Response {
	var err error
	var req *http.Request
	var res *http.Response
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	for _, v := range c.cookies {
		req.AddCookie(v)
	}
	res, err = c.Do(req)
	if err != nil {
		panic(err)
	} else {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(b)[0:100])
	}
	for _, v := range res.Cookies() {
		Conf.Cookies[v.Name] = v
	}
	CreateConfig()
	return res
}
