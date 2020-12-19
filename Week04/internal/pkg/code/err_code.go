package code

import "github.com/pkg/errors"

var StoreFail = errors.New("资产入库失败")

var NotFound = errors.New("资产不存在")
