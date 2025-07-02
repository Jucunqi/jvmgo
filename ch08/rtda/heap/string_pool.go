package heap

import "unicode/utf16"

// 用map表示字符串池，key是go字符串，value是Java字符串
var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {

	// 从字符串池中获取
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}
	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)
	internedStrings[goStr] = jStr
	return jStr
}

func stringToUtf16(str string) []uint16 {
	runes := []rune(str) // utf32
	return utf16.Encode(runes)
}

func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s)
	return string(runes)
}
