package idgen

import (
	"testing"
)

func TestGetId(t *testing.T) {
	idGenerator:=NewIdGenerator(1,1)
	id:=(&idGenerator).GetId()
}
