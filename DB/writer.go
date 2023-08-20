package db

import (
	collector "cc/Collector"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
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

func (DBClient *Client) WriteData(Globaldata collector.GlobalDATA) error {

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
	f, err := os.Create("History/" + CurrentYear + "/" + CurrentMonth + "/" + CurrentDay + "/" + CurrentHour + "/" + CurrentMinute + "/" + CurrentMinute + "info.json")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	err = json.NewEncoder(f).Encode(Globaldata)
	if err != nil {
		return err
	}

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

func (DBClient *Client) ReadData(Year string, Month string, Day string, Hour string, Minute string, object string, name string) error {

	// read the data from the database
	// and return the data

	if Year > strconv.Itoa(time.Now().Year()) {
		if Month > strconv.Itoa(int(time.Now().Month())) {
			if Day > strconv.Itoa(time.Now().Day()) {
				if Hour > strconv.Itoa(time.Now().Hour()) {
					if Minute > strconv.Itoa(time.Now().Minute()) {
						return fmt.Errorf("invalid time")
					}
				}
			}
		}
	}

	if DBClient.ReadAccess == false {
		return fmt.Errorf("read Access is not given")
	}

	SearchPath := "History/" + Year + "/" + Month + "/" + Day + "/" + Hour + "/" + Minute + "/" + Minute + "info.json"

	file, err := os.Open(SearchPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var data collector.GlobalDATA

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return err
	}

	// use switch case to return the data

	if name != "" {
		switch object {
		case "deployments":
			for _, v := range data.DeployInfo.Deployments {
				if v.Name == name {
					fmt.Println("─────────────", v.Name, "─────────────")
					fmt.Println(v.Namespace)
					fmt.Println(v.CreationTimestamp)
					fmt.Println(v.Replicas)
					fmt.Println(v.AvailableRep)
					fmt.Println(v.ReadyRep)
					fmt.Println(v.UpToDateRep)
					fmt.Println(v.AvailableUpToDateRep)
					fmt.Println(v.Age)
					fmt.Println(v.Spec)
					fmt.Println(v.Labels)
					fmt.Println(v.Selector)
					fmt.Println(v.Strategy)
					fmt.Println(v.MinReadySeconds)
					fmt.Println(v.RevisionHistoryLimit)
					fmt.Println(v.Paused)
					fmt.Println(v.ProgressDeadlineSeconds)
					fmt.Println(v.ReplicaSet)
					fmt.Println(v.Conditions)

					fmt.Println("──────────────────────────────────")
				}
			}

		case "pods":
			for _, v := range data.PodInfo.PODS {
				if v.PodName == name {
					fmt.Println("─────────────", v.PodName, "─────────────")
					fmt.Println(v.PodCreation)
					fmt.Println(v.PodDeletionGracePeriodSeconds)
					fmt.Println(v.PodPhase)
					fmt.Println(v.PodRunningOn)
					fmt.Println(v.PodLabels)
					fmt.Println(v.PodNamespace)
					fmt.Println(v.PodAnnotation)

					fmt.Println(v.PodOwnerReferences)
					fmt.Println(v.PodResourceVersion)
					fmt.Println(v.PodUID)
					fmt.Println(v.PodDNSConfig)
					fmt.Println(v.PodDNSPolicy)
					fmt.Println(v.PodEnableServiceLinks)
					fmt.Println(v.PodEphemeralContainers)
					fmt.Println(v.PodHostIpc)
					fmt.Println(v.PodHostNetwork)
					fmt.Println(v.PodHostPID)
					fmt.Println(v.PodHostUsers)
					fmt.Println(v.PodImagePullSecrets)
					fmt.Println(v.PodInitContainers)
					fmt.Println(v.PodRestartPolicy)
					fmt.Println(v.PodRuntimeClassName)
					fmt.Println("──────────────────────────────────")
				}
			}

		case "default":
			return fmt.Errorf("invalid object")
		}

	} else {

		switch object {
		case "deployments":
			fmt.Println(data.DeployInfo.Deployments)

		case "pods":
			fmt.Println(data.PodInfo)

		case "default":
			return fmt.Errorf("invalid object")
		}

	}

	return nil
}
