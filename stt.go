package watson

import (
	"io"
	"path"
	"bytes"
	"errors"
	"net/url"
	"net/http"
	"crypto/tls"
	"encoding/json"
)

const (
	sttVer = "v1"
)

type Model struct {
	Rate int			`json:"rate"`
	Name string			`json:"name"`
	Lang string			`json:"language"`
	Desc string			`json:"description"`
}

type Models struct {
	Models []Model			`json:"models"`
}

func (w *Watson) GetModels() (*Models, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	
	uri := path.Join(watsonUrl, "speech-to-text", "api", sttVer, "models")

	req, err := http.NewRequest("GET", httpsScheme + uri, nil)
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

	m := &Models{}
	err = dec.Decode(m)
	if err != nil {
		return nil, err
	}
	
	return m, nil
}

type Session struct {
	SessionId string		`json:"session_id"`
	NewSessionUri string		`json:"new_session_uri"`
	Recognize string		`json:"recognize"`
	RecognizeWS string		`json:"recognizeWS"`
	ObserveResult string		`json:"observe_result"`
	Cookies []*http.Cookie
}

func (w *Watson) createSession(model string) (*Session, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	data := url.Values{}
	data.Set("model", model)
	
	uri := path.Join(watsonUrl, "speech-to-text", "api", sttVer, "sessions")

	buf := bytes.NewBufferString(data.Encode())

	req, err := http.NewRequest("POST", httpsScheme + uri, buf)
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

	if res.StatusCode != 201 {
		return nil, watsonError(res)
	}

	dec := json.NewDecoder(res.Body)

	ss := &Session{}
	err = dec.Decode(ss)
	if err != nil {
		return nil, err
	}

	ss.Cookies = readSetCookies(res.Header) 	
	return ss, nil
}

type Word struct {
	Token string
	Begin float64
	End float64
	Confidence float64
}

type Text struct {
	Words []Word
}

type RHeader struct {
	Type string			`json:"part_content_type"`
	Count int			`json:"data_parts_count"`
	Continuous bool			`json:"continuous"`
	Timeout int			`json:"inactivity_timeout"`
	Alternatives int		`json:"max_alternatives"`
	Timestamps bool			`json:"timestamps"`
	Confidence bool			`json:"word_confidence"`
}

type RMeta struct {
	Data RHeader			`json:"metadata"`
	Files []string			`json:"files"`
}

type Results struct {
	Index int			`json:"result_index"`
	List []Result			`json:"results"`
}

type Result struct {
	Alternatives []Alternative	`json:"alternatives"`
	Final bool			`json:"final"`
}

type Alternative struct {
	Confidence float64		`json:"confidence"`
	Timestamps [][]interface{}	`json:"timestamps"`
	Transcript string		`json:"transcript"`
	Words [][]interface{}		`json:"word_confidence"`
}

//
// TODO: recognize requests don't work yet with sessions.
//
func (w *Watson) Recognize(is io.Reader, model, afmt string) (*Text, error) {
/*
	ss, err := w.createSession(model)
	if err != nil {
		return nil, err
	}
*/

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives: true,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	tp := ""
	switch afmt {
	case "wav":
		tp = "audio/wav"
	case "flac":
		tp = "audio/flac"
	default:
		return nil, errors.New("Invalid file format")
	}

/*
	uri := ss.NewSessionUri
*/
	uri := httpsScheme + path.Join(watsonUrl, "speech-to-text", "api", sttVer, "recognize")

	if false {
		uri = "http://philippe-anel.fr:3000"
		client = &http.Client{}
	}

	// TODO: parameters 
	req, err := http.NewRequest("POST", 
		uri +
		"?continuous=true" + 
		"&max_alternatives=1" +
		"&timestamps=true" +
		"&word_confidence=true" +
		"", is)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(w.user, w.pass);

	req.Header.Set("Content-Type", tp)
	req.Header.Set("accept", `application/json`)
	req.Header.Set("User-Agent", watsonUserAgent)

/*
	for _, c := range ss.Cookies {
		req.AddCookie(c)
	}
*/

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, watsonError(res)
	}
	
	return ParseResponse(res.Body)
}

func ParseResponse(rd io.Reader) (*Text, error) {
	dec := json.NewDecoder(rd)

	rs := &Results{}

	err := dec.Decode(rs)
	if err != nil {
		return nil, err
	}

	tt := &Text{}

	for _, r := range rs.List {
		a := r.Alternatives[0]
		for i, t := range a.Timestamps {
			w := a.Words[i]
			
			n := t[0]
			b := t[1]
			e := t[2]
			c := w[1]

			tt.Words = append(tt.Words, Word{
				Token: n.(string),
				Begin: b.(float64),
				End: e.(float64),
				Confidence: c.(float64),
			})
		}
	}

	return tt, nil
}
