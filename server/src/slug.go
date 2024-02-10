package main

type slug struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Domain string `json:"redirect"`
	UserID string `json:"uid"`
}
