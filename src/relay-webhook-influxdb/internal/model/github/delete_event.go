package model

type DeleteEvent struct {
	Ref        string     `json:"ref"`
	RefType    string     `json:"ref_type"`
	Repository Repository `json:"repository"`
	Sender     Sender     `json:"sender"`
}
