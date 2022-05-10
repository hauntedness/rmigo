package pkg

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

var (
	ErrInvalidCookie       error = errors.New("cookie is expired")
	ErrInvalidUserPassword error = errors.New("user password is wrong")
)

type MockData struct {
	projectList map[string]Project
	sprintList  map[string]Sprint
	storyList   map[string]Story
	defectList  map[string]Defect
}

func Mock() MockData {
	var mockData MockData

	var p1 = Node{
		ID:     "p1",
		Type:   "project",
		Name:   "p1",
		Parent: root,
	}
	var p2 = Node{
		ID:     "p2",
		Type:   "project",
		Name:   "p2",
		Parent: root,
	}
	var i11 = Node{
		ID:     "sprint11",
		Type:   "sprint",
		Name:   "sprint11",
		Parent: &p1,
	}
	var i12 = Node{
		ID:     "sprint12",
		Type:   "sprint",
		Name:   "sprint12",
		Parent: &p1,
	}
	var i21 = Node{
		ID:     "sprint21",
		Type:   "sprint",
		Name:   "sprint21",
		Parent: &p1,
	}

	var s1 = Story{
		Item: Item{
			ID:     "story1",
			Desc:   "sssdfsadfdasfdsafdsfdsafdsa",
			Parent: nil,
			Dir:    &i11,
		},
		Acceptance: "",
		Attachment: "",
		Owner:      "user1",
		Status:     "completed",
	}
	var s2 = Story{
		Item: Item{
			ID:     "story2",
			Desc:   "sssdfsadfdasfdsafdsfdsafdsa",
			Parent: nil,
			Dir:    &i11,
		},
		Acceptance: "",
		Attachment: "",
		Owner:      "user1",
		Status:     "completed",
	}
	var s3 = Story{
		Item: Item{
			ID:     "story3",
			Desc:   "sssdfsadfdasfdsafdsfdsafdsa",
			Parent: nil,
			Dir:    &i12,
		},
		Acceptance: "",
		Attachment: "",
		Owner:      "user1",
		Status:     "completed",
	}
	var s4 = Story{
		Item: Item{
			ID:     "story4",
			Desc:   "sssdfsadfdasfdsafdsfdsafdsa",
			Parent: nil,
			Dir:    &i21,
		},
		Acceptance: "",
		Attachment: "",
		Owner:      "user1",
		Status:     "completed",
	}
	var d1 = Defect{
		Item: Item{
			ID:     "defect1",
			Desc:   "sssdfsadfdasfdsafdsfdsafdsa",
			Parent: &s1.Item,
			Dir:    &i11,
		},
		Acceptance: "",
		Attachment: "",
		Owner:      "user1",
		Status:     "completed",
	}
	var d2 = Defect{
		Item: Item{
			ID:     "defect2",
			Desc:   "sssdfsadfdasfdsafdsfdsafdsa",
			Parent: &s1.Item,
			Dir:    &i11,
		},
		Acceptance: "",
		Attachment: "",
		Owner:      "user1",
		Status:     "completed",
	}
	var d3 = Defect{
		Item: Item{
			ID:     "defect3",
			Desc:   "sssdfsadfdasfdsafdsfdsafdsa",
			Parent: &s2.Item,
			Dir:    &i12,
		},
		Acceptance: "",
		Attachment: "",
		Owner:      "user1",
		Status:     "completed",
	}
	var d4 = Defect{
		Item: Item{
			ID:     "defect4",
			Desc:   "sssdfsadfdasfdsafdsfdsafdsa",
			Parent: &s4.Item,
			Dir:    &i21,
		},
		Acceptance: "",
		Attachment: "",
		Owner:      "user1",
		Status:     "completed",
	}
	mockData.sprintList = map[string]Sprint{
		i11.ID: {&i11}, i12.ID: {&i12}, i21.ID: {&i21},
	}
	mockData.projectList = map[string]Project{
		p1.ID: {&p1}, p2.ID: {&p2},
	}
	mockData.storyList = map[string]Story{
		s1.ID: s1, s2.ID: s2, s3.ID: s3, s4.ID: s4,
	}
	mockData.defectList = map[string]Defect{
		d1.ID: d1, d2.ID: d2, d3.ID: d3, d4.ID: d4,
	}
	return mockData
}

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
	return nil
}

//TODO
func (c *RTCClient) ListProjects() (list []Project, err error) {
	for _, p := range Mock().projectList {
		list = append(list, p)
	}
	return
}

func (c *RTCClient) GetProject(projectId string) (project Project, err error) {
	if v, ok := Mock().projectList[projectId]; ok {
		project = v
	} else {
		err = errors.New("project not found")
	}
	return
}

//TODO
func (c *RTCClient) ListSprints(projectId string) (list []Sprint, err error) {
	md := Mock()
	p := md.projectList[projectId]
	for _, s := range md.sprintList {
		if s.Parent.ID == p.ID {
			list = append(list, s)
		}
	}
	return
}

func (c *RTCClient) GetSprint(id string) (sprint Sprint, err error) {
	if v, ok := Mock().sprintList[id]; ok {
		sprint = v
	} else {
		err = errors.New("sprint not found")
	}
	return
}

//TODO
func (c *RTCClient) ListStoryOfProject(id string) (list []*Story) {
	for _, s2 := range Mock().storyList {
		if s2.Dir == nil {
			continue
		} else if s2.Dir.Parent == nil {
			continue
		} else if s2.Dir.Parent.ID == id {
			list = append(list, &s2)
		}
	}
	return
}

//TODO
func (c *RTCClient) ListStoryOfSprint(id string) (list []*Story) {
	for _, s2 := range Mock().storyList {
		if s2.Dir == nil {
			continue
		} else if s2.Dir.ID == id {
			list = append(list, &s2)
		}
	}
	return
}

//TODO
func (c *RTCClient) GetStoryDetail(id string) (s *Story, err error) {
	if s2, ok := Mock().storyList[id]; ok {
		s = &s2
	} else {
		err = errors.New("story not found")
	}
	return
}

//TODO
func (c *RTCClient) GetDefectDetail(id string) (defect *Defect, err error) {
	if v, ok := Mock().defectList[id]; ok {
		defect = &v
	} else {
		err = errors.New("defect not found")
	}
	return
}

//TODO
func (c *RTCClient) SetStoryDetail(storyId string) error {
	return nil
}

//TODO
func (c *RTCClient) setDefectDetail(defectId string) error {
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
