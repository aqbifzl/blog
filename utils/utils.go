package utils

func IsFilenameSafe(s string) bool {
	for _, r := range s {
		if !((r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '_') {
			return false
		}
	}

	return true
}

func GetPageList(page, pages int) []int {
	if pages <= 3 {
		out := make([]int, pages)
		for i := 1; i <= pages; i++ {
			out[i-1] = i
		}
		return out
	}

	start := min(max(page-1, 1), pages-2)

	return []int{start, start + 1, start + 2}
}
