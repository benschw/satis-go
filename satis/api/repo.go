package api

import (
	"hash/crc32"
)

type Repo struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

func NewRepo(t string, u string) *Repo {
	h := crc32.NewIEEE()
	n, _ := h.Write([]byte(u))

	return &Repo{
		Id:   n,
		Type: t,
		Url:  u,
	}
}
