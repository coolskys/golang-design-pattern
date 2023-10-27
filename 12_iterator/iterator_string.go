package iterator

type Strings struct {
	start rune
	end   rune
}

func NewStrings(start rune, end rune) *Strings {
	return &Strings{
		start: start,
		end:   end,
	}
}

func (s *Strings) Iterator() Iterator {
	return &StringsIterator{
		ss:   s,
		next: s.start,
	}
}

type StringsIterator struct {
	ss   *Strings
	next rune
}

func (s *StringsIterator) First() {
	s.next = s.ss.start
}

func (s *StringsIterator) IsDone() bool {
	return s.next > s.ss.end
}

func (s *StringsIterator) Next() interface{} {
	if !s.IsDone() {
		next := s.next
		s.next++
		return next
	}
	return nil
}
