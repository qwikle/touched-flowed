package entities

type Token struct {
	Id    uint64 `json:"id"`
	Token string `json:"token"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func TokenFromJson(data interface{}) *Token {
	m := data.(map[string]interface{})
	return &Token{
		Id: uint64(m["id"].(float64)),
	}
}
