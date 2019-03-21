package gen

import "github.com/vikingsc2007/tabtoy/v3/model"

type GenFunc func(globals *model.Globals) (data []byte, err error)
