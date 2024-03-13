package ocpc_media

import (
	"context"
	"sync"
)

/**
 * 华为
 */

type HuaWei struct {
	SecretKey string `json:"secret_key"`
}

var hwpool = sync.Pool{
	New: func() interface{} {
		return &HuaWei{}
	},
}

func NewHuaWei(secret string) IReporter {
	hw, ok := hwpool.Get().(*HuaWei)
	if !ok {
		return nil
	}
	hw.SecretKey = secret
	return hw
}
func (hw *HuaWei) Free() {
	hw.SecretKey = ""
	hwpool.Put(hw)
}

/**
 * 百度
 */

type BaiDu struct {
	SecretKey string `json:"secret_key"`
}

var bdpool = sync.Pool{New: func() interface{} { return &BaiDu{} }}

func NewBaiDu(secret string) IReporter {
	bd, ok := bdpool.Get().(*BaiDu)
	if !ok {
		return nil
	}
	bd.SecretKey = secret
	return bd
}

func (bd *BaiDu) Free() {
	bd.SecretKey = ""
	bdpool.Put(bd)
}

/**
 * 爱奇艺
 */

type AiQiYi struct {
	SecretKey string `json:"secret_key"`
}

var aqypoool = sync.Pool{New: func() interface{} { return &AiQiYi{} }}

func NewAiQiYi(secret string) IReporter {
	a, ok := aqypoool.Get().(*AiQiYi)
	if !ok {
		return nil
	}
	a.SecretKey = secret
	return a
}

func (a *AiQiYi) Free() {
	a.SecretKey = ""
	aqypoool.Put(a)
}

/**
 * 小米
 */

type XiaoMi struct {
	SecretKey string `json:"secret_key"`
}

var xmpoool = sync.Pool{New: func() interface{} { return &XiaoMi{} }}

func NewXiaoMi(secret string) IReporter {
	x, ok := xmpoool.Get().(*XiaoMi)
	if !ok {
		return nil
	}
	x.SecretKey = secret
	return x
}

func (x *XiaoMi) Free() {
	x.SecretKey = ""
	xmpoool.Put(x)
}

type IReporter interface {
	Report(ctx context.Context, payload []byte) (string, error)
	Free()
}
