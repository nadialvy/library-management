package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// middleware digunakan untuk memvalidasi jwt dan memastikan pengguna itu valid untuk
// melakukan suatu request ke suatu endpoint
// middleware berjalan diantara request dan response
// req -> middleware -> resp
// middleware dipanggil setiap kali ada req ke route terkait

// fungsi utama yang ngembaliin gin.HandlerFunc
// middleware nya dibungkus gungsi biar bisa dipake di berbagai route
// Ini adalah bentuk dari higher order function
func AuthMiddleware() gin.HandlerFunc {

	// fungsi ini yang bakalan dieksekusi tiap kali middleware dipanggil
	// *gin.Context adalah objek untuk menyimpan semua data yang berkaitan dengan HTTP request dan response
	// dengan gin.Context kita bisa ngambil data dari request = c.GetHeader("Authorization") dll

	return func(c *gin.Context) {
		// ambil token dari header auth
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// token harus dimulai dengan "Bearer "
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid format token"})
			c.Abort()
			return
		}

		// token validation
		// jwt parse digunakan untuk ngeparsing token jwt dan ngeverif apakah signature nya valid
		// parameternya ada 2
		// 1. tokenString = token jwt dari client
		// 2. callback function = fungsi yang ngasi tau jwt librari giamna cara dapetin jwtsecret untuk verif token
		// kenapa callback? memberikan fleksibilitas, kalau pengen pake public key dalam algoritma seperti rs256
		// hasil dari jwt.Parse adalah token atau error
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		// token yang dihasilkan disini adalah objeck dari tipe *jwt.Token (struct dari library jwt)
		// strukturnya kurang lebih
		// type Token struct {
		// Raw       string      // Token JWT asli (dalam bentuk string)
		// Method    SigningMethod // Algoritma signing yang digunakan
		// Header    map[string]interface{} // Header dari token JWT
		// Claims    Claims      // Payload (data di dalam token)
		// Signature string      // Signature token
		// Valid     bool        // Apakah token valid atau tidak
		// }

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// simpan klaim token di context. Claim itu isinya userid, role dan exp
		// claims adalah variable lokal untuk menyimpan payload token kalau type assertion berhasil.
		// IYA..claims tuh di declare didalam if.... ANEH BANGET FUCK
		// ada 2 statement di sini
		// 1. claims, ok := token.Claims.(jwt.MapClaims)
		// 2. ok

		// if variable, kondisi := expression; kondisi
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			if time.Now().Unix() > exp {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
				c.Abort()
				return
			}
			c.Set("userID", claims["userID"])
			c.Set("role", claims["role"])
		}

		// penjelasan if diatas
		// misal kita punya token payloadnya
		// { "userID": 123, "role": "admin", "exp": 170000 }
		// nah tipe assertion diatas memastikan bahwa payload dari token (token.Claims)
		c.Next()
	}
}
