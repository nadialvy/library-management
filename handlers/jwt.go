package handlers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// []byte merupakan tipe data yang merepresentasikan array dalam bentuk byte
// ex: data := []byte{65,66,67}
// []byte(something) artinya mengconvert something ini ke dalam bentuk byte
var jwtSecret = []byte(os.Getenv("JWT_SECRET")) //convert string JWT_SECRET ke dalam bentuk byte

// strukturnya
// func namaFunc (param) (return type)
// kalau di return type nya ada >1 maka return bisa dalam 2 bentuk => go mendukung multiple return values
func GenerateToken(userID uint, role string) (string, error) {
	// payload JWT
	claims := jwt.MapClaims{
		"userID": userID,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), //24 jam exp
	}

	// bikin jwt dengan method NewWithClaims
	// ES256 itu algoritmanya, dan yang paling populer karena keamanan tinggi, ukuran signature lebih kecil, dan performa lebih baik
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// menandatangain token = memastikan tokennya valid dan ga di modif
	return token.SignedString(jwtSecret)
}
