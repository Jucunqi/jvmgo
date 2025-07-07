package heap

type ExceptionHandler struct {
	startPc   int       // 其实pc
	endPc     int       // 结束pc
	handlerPc int       // 跳转pc
	catchType *ClassRef // 捕获类型
}

type ExceptionTable []*ExceptionHandler

func (t ExceptionTable) findExceptionHandler(exClass *Class, pc int) *ExceptionHandler {

	for _, handler := range t {
		// 判断pc是否处于当前异常类的起始和结束区间内
		if pc > handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler // catch - all
			}
			class := handler.catchType.ResolveClass()
			if class == exClass || exClass.IsSubClassOf(class) {
				// 并且处理类型是exClass的子类
				return handler
			}
		}
	}
	return nil
}
