package custom_bcrypt

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func CreateBcryptPassword() []byte {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envData, err := godotenv.Read(".env")
	password := []byte(envData["BCRYPT_PASSWORD"])

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return hashedPassword
}

func CompareBcryptPassword(password string) bool {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envData, err := godotenv.Read(".env")
	byte_password := []byte(password)
	hash_byte_password := []byte(envData["BCRYPT_PASSWORD"])
	hashedPassword, err := bcrypt.GenerateFromPassword(hash_byte_password, bcrypt.DefaultCost)

	err = bcrypt.CompareHashAndPassword(hashedPassword, byte_password)
	fmt.Println(string(hashedPassword))
	fmt.Println(err)
	if err != nil {
		return false
	}

	return true

}
