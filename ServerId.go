package serverid

import (
	"github.com/google/uuid"
)

//type ServerId uuid.UUID

type ServerIdService interface {
	GetServerId() uuid.UUID
}

type memServerIdService struct {
	sId uuid.UUID
}

func (this *memServerIdService) GetServerId() uuid.UUID {
	return this.sId
}
