package idgen

import (
	"testing"
	"fmt"
)

func TestGetId(t *testing.T) {
	idGenerator:=NewIdGenerator(1,1)
	id:=(&idGenerator).GetId()
	fmt.Println(id)
}



















