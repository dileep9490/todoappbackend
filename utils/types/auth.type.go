package types

//TODO: add validators
type LoginType struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type SignUpType struct {
	LoginType
	Name string `json:"name"`
}
