package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"
	"github.com/gorilla/mux"
)

// Kullanıcı modeli
type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
	Token    string `json:"token"`
}

// Kullanıcılar veritabanı
var users = make(map[string]User)

// AddUser endpoint'i
func addUser(w http.ResponseWriter, r *http.Request) {
	// Giriş DTO'su
	var input User

	// JSON'dan DTO'ya map etme
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Giriş doğrulama
	if !isValidUsername(input.Username) {
		http.Error(w, "Invalid username format", http.StatusBadRequest)
		return
	}
	if !isValidPassword(input.Password) {
		http.Error(w, "Invalid password format", http.StatusBadRequest)
		return
	}

	// Parola hashleme
	hashedPassword := hashPassword(input.Password)

	// Token oluşturma
	token := generateToken(input.Username)

	// Kullanıcı oluşturma
	user := User{
		Username: input.Username,
		Password: hashedPassword,
		Token:    token,
	}

	// Kullanıcıyı veritabanına ekleme
	users[input.Username] = user

	// Çıkış DTO'su
	output := User{
		Username: input.Username,
		Token:    token,
	}

	// JSON'a map etme ve gönderme
	json.NewEncoder(w).Encode(output)
}

// RemoveUser endpoint'i
func removeUser(w http.ResponseWriter, r *http.Request) {
	// Username alma
	username := mux.Vars(r)["username"]

	// Kullanıcıyı veritabanından silme
	delete(users, username)

	// Başarılı ise 200 döndür
	w.WriteHeader(http.StatusOK)
}

// ActivateUser endpoint'i
func activateUser(w http.ResponseWriter, r *http.Request) {
	// Username alma
	username := mux.Vars(r)["username"]

	// Kullanıcıyı bulma
	user, found := users[username]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kullanıcının token'ını güncelleme
	user.Token = generateToken(username)
	users[username] = user

	// Çıkış DTO'su
	output := User{
		Username: username,
		Token:    user.Token,
	}

	// JSON'a map etme ve gönderme
	json.NewEncoder(w).Encode(output)
}

// DeactivateUser endpoint'i
func deactivateUser(w http.ResponseWriter, r *http.Request) {
	// Username alma
	username := mux.Vars(r)["username"]

	// Kullanıcıyı bulma
	user, found := users[username]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Kullanıcının token'ını kaldırma
	user.Token = ""
	users[username] = user

	// Başarılı ise 200 döndür
	w.WriteHeader(http.StatusOK)
}

// isValidUsername belirli bir kullanıcı adının geçerli olup olmadığını kontrol eder
func isValidUsername(username string) bool {
	// Kullanıcı adı için basit bir format kontrolü
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return usernameRegex.MatchString(username)
}

// isValidPassword belirli bir parolanın geçerli olup olmadığını kontrol eder
func isValidPassword(password string) bool {
	// Parola için basit bir format kontrolü
	return len(password) >= 6
}

// hashPassword parolayı hashler
func hashPassword(password string) string {
	// Burada, daha güçlü bir hashleme algoritması kullanmak daha iyi olurdu, ancak bu sadece bir örnek olduğu için MD5 kullanıyoruz
	hasher := md5.New()
	hasher.Write([]byte(password))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

// generateToken kullanıcı için bir token oluşturur
func generateToken(username string) string {
	// Token oluşturma - basitçe, kullanıcı adı ve şu anki zamanın bir kombinasyonu
	return fmt.Sprintf("%s-%d", username, time.Now().UnixNano())
}

func main() {
	// Router oluşturma
	router := mux.NewRouter()

	// Endpoitler
	router.HandleFunc("/addUser", addUser).Methods("POST")
	router.HandleFunc("/removeUser/{username}", removeUser).Methods("DELETE")
	router.HandleFunc("/activateUser/{username}", activateUser).Methods("PUT")
	router.HandleFunc("/deactivateUser/{username}", deactivateUser).Methods("PUT")

	// Sunucu başlatma
	http.ListenAndServe(":8080", router)
}
