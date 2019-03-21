package luasrc

import (
	"github.com/davyxu/protoplus/codegen"
	"github.com/vikingsc2007/tabtoy/v3/gen"
	"github.com/vikingsc2007/tabtoy/v3/model"
)

func Generate(globals *model.Globals) (data []byte, err error) {

	err = codegen.NewCodeGen("luasrc").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(gen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc).
		ParseTemplate(templateText, globals).
		WriteBytes(&data).Error()

	return
}
