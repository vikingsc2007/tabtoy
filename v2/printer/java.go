package printer

import (
	"fmt"
	"text/template"

	"github.com/davyxu/tabtoy/v2/i18n"
	"github.com/davyxu/tabtoy/v2/model"
	"strings"
)

const javaTemplate = `// Generated by github.com/davyxu/tabtoy
// Version: {{.ToolVersion}}
// DO NOT EDIT!!
#include <vector>
#include <map>
#include <string>

namespace {{.Namespace}}{{$globalIndex:=.Indexes}}{{$verticalFields:=.VerticalFields}}
{
	{{range .Enums}}
	// Defined in table: {{.DefinedTable}}
	enum class {{.Name}}
	{
	{{range .Fields}}	
		{{.Comment}}
		{{.FieldDescriptor.Name}} = {{.FieldDescriptor.EnumValue}}, {{.Alias}}
	{{end}}
	};
	{{end}}

	{{range .Classes}}

	{{if not .IsCombine}}

	// Defined in table: {{.DefinedTable}}
	class {{.Name}}
	{
	public:
	{{range .Fields}}
		{{.Comment}}
		{{.TypeCode}} {{.Alias}}
	{{end}}

	}; {{end}}
	{{end}}

	{{range .Classes}}
	{{if .IsCombine}}

	// Defined in table: {{.DefinedTable}}	
	class {{.Name}}
	{
	{{if .IsCombine}}
	public:
		tabtoy::Logger TableLogger;
	{{end}}
	{{range .Fields}}	
		{{.Comment}}
		{{.TypeCode}} {{.Alias}}
	{{end}}
	{{if .IsCombine}}
		//#region Index code
	 	{{range $globalIndex}}std::map<{{.IndexType}}, {{.RowType}}> _{{.RowName}}By{{.IndexName}};
	public:
		class {{.RowType}}* Get{{.RowName}}By{{.IndexName}}({{.IndexType}} {{.IndexName}}, {{.RowType}}* def = nullptr)
        {
            auto ret = _{{.RowName}}By{{.IndexName}}.find( {{.IndexName}} );
            if ( ret != _{{.RowName}}By{{.IndexName}}.end() )
            {
                return &ret->second;
            }
			
			if ( def == nullptr )
			{
				TableLogger.ErrorLine("Get{{.RowName}}By{{.IndexName}} failed, {{.IndexName}}: %s", {{.IndexName}});
			}

            return def;
        }
		{{end}}
	{{range $verticalFields}}
	public:
		class {{.StructName}}* Get{{.Name}}( )
		{
			return &{{.Name}}_[0];
		}	
	{{end}}
		//#endregion
		//#region Deserialize code
		{{range $.Classes}}
	public:
		static void Deserialize( {{.Name}}& ins, tabtoy::DataReader& reader )
		{
 			int tag = -1;
            while ( -1 != (tag = reader.ReadTag()))
            {
                switch (tag)
                { {{range .Fields}}
                	case {{.Tag}}:
                	{
						{{.ReadCode}}
                	}
                	break; {{end}}
                }
             }

			{{range $a, $row :=.IndexedFields}}
			// Build {{$row.FieldDescriptor.Name}} Index
			for( size_t i = 0;i< ins.{{$row.FieldDescriptor.Name}}_.size();i++)
			{
				auto element = ins.{{$row.FieldDescriptor.Name}}_[i];
				{{range $b, $key := .IndexKeys}}
				ins._{{$row.FieldDescriptor.Name}}By{{$key.Name}}.emplace(std::make_pair(element.{{$key.Name}}_, element));
				{{end}}
			}
			{{end}}
		}{{end}}
		//#endregion
	{{end}}

	};
	{{end}}
	{{end}}
}
`

type javaIndexField struct {
	TableIndex
}

func (self javaIndexField) IndexName() string {
	return self.Index.Name
}

func (self javaIndexField) RowType() string {
	return self.Row.Complex.Name
}

func (self javaIndexField) RowName() string {
	return self.Row.Name
}

func (self javaIndexField) IndexType() string {

	switch self.Index.Type {
	case model.FieldType_Int32:
		return "int"
	case model.FieldType_UInt32:
		return "long"
	case model.FieldType_Int64:
		return "long"
	case model.FieldType_UInt64:
		return "long"
	case model.FieldType_String:
		return "String"
	case model.FieldType_Float:
		return "float"
	case model.FieldType_Bool:
		return "boolean"
	case model.FieldType_Enum:

		return self.Index.Complex.Name
	default:
		log.Errorf("%s can not be index ", self.Index.String())
	}

	return "unknown"
}

type javaField struct {
	*model.FieldDescriptor

	IndexKeys []*model.FieldDescriptor

	parentStruct *javaStructModel
}

func (self javaField) Alias() string {

	v := self.FieldDescriptor.Meta.GetString("Alias")
	if v == "" {
		return ""
	}

	return "// " + v
}

func (self javaField) Comment() string {

	if self.FieldDescriptor.Comment == "" {
		return ""
	}

	// zjwps 建议修改
	return "/// <summary> \n		/// " + strings.Replace(self.FieldDescriptor.Comment, "\n", "\n		///", -1) + "\n		/// </summary>"
}

func (self javaField) ReadCode() string {

	var baseType string

	var descHandlerCode string

	switch self.Type {
	case model.FieldType_Int32:
		baseType = "Int32"
	case model.FieldType_UInt32:
		baseType = "UInt32"
	case model.FieldType_Int64:
		baseType = "Int64"
	case model.FieldType_UInt64:
		baseType = "UInt64"
	case model.FieldType_String:
		baseType = "String"
	case model.FieldType_Float:
		baseType = "Float"
	case model.FieldType_Bool:
		baseType = "Bool"
	case model.FieldType_Enum:

		if self.Complex == nil {
			return "unknown"
		}

		baseType = fmt.Sprintf("Enum<%s>", self.Complex.Name)

	case model.FieldType_Struct:
		if self.Complex == nil {
			return "unknown"
		}

		baseType = fmt.Sprintf("Struct<%s>", self.Complex.Name)

	}

	if self.Type == model.FieldType_Struct {
		descHandlerCode = "Deserialize"
		//descHandlerCode = fmt.Sprintf("%sDeserializeHandler", self.Complex.Name)
	}

	if self.IsRepeated {
		return fmt.Sprintf("ins.%s_.emplace_back( reader.Read%s(%s) );", self.Name, baseType, descHandlerCode)
	} else {
		return fmt.Sprintf("ins.%s_ = reader.Read%s(%s);", self.Name, baseType, descHandlerCode)
	}

}

func (self javaField) Tag() string {

	if self.parentStruct.IsCombine() {
		tag := model.MakeTag(int32(model.FieldType_Table), self.Order)

		return fmt.Sprintf("0x%x", tag)
	}

	return fmt.Sprintf("0x%x", self.FieldDescriptor.Tag())
}

func (self javaField) StructName() string {
	if self.Complex == nil {
		return "[NotComplex]"
	}

	return self.Complex.Name
}

func (self javaField) IsVerticalStruct() bool {
	if self.FieldDescriptor.Complex == nil {
		return false
	}

	return self.FieldDescriptor.Complex.File.Pragma.GetBool("Vertical")
}

func (self javaField) TypeCode() string {

	var raw string

	switch self.Type {
	case model.FieldType_Int32:
		raw = "int"
	case model.FieldType_UInt32:
		raw = "unsigned int"
	case model.FieldType_Int64:
		raw = "long long"
	case model.FieldType_UInt64:
		raw = "unsigned long long"
	case model.FieldType_String:
		raw = "std::string"
	case model.FieldType_Float:
		raw = "float"
	case model.FieldType_Bool:
		raw = "bool"
	case model.FieldType_Enum:
		if self.Complex == nil {
			log.Errorln("unknown enum type ", self.Type)
			return "unknown"
		}

		raw = self.Complex.Name
	case model.FieldType_Struct:
		if self.Complex == nil {
			log.Errorln("unknown struct type ", self.Type, self.FieldDescriptor.Name, self.FieldDescriptor.Parent.Name)
			return "unknown"
		}

		raw = self.Complex.Name

		// 非repeated的结构体
		if !self.IsRepeated {
			return fmt.Sprintf("public:\r\n \t\t%s %s_;", raw, self.Name)
		}

	default:
		raw = "unknown"
	}

	if self.IsRepeated {
		return fmt.Sprintf("public:\r\n \t\tstd::vector<%s> %s_;", raw, self.Name)
	}

	return fmt.Sprintf("public:\r\n \t\t%s %s_ = %s;", raw, self.Name, wrapJavaDefaultValue(self.FieldDescriptor))
}

func wrapJavaDefaultValue(fd *model.FieldDescriptor) string {
	switch fd.Type {
	case model.FieldType_Enum:
		return fmt.Sprintf("%s::%s", fd.Complex.Name, fd.DefaultValue())
	case model.FieldType_String:
		return fmt.Sprintf("\"%s\"", fd.DefaultValue())
	case model.FieldType_Float:
		var defValue = fd.DefaultValue()
		if !strings.ContainsAny(defValue, ".") {
			return fmt.Sprintf("%s.0f", defValue)
		}
		return fmt.Sprintf("%sf", defValue)
	}

	return fd.DefaultValue()
}

type javaStructModel struct {
	*model.Descriptor
	Fields        []javaField
	IndexedFields []javaField // 与javaField.IndexKeys组成树状的索引层次
}

func (self *javaStructModel) DefinedTable() string {
	return self.File.Name
}

func (self *javaStructModel) Name() string {
	return self.Descriptor.Name
}

func (self *javaStructModel) IsCombine() bool {
	return self.Descriptor.Usage == model.DescriptorUsage_CombineStruct
}

type javaFileModel struct {
	Namespace   string
	ToolVersion string
	Classes     []*javaStructModel
	Enums       []*javaStructModel
	Indexes     []javaIndexField // 全局的索引

	VerticalFields []javaField
}

type javaPrinter struct {
}

func (self *javaPrinter) Run(g *Globals) *Stream {

	tpl, err := template.New("java").Parse(javaTemplate)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	var m javaFileModel

	m.Namespace = g.FileDescriptor.Pragma.GetString("Package")
	m.ToolVersion = g.Version

	// combinestruct的全局索引
	for _, ti := range g.GlobalIndexes {

		// 索引也限制
		if !ti.Index.Parent.File.MatchTag(".java") {
			continue
		}

		m.Indexes = append(m.Indexes, javaIndexField{TableIndex: ti})
	}

	// 遍历所有类型
	for _, d := range g.FileDescriptor.Descriptors {

		// 这给被限制输出
		if !d.File.MatchTag(".java") {
			log.Infof("%s: %s", i18n.String(i18n.Printer_IgnoredByOutputTag), d.Name)
			continue
		}

		var sm javaStructModel
		sm.Descriptor = d

		switch d.Kind {
		case model.DescriptorKind_Struct:
			m.Classes = append(m.Classes, &sm)
		case model.DescriptorKind_Enum:
			m.Enums = append(m.Enums, &sm)
		}

		// 遍历字段
		for _, fd := range d.Fields {

			// 对CombineStruct的XXDefine对应的字段
			if d.Usage == model.DescriptorUsage_CombineStruct {

				// 这个字段被限制输出
				if fd.Complex != nil && !fd.Complex.File.MatchTag(".java") {
					continue
				}

				// 这个结构有索引才创建
				if fd.Complex != nil && len(fd.Complex.Indexes) > 0 {

					// 被索引的结构
					indexedField := javaField{FieldDescriptor: fd, parentStruct: &sm}

					// 索引字段
					for _, key := range fd.Complex.Indexes {
						indexedField.IndexKeys = append(indexedField.IndexKeys, key)
					}

					sm.IndexedFields = append(sm.IndexedFields, indexedField)
				}

				if fd.Complex != nil && fd.Complex.File.Pragma.GetBool("Vertical") {
					m.VerticalFields = append(m.VerticalFields, javaField{FieldDescriptor: fd, parentStruct: &sm})
				}

			}

			csField := javaField{FieldDescriptor: fd, parentStruct: &sm}

			sm.Fields = append(sm.Fields, csField)

		}

	}

	bf := NewStream()

	err = tpl.Execute(bf.Buffer(), &m)
	if err != nil {
		log.Errorln(err)
		return nil
	}

	return bf
}

func init() {

	RegisterPrinter("java", &javaPrinter{})

}
