package deezer

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	redirectURI = "http://localhost:10029"
)

var TokenPath string = ""
var Token string = ""

func init() {
}

func Login(applicationId string, applicationSecret string) error {

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	TokenPath = filepath.Join(usr.HomeDir, ".musicapp.deezer")
	state, err := generateRandomString(32)
	if err != nil {
		return err
	}

	ch := make(chan string)

	http.Handle("/", &authHandler{state: state, ch: ch, applicationId: applicationId, applicationSecret: applicationSecret})
	go http.ListenAndServe("localhost:10029", nil)

	url := fmt.Sprintf("https://connect.deezer.com/oauth/auth.php?app_id=%s&redirect_uri=%s&perms=basic_access,email", applicationId, redirectURI)
	fmt.Println("Please log in to Deezer by visiting the following page in your browser:", url)

	tok := <-ch
	Token = tok

	if err := SaveToken(tok); err != nil {
		return err
	}

	return nil
}

func SaveToken(tok string) error {
	f, err := os.OpenFile(TokenPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	return enc.Encode(tok)
}

func ReadToken() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	TokenPath = filepath.Join(usr.HomeDir, ".musicapp.deezer")
	content, err := ioutil.ReadFile(TokenPath)
	if err != nil {
		return "", err
	}

	var tok string
	if err := json.Unmarshal(content, &tok); err != nil {
		return "", err
	}

	return tok, nil
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

type authHandler struct {
	applicationId     string
	applicationSecret string

	state string
	ch    chan string
	token string
}

func (a *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if a.token != "" {
		return
	}

	code := r.FormValue("code")
	url := fmt.Sprintf("https://connect.deezer.com/oauth/access_token.php?app_id=%s&secret=%s&code=%s",
		a.applicationId, a.applicationSecret, code)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	token := ""

	segments := strings.Split(string(contents), "&")
	if len(segments) > 0 {
		values := strings.Split(segments[0], "=")
		if len(values) > 0 {
			token = values[1]
			a.token = token
		}
	}

	fmt.Println("Login successfully. Please return to your terminal.")

	a.ch <- token
}
