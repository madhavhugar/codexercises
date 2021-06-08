package codility

/*
 * Delivery Hero
 * Initech is a company which has CEO Bill and a hierarchy of employees.
 * All employees have a name and an id; employees can also have a list of other employees reporting to them,
 * which can themselves have reports, and so on. An employee with at least one report is called a manager.
 *
 * Please implement the `closestCommonManager()` method to find the closest manager for two given employees
 * (i.e. the manager farthest from the CEO that both employees report up to).

 */

type Employee struct {
	Name    string
	Reports []*Employee
}

func employeePath(root *Employee, target *Employee, path *[]*Employee) bool {
	*path = append(*path, root)
	if root == target {
		return true
	}

	for _, r := range root.Reports {
		if res := employeePath(r, target, path); res {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
}

func closestCommonManager(path1 []*Employee, path2 []*Employee) string {
	m := len(path1)
	n := len(path2)

	var start int
	if m > n {
		start = n
	} else {
		start = m
	}

	for i := start - 1; i >= 0; i-- {
		if path1[i] == path2[i] {
			return path1[i].Name
		}
	}

	return path1[0].Name
}
