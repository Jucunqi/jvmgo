package heap

func getArrayClassName(name string) string {
	return "[" + toDescriptor(name)
}

func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}

var primitiveTypes = map[string]string{

	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

func getComponentClassName(className string) string {

	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

func toClassName(descriptor string) string {

	// array
	if descriptor[0] == '[' {
		return descriptor
	}

	// object
	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}

	// primitive
	for className, d := range primitiveTypes {
		if d == descriptor {
			return className
		}
	}
	panic("Invalid descriptor : " + descriptor)
}
