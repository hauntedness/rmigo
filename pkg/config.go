package pkg

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

type Config struct {
	BaseUrl     string                  `json:"base_url"`          // rtc base url
	Username    string                  `json:"username"`          // username
	Cookies     map[string]*http.Cookie `json:"cookies,omitempty"` // store password, expiration date
	Spell       string                  `json:"spell"`             // words encrypted by runes, used to recover password
	CurrentDir  *Node                   `json:"current_dir,omitempty"`
	CurrentFile string                  `json:"current_file"`
	runes       string                  `` // secret
	password    string                  `` //
	gcm         cipher.AEAD
}

var Conf *Config
var AppName = "rmigo"
var conf_path string

var root = &Node{
	ID:     "root",
	Type:   "void",
	Name:   "~",
	Parent: nil,
}

func init() {
	configdir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(path.Join(configdir, AppName), 0666)
	if err != nil {
		panic(err)
	}
	Conf = &Config{}
	conf_path = path.Join(configdir, AppName, "config.json")
	Conf.runes = time.Now().Format("200601") + "-3.1415926"
	Conf.Cookies = make(map[string]*http.Cookie)
	Conf.CurrentDir = root
	cphr, err := aes.NewCipher([]byte(Conf.runes))
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(cphr)
	if err != nil {
		panic(err)
	}
	Conf.gcm = gcm
}

func InitConfig(username, password string) {
	Conf.Username = username
	Conf.password = password
	Conf.runes = time.Now().Format("200601") + "-3.1415926"

	nonce := make([]byte, Conf.gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	b2 := Conf.gcm.Seal(nonce, nonce, []byte(Conf.password), nil)
	base64dst := base64.RawStdEncoding.EncodeToString(b2)
	Conf.Spell = string(base64dst)
	CreateConfig()
}

func LoadConfig() {
	file, err := os.OpenFile(conf_path, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		log.Println(errors.New("can't not open config file, run rtc.exe init"))
		panic(err)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &Conf)
	if err != nil {
		panic(err)
	}

	b2, err := base64.RawStdEncoding.DecodeString(Conf.Spell)
	if err != nil {
		panic(err)
	}

	nonce, encrypted := b2[:Conf.gcm.NonceSize()], b2[Conf.gcm.NonceSize():]
	b3, err := Conf.gcm.Open(nil, nonce, encrypted, nil)
	if err != nil {
		panic(err)
	}
	Conf.password = string(b3)
}

func CreateConfig() {
	file, err := os.Create(conf_path)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(Conf, "", "\t")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(b)
	if err != nil {
		panic(err)
	}
}
