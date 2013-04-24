package main

import (
	"strconv"
)

type Candidate struct {
	mir_id        int
	party_id      int
	candidate_num int
	name          string
}

type Candidates []Candidate

func NewCandidates() Candidates {
	return make(Candidates, 0, 0)
}

func (c Candidates) Add(record []string) {
	mir_id, _ := strconv.Atoi(record[0])
	party_id, _ := strconv.Atoi(record[1])
	candidate_num, _ := strconv.Atoi(record[2])
	c = append(c, Candidate{mir_id, party_id, candidate_num, record[3]})
}

func (c Candidates) Load(path string) {
	file_name := path + CANDIDATE_FILENAME
	loadFile(c, file_name)
}