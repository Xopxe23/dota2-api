package dota2

import "fmt"

type Hero struct {
	Id          int      `json:"id"`
	Name        string   `json:"localized_name"`
	PrimaryAttr string   `json:"primary_attr"`
	AttackType  string   `json:"attack_type"`
	Roles       []string `json:"roles"`
}

type heroesResponce struct {
	heroes []Hero
}

func (h Hero) Info() string {
	return fmt.Sprintf("[ID]: %d | [Name]: %s | [Primary Attr]: %s | [Attack Type]: %s | [roles]: %s",
		h.Id, h.Name, h.PrimaryAttr, h.AttackType, h.Roles)
}
