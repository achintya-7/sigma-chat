package models

type Send struct {
	Room string `json:"room"`
	Msg  string `json:"msg"`
}