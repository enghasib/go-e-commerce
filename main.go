package main

import (
	"fmt"

	"github.com/enghasib/server/internal/utils"
)

// "github.com/enghasib/server/internal/cmd"

func main() {
	// cmd.Serve()

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

	isVerifiedToken, err := utils.Verify("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI0NTY1IiwibmFtZSI6ImFsIGhhc2liIiwiZW1haWwiOiJoYXNpYkBleGFtcGxlLmNvbSIsImlzX3Nob3Bfb3duZXIiOnRydWV9.y5d0bRP6Fjv1vOrYwzsjccqOe_b4zHZkO8fulgyhuzg", "my-secret")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Token verified:", isVerifiedToken)
}
