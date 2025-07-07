package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func (s *Stack) push(frame *Frame) {

	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if s._top != nil {
		frame.lower = s._top
	}
	s._top = frame
	s.size++
}

func (s *Stack) pop() *Frame {

	if s._top == nil {
		panic("jvm stack is empty")
	}
	top := s._top
	s._top = top.lower
	s.size--
	return top

}

func (s *Stack) top() *Frame {

	if s._top == nil {
		panic("jvm stack is empty!")
	}
	return s._top
}

func (s *Stack) isEmpty() bool {
	return s._top == nil
}

func (s *Stack) Clear() {
	for !s.isEmpty() {
		s.pop()
	}
}

func (s *Stack) GetFrames() []*Frame {
	frames := make([]*Frame, 0, s.size)
	for frame := s._top; frame != nil; frame = frame.lower {
		frames = append(frames, frame)
	}
	return frames
}

func newStack(size uint) *Stack {
	return &Stack{maxSize: size}
}
