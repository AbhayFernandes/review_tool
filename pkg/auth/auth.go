package auth

func ValidateToken(token string) bool {
    return token == "valid-token"
}

