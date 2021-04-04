package authentication

type RegistrationBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
}
