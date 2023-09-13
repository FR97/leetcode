package easy

import "sort"

func longestCommonPrefix(strs []string) string {
	sort.Strings(strs)

	for i := 0; i < len(strs[0]); i++ {
		first := strs[0][i]
		last := strs[len(strs)-1][i]
		// First and last string would be most different
		if first != last {
			return strs[0][:i]
		}
	}

	return strs[0]
}
