package practice_leetcode

type stack struct {
	s []interface{}
}

func (s *stack) push(v interface{}) {
	s.s = append(s.s, v)
}

func (s *stack) pop() interface{} {
	if s.len() == 0 {
		return nil
	}
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}

func (s *stack) len() int {
	return len(s.s)
}

func isValid(v string) bool {
	s := new(stack)
	for _, ch := range v {
		switch ch {
		case '(':
			s.push(')')
		case '[':
			s.push(']')
		case '{':
			s.push('}')
		case ')', ']', '}':
			v := s.pop()
			if v == nil || v.(rune) != ch {
				return false
			}
		}
	}
	return s.len() == 0
}
