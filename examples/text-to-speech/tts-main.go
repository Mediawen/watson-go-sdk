package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"

	"../../"	// "github.com/mediawen/watson-go-sdk"
)

type Cfg struct {
	User string 			`json:"user"`
	Pass string 			`json:"pass"`
}

func loadCfg(name string) (*Cfg, error) {
	f, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	cfg := &Cfg{}

	err = json.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf(
			"usage: tts out.[wav,flac] voice text\n" +
			"       tts -l\n")
		return
	}

	cfg, err := loadCfg("./tts.cfg.json")
	if err != nil {
		log.Fatal(err)
	}

	w := watson.New(cfg.User, cfg.Pass)

	vl, err := w.GetVoices()
	if err != nil {
		log.Fatal(err)
	}

	if args[1] == "-l" {
		if len(args) > 2 {
			log.Fatal("too many args")
		}

		for _, v := range vl.Voices {
			fmt.Printf("%s %-8s => %s\n", v.Lang, v.Gender, v.Name)
		}

		return
	}

	out := args[1]
	ext := ""
	switch path.Ext(out) {
	case ".wav":
		ext = "wav"
	case ".flac":
		ext = "flac"
	default:
		log.Fatal("unknown file format")
	}

	voice := args[2]
	text  := args[3]

	found := false
	for _, v := range vl.Voices {
		if v.Name == voice {
			found = true
		}
	}
	if !found {
		log.Fatal("voice not found")
	}

	a, err := w.Synthesize(text, voice, ext)
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	f, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	
	n, err := io.Copy(f, a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s: %d bytes written\n", out, n)
}
