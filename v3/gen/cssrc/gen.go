package cssrc

import (
	"github.com/davyxu/protoplus/codegen"
	"github.com/vikingsc2007/tabtoy/v3/gen"
	"github.com/vikingsc2007/tabtoy/v3/model"
)

func Generate(globals *model.Globals) (data []byte, err error) {

	cg := codegen.NewCodeGen("cssrc").
		RegisterTemplateFunc(codegen.UsefulFunc).
		RegisterTemplateFunc(gen.UsefulFunc).
		RegisterTemplateFunc(UsefulFunc)

	err = cg.ParseTemplate(templateText, globals).Error()
	if err != nil {
		return
	}

	err = cg.WriteBytes(&data).Error()

	return
}
