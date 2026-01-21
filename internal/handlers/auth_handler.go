func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string
		Password string
		OTP string
	}
	json.NewDecoder(r.Body).Decode(&req)

	ar adminID, hash, secret string
	err := db.QueryRow(`
		SELECT id, password_hash, totp_secret
		FROM admins WHERE email=$1
	`, req.Email).Scan(&adminID, &hash, &secret)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(req.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := auth.GenerateJWT(adminID)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}