package router

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"

	"github.com/MD-Levitan/mqqt-app/config"
	"github.com/MD-Levitan/mqqt-app/models"
)

func jsonResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func JSONHandler(handler func(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		encoder := json.NewEncoder(w)
		if err := handler(decoder, encoder, w, r); err != nil {
			w.WriteHeader(400)
			encoder.Encode(models.Error{err.Error()})
			return
		}
	}
}

func getUserContext(s *sessions.Session) *models.UserContext {
	val := s.Values["Context"]
	var user = &models.UserContext{}
	user, ok := val.(*models.UserContext)
	if !ok {
		logrus.Error("cannot get user context from session")
		return nil
	}
	return user
}

func authorizeByCookie(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		session, err := config.GetStore().Get(r, "Rcookie")
		if err != nil || session.IsNew {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		context := getUserContext(session)
		if context == nil {
			session.AddFlash("You don't have access!")
			err = session.Save(r, w)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authHandler(dec *json.Decoder, enc *json.Encoder, w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Printf("authHandler")
	return fmt.Errorf("test error")
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		logrus.Error(err)
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

// func createJWT(user models.User) (string, error) {
// 	conf := config.GetConfig()
// 	password, err := encrypt([]byte(conf.Web.SessionKey), []byte(user.Password))
// 	if err != nil {
// 		logrus.Error(err)
// 		return "", err
// 	}
// 	expirationTime := time.Now().Add(time.Hour)
// 	claims := &models.UserClaim{
// 		Username: user.Username,
// 		Password: base64.RawStdEncoding.EncodeToString(password),
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(conf.Web.JWTSecret))

// }

// func authorizeByJWT(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		conf := config.GetConfig()
// 		user := models.UserClaim{}

// 		tokenString := r.Header.Get("Authorization")

// 		if !strings.HasPrefix(tokenString, "Bearer") {
// 			logrus.Error("token string doesn't contain Bearer")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		} else {
// 			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
// 		}

// 		token, err := jwt.ParseWithClaims(tokenString, &user, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}

// 			return []byte(conf.Web.JWTSecret), nil
// 		})

// 		if err != nil || !token.Valid {
// 			logrus.Error(err)
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
