package union

import (
	"fmt"
	"testing"
)

func TestUnion_Union(t *testing.T) {
	u := NewUnion(5)
	u.Union(0, 1)
	u.Union(0, 3)
	u.Union(0, 4)
	u.Union(2, 3)
	fmt.Println(u.parents)
	fmt.Println(u.Find(3))
}
