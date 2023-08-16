package db

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// in structure we will create a table that will have the address of all the folders that are created or
// we saved the json files in it .

// suppose user today created a folder and saved a json file in it . so we will save the address of that folder

// lets see an example

/*

	today date : 12/12/2020
	user wan to acess yesterday data so he will go to the folder of 11/12/2020 and will get the data from there


	we will save the address of the folder when it was created and will save it in the database
	then we will access the data from that folder using the address that we have saved in the database



*/

/*
 	things to do
	// create a only way to write and access the data from the database
	// we will create client and then we will use that client to write and read the data from the database

	create a function that will create a client and will return the client
	create a function that will write the data in the database
	create a function that will read the data from the database

*/

var (
	ConstHashId string
)

type Client struct {
	ReadAccess  bool
	WriteAccess bool
	HashId      string
}

type AcessDB interface {
	WriteDATA(CLIENT *Client) error
	ReadDAta(CLIENT *Client) error
}

func CreateClient() *Client {
	return &Client{}
}

func GetHAshID() string {
	return ConstHashId
}

func GenerateHashID() []byte {

	key := []byte("123456789qwertyuioplkjhgfdsazxcv") // 32-byte AES key
	plaintext := []byte("Hello, World!")

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	return ciphertext
}

func DecryptHashID() {

}

func (CLIENT *Client) WriteData() error {

	return nil
}

func (CLIENT *Client) ReadData() error {

	return nil
}
