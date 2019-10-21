//对形如map[string]struct结构进行排序，排序字段为struct中的某一个int字段
package sortutil

type MapStringKeyIntValue struct {
	Key         string
	SortedValue int
}

type MapStringKeyIntValueSorter []MapStringKeyIntValue

func (m MapStringKeyIntValueSorter) Len() int {
	return len(m)
}

func (m MapStringKeyIntValueSorter) Less(i, j int) bool {
	return m[i].SortedValue < m[j].SortedValue
}

func (m MapStringKeyIntValueSorter) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
