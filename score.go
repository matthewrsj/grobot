package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var scores sync.Map
var fMux sync.Mutex

const scoreFile = "/home/mrsj/.score" // nick appended so we can have multiple instances

func updateScore(nick string, change int) {
	var current int
	if val, ok := scores.Load(nick); !ok {
		current = 0
	} else {
		current = val.(int)
	}

	scores.Store(nick, current+change)
	fMux.Lock()
	err := writeMapToScoreFile()
	fMux.Unlock()
	if err != nil {
		log.Println(err)
	}
}

func getScore(nick string) int {
	val, ok := scores.Load(nick)
	if !ok {
		val, ok = setScoreFromFile(nick)
		if !ok {
			return 0
		}
	}
	return val.(int)
}

func initScoresFromFile() {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s.%s", scoreFile, *nick))
	if err != nil {
		return
	}
	r := csv.NewReader(strings.NewReader(string(b)))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return
		}
		if len(record) != 2 {
			log.Println("record should be two elements long")
			return
		}
		i, err := strconv.Atoi(record[1])
		if err != nil {
			log.Println("malformed score")
			return
		}
		scores.Store(record[0], i)
	}
}

func setScoreFromFile(n string) (int, bool) {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s.%s", scoreFile, *nick))
	if err != nil {
		return 0, false
	}

	r := csv.NewReader(strings.NewReader(string(b)))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return 0, false
		}
		if len(record) != 2 {
			log.Println("record should be two elements long")
			return 0, false
		}
		if record[0] == n {
			i, err := strconv.Atoi(record[1])
			if err != nil {
				log.Println("malformed score")
				return 0, false
			}
			scores.Store(n, i)
			return i, true
		}
	}
	return 0, false
}

func writeMapToScoreFile() error {
	f, err := os.Create(fmt.Sprintf("%s.%s", scoreFile, *nick))
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	scores.Range(func(k, v interface{}) bool {
		fmt.Fprintf(f, "%s,%d\n", k.(string), v.(int))
		return true
	})
	return nil
}

type scoreRecord struct {
	nick  string
	score int
}

type srs []scoreRecord

func allScores() srs {
	var sr srs
	scores.Range(func(k, v interface{}) bool {
		sr = append(sr, scoreRecord{k.(string), v.(int)})
		return true
	})
	return sr
}

func topScoresString(top int) string {
	sr := allScores()
	sort.Slice(sr, func(i, j int) bool { return sr[i].score > sr[j].score })
	return sr[:top].String()
}

func (sr srs) String() string {
	sort.Slice(sr, func(i, j int) bool { return sr[i].score > sr[j].score })
	var s string
	for i := range sr {
		s += fmt.Sprintf("[%s: %d]  ", sr[i].nick, sr[i].score)
	}
	return s
}

func allScoresString() string {
	sr := allScores()
	return sr.String()
}
