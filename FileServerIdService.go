package serverid

import (
	"encoding/json"
	"io"
	"os"

	"github.com/google/uuid"
)

const (
	defaultSIdFilePath string = "server-id.json"
)

type serverIdJsonFile struct {
	ServerIdString string `json:"server-id"`
}

func NewFileServerIdService(fPath string) (ServerIdService, error) {
	if fPath == "" {
		fPath = defaultSIdFilePath
	}

	r, err := getServerIdReader(fPath)
	defer r.Close()
	var sId uuid.UUID
	if err == nil {
		sId, err = getServerIdFromReader(r)
	}

	if err == nil {
		return &memServerIdService{sId: sId}, nil
	}
	return nil, err
}

func getServerIdFromReader(r io.Reader) (uuid.UUID, error) {
	var sId uuid.UUID
	decoder := json.NewDecoder(r)
	sIdFile := &serverIdJsonFile{}
	err := decoder.Decode(sIdFile)
	if err == nil {
		sId, err = uuid.Parse(sIdFile.ServerIdString)

	}
	return sId, err
}

func getServerIdReader(path string) (io.ReadCloser, error) {

	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		//We need to create the file
		err = createNewServerIdFile(path)
	}

	var f *os.File
	if err == nil {
		f, err = os.Open(path)
	}
	return f, err
}

func createNewServerIdFile(path string) error {
	var f *os.File
	var err error
	f, err = os.Create(path)
	defer f.Close()
	if err == nil {
		encoder := json.NewEncoder(f)
		newId := uuid.New()
		sIdFile := &serverIdJsonFile{ServerIdString: newId.String()}
		err = encoder.Encode(sIdFile)
	}
	return err
}
