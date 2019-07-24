package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) { //这里的for相当于while
		s.words = append(s.words, 0)
		//fmt.Println(word)
	}
	s.words[word] |= 1 << bit //这里左移一位是为了分辨有没有存在这个数，0左移还是0

}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func (s *IntSet) Len() int {
	var flag int
	for _, word := range s.words {
		if word != 0 {
			flag++
		}

	}

	return flag
} // return the number of elements
func (s *IntSet) Remove(x int) {
	word := x / 64
	ok := s.words[word]
	if ok != 0 {
		s.words[word] = 0
	}
} // remove x from the set

func (s *IntSet) Clear() {
	s.words = s.words[:0]
} // remove all elements from the set

func (s *IntSet) Copy() *IntSet {
	newword := make([]uint64, len(s.words))
	newIntset := new(IntSet)
	newIntset.words = newword[:len(s.words)]
	return newIntset
} // return a copy of the set

//定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。
func (s *IntSet) AddAll(a ...int) {
	if len(a) != 0 {
		for _, value := range a {
			s.Add(value)
		}

	}

}

//(*IntSet).UnionWith会用|操作符计算两个集合的交集，我们再为IntSet实现另外的几个函数IntersectWith
// (交集：元素在A集合B集合均出现),DifferenceWith(差集：元素出现在A集合，
// 未出现在B集合),SymmetricDifference(并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A)。

func (s *IntSet) IntersectWith(ss *IntSet) *IntSet {
	var w IntSet
	for word, bit := range s.words {
		if bit == 0 {
			continue
		}

		has_x := int(word) * 64
		for j := 0; j < 64; j++ {

			if ((bit & (1 << uint(j))) != 0) && (ss.Has(has_x + j)) {
				//fmt.Println(has_x+j)
				w.Add(has_x + j)
			}

		}
	}
	return &w
}
func (s *IntSet) DifferenceWith(ss *IntSet) *IntSet {
	var w IntSet
	for word, bit := range s.words {
		if bit == 0 {
			continue
		}
		//fmt.Println("bit:",bit)
		has_x := int(word) * 64
		for j := 0; j < 64; j++ {

			if ((bit & (1 << uint(j))) != 0) && !ss.Has(has_x+j) {
				w.Add(has_x + j)
			}

		}

	}
	return &w
}

func (s *IntSet) SymmetricDifference(ss *IntSet) *IntSet {
	var w IntSet
	for word, bit := range s.words {
		if bit == 0 {
			continue
		}
		has_x := int(word) * 64
		for j := 0; j < 64; j++ {

			if ((bit & (1 << uint(j))) != 0) && (!ss.Has(has_x + j)) {
				//fmt.Println(has_x+j)
				w.Add(has_x + j)
			}
		}
	}
	for word, bit := range ss.words {
		if bit == 0 {
			continue
		}
		//fmt.Println("bit",bit)
		has_x := int(word) * 64
		for j := 0; j < 64; j++ {

			if (bit&(1<<uint(j)) != 0) && (!s.Has(has_x + j)) {
				//fmt.Println(has_x+j)
				w.Add(has_x + j)
			}

		}
	}
	return &w
}
func (s *IntSet) Elem() []int {
	var haha []int
	for _, bit := range s.words {
		if bit != 0 {
			for j := 0; j < 64; j++ {
				if bit&(1<<uint(j)) != 0 {
					haha = append(haha, int(j))
				}

			}
		}
	}
	return haha
}

// 我们这章定义的IntSet里的每个字都是用的uint64类型，但是64位的数值可能在32位的平台上不高效。
// 修改程序，使其使用uint类型，这种类型对于32位平台来说更合适。当然了，这里我们可以不用简单粗暴地除64，
// 可以定义一个常量来决定是用32还是64，这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
func main() {
	IntSet1 := new(IntSet)
	IntSet2 := new(IntSet)
	IntSet1.Add(1)

	IntSet1.Add(2)

	IntSet1.Add(3)

	IntSet2.Add(3)

	IntSet2.Add(2)

	IntSet2.Add(4)

	//IntSet1.UnionWith(IntSet2)
	fmt.Print("ss")
	fmt.Println(IntSet1.IntersectWith(IntSet2).String())
	fmt.Println(IntSet1.DifferenceWith(IntSet2).String())
	fmt.Println(IntSet2.String())
	fmt.Println(IntSet1.SymmetricDifference(IntSet2).String())
	fmt.Println(IntSet1.Elem())
}
