package strset

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	testCases := []struct {
		s1 Set
		out string
		s2 Set
	}{
		{Make(), "a", Make()},
		{Make("a"), "a",  Make()},
		{Make("a"), "b", Make("a")},
		{Make("a", "b"), "a", Make("b")},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v remove %v remain %v", tc.s1, tc.out, tc.s2),
			func(t *testing.T) {
				tc.s1.Remove(tc.out)
				assert.True(t, tc.s1.Equal(tc.s2))
			})
	}
}

func TestRemoveAll(t *testing.T) {
	for _, tc := range differenceTestCases {
		t.Run(fmt.Sprintf("%v.RemoveAll(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.RemoveAll(tc.set2.Elems()...)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPop(t *testing.T) {
	testCases := []struct {
		s1 Set
		wantElem string
		wantFound bool
		s2 Set

	}{
		{Make(), "", false, Make()},
		{Make("a"), "a", true, Make()},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v pop, got (%v, %v) remain %v",
			tc.s1, tc.wantElem, tc.wantFound, tc.s2),
			func(t *testing.T) {
				elem, found := tc.s1.Pop()
				assert.Equal(t, tc.wantElem, elem)
				assert.Equal(t, tc.wantFound, found)
				assert.True(t, tc.s1.Equal(tc.s2))
			})
	}
}

func TestPop_3(t *testing.T) {
	s := Make("a", "b", "c")
	for wantLen:=2; wantLen>=0; wantLen-- {
		elem, found := s.Pop()
		assert.Equal(t,wantLen, s.Len())
		assert.NotEqual(t, "", elem)
		assert.True(t, found)
	}
	elem, found := s.Pop()
	assert.Equal(t, 0, s.Len())
	assert.Equal(t, "", elem)
	assert.False(t, found)

}

func TestClear(t *testing.T) {
	testCases := []Set {
		Make(),
		Make("a"),
		Make("a", "b"),
	}
	for _, set := range testCases {
		t.Run(fmt.Sprintf("%v.Clear()", set), func(t *testing.T) {
			set.Clear()
			assert.Equal(t, set.Len(), 0)
		})
	}
}

func TestIntersectionUpdate(t *testing.T) {
	for _, tc := range intersectionTestCases {
		t.Run(fmt.Sprintf("%v.IntersectionUpdate(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.IntersectionUpdate(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestUnionUpdate(t *testing.T) {
	for _, tc := range unionTestCases {
		t.Run(fmt.Sprintf("%v.UnionUpdate(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.UnionUpdate(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDifferenceUpdate(t *testing.T) {
	for _, tc := range differenceTestCases {
		t.Run(fmt.Sprintf("%v.DifferenceUpdate(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.DifferenceUpdate(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSymmetricDifferenceUpdate(t *testing.T) {
	for _, tc := range symmetricDifferenceTestCases {
		t.Run(fmt.Sprintf("%v.SymmetricDifferenceUpdate(%v) is %v", tc.set1, tc.set2, tc.want), func(t *testing.T) {
			got := tc.set1.Copy()
			got.SymmetricDifferenceUpdate(tc.set2)
			assert.Equal(t, tc.want, got)
		})
	}
}
