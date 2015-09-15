package watson

import (
	"io"
	"path"
	"bytes"

	"net/http"
	"crypto/tls"
	"encoding/json"
)

const (
	ttsVer = "v1"
)

type synth struct {
	Text string			`json:"text"`
}

type Voice struct {
	Name string			`json:"name"`
	Lang string			`json:"language"`
	Gender string			`json:"gender"`
}

type Voices struct {
	Voices []Voice			`json:"voices"`
}

func (w *Watson) GetVoices() (*Voices, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	
	uri := path.Join(watsonUrl, "text-to-speech", "api", ttsVer, "voices")

	req, err := http.NewRequest("GET", "https://" + uri, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(w.user, w.pass);

	req.Header.Set("User-Agent", watsonUserAgent)
	req.Header.Set("accept", `application/json`)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, watsonError(res)
	}

	dec := json.NewDecoder(res.Body)

	v := &Voices{}
	err = dec.Decode(v)
	if err != nil {
		return nil, err
	}
	
	return v, nil
}

func (w *Watson) Synthesize(text string, voice string, fmt string) (io.ReadCloser, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	
	uri := path.Join(watsonUrl, "text-to-speech", "api", ttsVer, "synthesize")

	var err error
	var bodyBuf []byte
	var req *http.Request

	body := synth{ Text: text }

	bodyBuf, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("POST", "https://" + uri + 
		"?accept=audio/" + fmt + "&voice=" + voice,
		bytes.NewBuffer(bodyBuf))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(w.user, w.pass);

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", watsonUserAgent)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		defer res.Body.Close()
		return nil, watsonError(res)
	}

	return res.Body, nil
}
