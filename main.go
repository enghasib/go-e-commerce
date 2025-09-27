package main

import (
	"github.com/enghasib/server/internal/cmd"
)

// "fmt"

// "github.com/enghasib/server/internal/utils"

// type People interface {
// 	GetDetails()
// }

// type BankUser interface {
// 	ReceiveMoney(money float64) float64
// }

// type user struct {
// 	name        string
// 	age         int
// 	nationality string
// 	money       float64
// }

// func (usr user) GetDetails() {
// 	fmt.Printf("My name is %s i'm %d years old. I'm a %s \n", usr.name, usr.age, usr.nationality)
// 	fmt.Println("Money:", usr.money)
// }

// func (user *user) ReceiveMoney(money float64) float64 {
// 	user.money = user.money + money
// 	return user.money
// }

func main() {

	cmd.Serve()

	// var usr1 People
	// usr1 = user{
	// 	name:        "al hasib",
	// 	age:         24,
	// 	nationality: "Bangladeshi",
	// 	money:       53.34,
	// }

	// usr1.GetDetails()

	// var usr2 BankUser
	// usr2 = &user{
	// 	name:        "shakil",
	// 	age:         23,
	// 	nationality: "bdshi",
	// 	money:       50.55,
	// }

	// usr2.ReceiveMoney(511)

	// usr3 := user{
	// 	name:        "mamun",
	// 	age:         25,
	// 	nationality: "bangladeshi",
	// 	money:       355.994,
	// }

	// usr3.ReceiveMoney(45.6)

	// fmt.Println("money:", usr3.money)

	// sha algorithm

	// payload, err := utils.DecodeToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOjEsInVzZXJfbmFtZSI6InNoYWtpbCIsImVtYWlsIjoic2hha2lsQGdtYWlsLmNvbSIsImlzX3Nob3Bfb3duZXIiOnRydWV9.1ovBg-CTGZZecbQuIIwx2xlHLmgByt7TrALFqhLfgSk", config.GetEnv("JWT_SECRET", "jwt"))

	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// fmt.Println("payload:", *payload)

	// a := "hello world "
	// byteArr := []byte(a)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// bs64str := enc.EncodeToString(byteArr)

	// fmt.Println("Encoded", bs64str)

	// decodedStr, err := enc.DecodeString(bs64str)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// fmt.Println("Decoded str:", string(decodedStr))

	//sha256
	// inp := []byte("hello")

	// sha := sha256.Sum256(inp)

	// fmt.Println("Secure Hash:", sha)

	//HMC

	// secret := []byte("my-secret")
	// message := []byte("hello world")

	// hmc := hmac.New(sha256.New, secret)

	// hash := hmc.Sum(message)
	// fmt.Println("Hash:", hash)
	// data := utils.Payload{
	// 	Sub:         "4565",
	// 	Name:        "al hasib",
	// 	Email:       "hasib@example.com",
	// 	IsShopOwner: true,
	// }

	// jwt, err := utils.CreateToken("my-secret", data)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// fmt.Println("Jwt:", jwt)

	// isVerifiedToken, err := utils.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0NTY1IiwibmFtZSI6ImFsIGhhc2liIiwiZW1haWwiOiJoYXNpYkBleGFtcGxlLmNvbSIsImlzX3Nob3Bfb3duZXIiOnRydWV9.y5d0bRP6Fjv1vOrYwzsjccqOe_b4zHZkO8fulgyhuzg", "my-secret")

	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// fmt.Println("Token verified:", isVerifiedToken)

}
