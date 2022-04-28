package serverid

import (
	"github.com/google/uuid"
)

type ServerId uuid.UUID

type ServerIdService interface {
	GetServerId() ServerId
}

type memServerIdService struct {
	sId ServerId
}

func (this *memServerIdService) GetServerId() ServerId {
	return this.sId
}
