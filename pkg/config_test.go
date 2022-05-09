package pkg

import "testing"

func TestInitConfig(t *testing.T) {
	Conf.CurrentDir = &Node{
		ID:     "test_project",
		Type:   "project",
		Name:   "what the hell",
		Parent: root,
	}
	InitConfig("user1", "pwd1")
	t.Log(Conf.Spell)
}

func TestLoadConfig(t *testing.T) {
	LoadConfig()
}
