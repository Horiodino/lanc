package db

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type HASHVALUE struct {
	hashid    []byte
	timestamp time.Time
}

/*
 	things to do
	create an only way to write and access the data from the database
	create a function that will read the data from the database

*/

var ConstHashIdAddtime time.Time
var ConstHashValue []byte

type Client struct {
	ReadAccess  bool
	WriteAccess bool
	HashId      string
}

type AcessDB interface {
	DecryptHashID(ciphertext []byte) ([]byte, error)
	GetHashID(DBClient *Client) (string, string)
	WriteDATA(DBClient *Client) error
	ReadDAta(DBClient *Client) error
	LockFile(DBClient *Client) error
}

func CreateClient() *Client {

	// check if hte hashid.txt file is present or not
	// and if it is present then read the data from it and save it in the hashid variable
	// check the timestamp and if the timestamp is more than 24 hours then generate a new hashid and save it in the hashid.txt file
	// and then return the client

	return &Client{
		ReadAccess:  true,
		WriteAccess: true,
		HashId:      string(ConstHashValue),
	}
}

func GenerateHashID() []byte {

	key := []byte("123456789qwertyuioplkjhgfdsazxcv") // 32-byte AES key
	plaintext := []byte("Hello, World!")

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return nil
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	ConstHashValue = ciphertext
	ConstHashIdAddtime = time.Now()

	fmt.Println(ciphertext)

	HASHID := HASHVALUE{
		hashid:    ciphertext,
		timestamp: ConstHashIdAddtime,
	}

	file, err := os.Create("hashid.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	_, err = fmt.Fprintf(file, "hashid: %x\n", HASHID.hashid)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(file, "timestamp: %s\n", HASHID.timestamp.Format(time.RFC3339Nano))
	if err != nil {
		log.Fatal(err)
	}

	// now  lock the file so that no one can access it in terms of write only until 24 hours
	// and after 24 hours, we will generate a new hashid and will save it in the file

	// NOTE this is temporary, so it will create a new hashid everytime the program is run

	DBCLIENT := CreateClient()
	_ = DBCLIENT.LockFile()

	return ciphertext

}

func (DBClient *Client) GetHashID() (string, string) {
	file, err := os.Open("hashid.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var hashid string
	var timestamp string

	_, err = fmt.Fscanf(file, "hashid: %s\n", &hashid)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fscanf(file, "timestamp: %s\n", &timestamp)
	if err != nil {
		log.Fatal(err)
	}

	return hashid, timestamp
}

func (DBClient *Client) DecryptHashID(ciphertext []byte) ([]byte, error) {
	key := []byte("123456789qwertyuioplkjhgfdsazxcv") // 32-byte AES key

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	nonceSize := 12
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func (DBClient *Client) LockFile() error {

	// we will lock the file for 23:59:59 hours
	// and after that we will generate a new hashid and will save it in the file
	//until then the file will be locked and no one can access it except read only

	return nil
}

func (DBClient *Client) WriteData() error {

	if DBClient.WriteAccess == false {
		return fmt.Errorf("write Access is not given")
	}

	CurrentYear := strconv.Itoa(time.Now().Year())
	CurrentMonth := strconv.Itoa(int(time.Now().Month()))
	CurrentDay := strconv.Itoa(time.Now().Day())
	CurrentHour := strconv.Itoa(time.Now().Hour())
	CurrentMinute := strconv.Itoa(time.Now().Minute())

	if exists(CurrentYear, CurrentMonth, CurrentDay, CurrentHour, CurrentMinute, 1) == false {
		err := os.Mkdir("History/"+CurrentYear, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if exists(CurrentYear, CurrentMonth, CurrentDay, CurrentHour, CurrentMinute, 2) == false {
		err := os.Mkdir("History/"+CurrentYear+"/"+CurrentMonth, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if exists(CurrentYear, CurrentMonth, CurrentDay, CurrentHour, CurrentMinute, 3) == false {
		err := os.Mkdir("History/"+CurrentYear+"/"+CurrentMonth+"/"+CurrentDay, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if exists(CurrentYear, CurrentMonth, CurrentDay, CurrentHour, CurrentMinute, 4) == false {
		err := os.Mkdir("History/"+CurrentYear+"/"+CurrentMonth+"/"+CurrentDay+"/"+CurrentHour, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	if exists(CurrentYear, CurrentMonth, CurrentDay, CurrentHour, CurrentMinute, 5) == false {
		err := os.Mkdir("History/"+CurrentYear+"/"+CurrentMonth+"/"+CurrentDay+"/"+CurrentHour+"/"+CurrentMinute, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}

	// now save the file in the folder
	// SaveDataC func can be only called like this WriteData.SaveDataC()

	return nil
}

func exists(current, CurrentMonth, CurrentDay, CurrentHour, CurrentMinute string, ID int) bool {
	path := "History/" + current

	switch ID {
	case 1:
		path += "/"
	case 2:
		path += "/" + CurrentMonth + "/"
	case 3:
		path += "/" + CurrentMonth + "/" + CurrentDay + "/"
	case 4:
		path += "/" + CurrentMonth + "/" + CurrentDay + "/" + CurrentHour + "/"
	case 5:
		path += "/" + CurrentMonth + "/" + CurrentDay + "/" + CurrentHour + "/" + CurrentMinute + "/"
	default:
		return false
	}

	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func (DBClient *Client) ReadData() error {

	return nil
}
