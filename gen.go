package bebop

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type GenerateSettings struct {
	PackageName string
}

func (f File) Validate() error {
	// TODO: check no user defined type name is the same as a primitive (or reserved, like map/array/struct)
	// TODO: check no duplicate type names are defined
	// TODO: check that all types make sense (are primitive, or exist in the file)
	// TODO: importing other files? oh god
	return nil
}

func (f File) Generate(w io.Writer, settings GenerateSettings) error {
	if err := f.Validate(); err != nil {
		return fmt.Errorf("cannot generate file: %w", err)
	}

	writeLine(w, "// Code generated by bebopc-go; DO NOT EDIT.")
	writeLine(w, "")
	writeLine(w, "package %s", settings.PackageName)
	writeLine(w, "")
	if f.hasDateType() {
		writeLine(w, "import (")
		writeLine(w, "\t\"time\"")
		writeLine(w, ")")
		writeLine(w, "")
	}

	for _, en := range f.Enums {
		en.Generate(w, settings)
	}
	for _, st := range f.Structs {
		st.Generate(w, settings)
	}
	for _, msg := range f.Messages {
		msg.Generate(w, settings)
	}
	return nil
}

func writeLine(w io.Writer, format string, args ...interface{}) {
	fmt.Fprintf(w, format+"\n", args...)
}

func (en Enum) Generate(w io.Writer, settings GenerateSettings) {
	exposedName := exposeName(en.Name)
	if en.Comment != "" {
		writeLine(w, "// %s", en.Comment)
	}
	writeLine(w, "type %s uint32", exposedName)
	writeLine(w, "")
	writeLine(w, "const(")
	for _, opt := range en.Options {
		if opt.Comment != "" {
			writeLine(w, "\t// %s", opt.Comment)
		}
		if opt.Deprecated {
			writeLine(w, "\t// Deprecated: %s", opt.DeprecatedMessage)
		}
		writeLine(w, "\t%s_%s %s = %d", exposedName, opt.Name, exposedName, opt.Value)
	}
	writeLine(w, ")")
	writeLine(w, "")
}

func (st Struct) Generate(w io.Writer, settings GenerateSettings) {
	exposedName := exposeName(st.Name)
	if st.OpCode != 0 {
		writeLine(w, "const %sOpCode = 0x%x", exposedName, st.OpCode)
		writeLine(w, "")
	}
	if st.Comment != "" {
		writeLine(w, "// %s", st.Comment)
	}
	writeLine(w, "type %s struct {", exposedName)
	for _, fd := range st.Fields {
		if fd.Comment != "" {
			writeLine(w, "\t// %s", fd.Comment)
		}
		if fd.Deprecated {
			writeLine(w, "\t// Deprecated: %s", fd.DeprecatedMessage)
		}

		name := exposeName(fd.Name)
		if st.ReadOnly {
			name = unexposeName(fd.Name)
		}
		writeLine(w, "\t%s %s", name, fd.FieldType.GoString())
	}
	writeLine(w, "}")
	writeLine(w, "")
	// TODO: encode, decode methods
	// TODO: getters for readonly
}

func (msg Message) Generate(w io.Writer, settings GenerateSettings) {
	exposedName := exposeName(msg.Name)
	if msg.OpCode != 0 {
		writeLine(w, "const %sOpCode = 0x%x", exposedName, msg.OpCode)
		writeLine(w, "")
	}
	if msg.Comment != "" {
		writeLine(w, "// %s", msg.Comment)
	}
	writeLine(w, "type %s struct {", exposedName)
	type FieldWithNumber struct {
		Field
		num int32
	}
	fields := make([]FieldWithNumber, 0, len(msg.Fields))
	for i, fd := range msg.Fields {
		fields = append(fields, FieldWithNumber{
			Field: fd,
			num:   i,
		})
	}
	sort.Slice(fields, func(i, j int) bool {
		return fields[i].num < fields[j].num
	})
	for _, fd := range fields {
		if fd.Comment != "" {
			writeLine(w, "\t// %s", fd.Comment)
		}
		if fd.Deprecated {
			writeLine(w, "\t// Deprecated: %s", fd.DeprecatedMessage)
		}

		name := exposeName(fd.Name)
		if msg.ReadOnly {
			name = unexposeName(fd.Name)
		}
		writeLine(w, "\t%s %s", name, fd.FieldType.GoString())
	}
	writeLine(w, "}")
	writeLine(w, "")
	// TODO: encode, decode methods
	// TODO: getters for readonly
}

func exposeName(name string) string {
	if len(name) > 1 {
		return strings.ToUpper(string(name[0])) + name[1:]
	}
	if len(name) > 0 {
		return strings.ToUpper(string(name[0]))
	}
	return ""
}

func unexposeName(name string) string {
	if len(name) > 1 {
		return strings.ToLower(string(name[0])) + name[1:]
	}
	if len(name) > 0 {
		return strings.ToLower(string(name[0]))
	}
	return ""
}
