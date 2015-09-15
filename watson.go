package watson

import (
	"net/http"
	"errors"

	"bytes"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Watson struct {
	user string
	pass string
}

type Error struct {
	Code int			`json:"code"`
	Error string			`json:"error"`
}

const (
	watsonUserAgent = "Golang Watson SDK v0.1"
	watsonUrl  = "stream.watsonplatform.net"
	httpsScheme = "https://"
)

func New(user, pass string) *Watson {
	return &Watson{
		user: user,
		pass: pass,
	}
}

func watsonError(res *http.Response) error {
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	
	e := &Error{}
	dec := json.NewDecoder(bytes.NewReader(b))

	err = dec.Decode(e)
	if err != nil {
		if res.StatusCode >= 500 {
			return errors.New("Watson server error")
		}
		return errors.New(fmt.Sprintf("Watson http error %d\n%s", 
			res.StatusCode, string(b)))
	}
	
	return errors.New("Watson says '" + e.Error + "'")
}
