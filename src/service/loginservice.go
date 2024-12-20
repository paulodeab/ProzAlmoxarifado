package service

func Login(login string, password string) bool {
	if login == "admin" && password == "1234" {
		return true
	}
	return false
}
