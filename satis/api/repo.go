package api

import (
	"hash/crc32"
)

type Repo struct {
	Id   uint32 `json:"id"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

func NewRepo(t string, u string) *Repo {
	h := crc32.NewIEEE()
	h.Write([]byte(u))
	return &Repo{
		Id:   h.Sum32(),
		Type: t,
		Url:  u,
	}
}
