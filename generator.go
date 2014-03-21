package idgen

import (
	"sync/atomic"
	"time"
	"sync"
)

type IdGenerator struct{
	MechineId uint64
	DataCenterId uint64
	idSequence *int64
	timeStamp uint64
	mutex sync.Mutex
}

var workIdShift=uint64(12)
var dataCenterIdShift =uint64(17)
var timestampLeftShift =uint64(22)
var sequnceMask=uint64(4095)
var baseTimestamp =uint64(1318323746000)

func NewIdGenerator(mechineId,dataCenterId uint64)IdGenerator{
	timestamp:=time.Now().UnixNano()/1000000
	zero:=int64(0)
	idGenerator:=IdGenerator{
		MechineId:mechineId,
		DataCenterId:dataCenterId,
		idSequence:&zero,
		timeStamp:uint64(timestamp),
	}
	go idGenerator.cleanSequence()
	return idGenerator
}

func (this IdGenerator) cleanSequence(){
	for{
		this.mutex.Lock()
		atomic.SwapInt64(this.idSequence,int64(0))
		this.timeStamp=uint64(time.Now().UnixNano()/1000000)
		this.mutex.Unlock()
		time.Sleep(time.Second)
	}
}

func (this *IdGenerator) GetId() uint64 {
	seq:=atomic.AddInt64(this.idSequence,1)
	newId := ((this.timeStamp - baseTimestamp) << timestampLeftShift) | (this.DataCenterId << dataCenterIdShift) | (this.MechineId << workIdShift) | uint64(seq)
	return uint64(newId)
}
