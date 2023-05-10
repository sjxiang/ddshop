package main

import (
	"github.com/sjxiang/ddshop/goods_srv/internal/pkg/conf"
	"github.com/sjxiang/ddshop/goods_srv/internal/dal"
)


func Init() {
	conf.Init()
	dal.Init()
}

func main() {
	Init()
}