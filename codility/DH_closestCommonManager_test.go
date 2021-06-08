package codility

import (
	"testing"
)

func TestClosestCommonManager(t *testing.T) {
	t.Run("basic case", func(t *testing.T) {
		e1 := &Employee{Name: "emp1"}
		e2 := &Employee{Name: "emp2"}
		ceo := &Employee{Name: "ceo", Reports: []*Employee{e1, e2}}

		var path1, path2 []*Employee

		employeePath(ceo, e1, &path1)
		employeePath(ceo, e2, &path2)

		expected := ceo.Name
		got := closestCommonManager(path1, path2)

		if expected != got {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})

	t.Run("more employees", func(t *testing.T) {
		e1 := &Employee{Name: "Emp1"}
		e2 := &Employee{Name: "Emp2"}
		m1 := &Employee{Name: "Manager1", Reports: []*Employee{e1}}
		m2 := &Employee{Name: "Manager2", Reports: []*Employee{e2}}
		mm := &Employee{Name: "ManagerOfManagers", Reports: []*Employee{m1, m2}}
		ceo := &Employee{Name: "CEO", Reports: []*Employee{mm}}

		var path1, path2 []*Employee
		employeePath(ceo, e1, &path1)
		employeePath(ceo, e2, &path2)

		expected := mm.Name
		got := closestCommonManager(path1, path2)

		if expected != got {
			t.Errorf("got != expected; %s != %s", got, expected)
		}
	})
}
