package nicclient

import (
	db "cc/DB"
	"fmt"
)

type Reciver struct {
	NodeIPS  []string
	NodeName []string
}

type INFO interface {
	NICINFO(NICCLIENT *Reciver) (string, error)
	NewReciver(DBClient *db.Client) (Reciver, error)
}

func NewReciver(DBClient *db.Client) (Reciver, error) {
	if DBClient.ReadAccess == false {
		err := fmt.Errorf("Read  Access is not granted")
		return Reciver{}, err
	}

	IPs, err := DBClient.GETIPS()
	if err != nil {
		return Reciver{}, err
	}
	// Names, err := DBClient.GetNames()

	return Reciver{
		NodeIPS:  IPs,
		NodeName: []string{},
	}, nil
}

func (NICCLIENT *Reciver) NICINFO() {
	fmt.Println(NICCLIENT.NodeIPS)
}
