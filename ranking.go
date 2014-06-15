package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

type Ranking struct {
	scores []int
}

func NewRanking() *Ranking {
	ranking := new(Ranking)
	ranking.scores = make([]int, 10)

	if fileExists(configFilePath()) {
		buf, err := ioutil.ReadFile(configFilePath())
		if err != nil {
			log.Fatal(err)
		}

		scoreTexts := strings.Split(string(buf), ",")
		for idx, text := range scoreTexts {
			num, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			ranking.scores[idx] = num
		}
	} else {
		for i := 0; i < 10; i++ {
			ranking.scores[i] = 0
		}
	}
	return ranking
}

func (r *Ranking) save() {
	texts := []string{}
	for _, sc := range r.scores {
		texts = append(texts, strconv.Itoa(sc))
	}
	config := strings.Join(texts, ",")
	ioutil.WriteFile(configFilePath(), []byte(config), 0644)
}

func (r *Ranking) insertScore(sc int) {
	for idx, rsc := range r.scores {
		if rsc < sc {
			r.slideScores(idx)
			r.scores[idx] = sc
			return
		}
	}
}

func (r *Ranking) slideScores(index int) {
	for i := len(r.scores) - 1; i > index; i-- {
		r.scores[i] = r.scores[i-1]
	}
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func configFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, ".tetris")
}
