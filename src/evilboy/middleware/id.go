package middleware

import (
	"math"
	"sync"
)

// ID生成器的接口类型。
type IdGenertor interface {
	GetUint32() uint32 // 获得一个uint32类型的ID。
}

// 创建ID生成器。
func NewIdGenertor() IdGenertor {
	return &cyclicIdGenertor{}
}

// ID生成器的实现类型。
type cyclicIdGenertor struct {
	sn    uint32     // 当前的ID。
	ended bool       // 前一个ID是否已经为其类型所能表示的最大值。
	mutex sync.Mutex // 互斥锁。
}

func (gen *cyclicIdGenertor) GetUint32() uint32 {
	gen.mutex.Lock()
	defer gen.mutex.Unlock()
	if gen.ended {
		defer func() { gen.ended = false }()
		gen.sn = 0
		return gen.sn
	}
	id := gen.sn
	if id < math.MaxUint32 {
		gen.sn++
	} else {
		gen.ended = true
	}
	return id
}

// ID生成器的接口类型2。
type IdGenertor2 interface {
	GetUint64() uint64 // 获得一个uint64类型的ID。
}

// 创建ID生成器2。
func NewIdGenertor2() IdGenertor2 {
	return &cyclicIdGenertor2{}
}

// ID生成器的实现类型2。
type cyclicIdGenertor2 struct {
	base       cyclicIdGenertor // 基本的ID生成器。
	cycleCount uint64           // 基于uint32类型的取值范围的周期计数。
}

func (gen *cyclicIdGenertor2) GetUint64() uint64 {
	var id64 uint64
	if gen.cycleCount%2 == 1 {
		id64 += math.MaxUint32
	}
	id32 := gen.base.GetUint32()
	if id32 == math.MaxUint32 {
		gen.cycleCount++
	}
	id64 += uint64(id32)
	return id64
}
