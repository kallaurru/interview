package algo_228

import "fmt"

func SummaryRanges(in []int) []string {
	if in == nil || len(in) == 0 {
		return []string{""}
	}

	ff := func(a, b int) string {
		if a == b {
			return fmt.Sprintf("%d", a)
		}
		return fmt.Sprintf("%d->%d", a, b)
	}

	begin := in[0]
	end := begin
	out := make([]string, 0, len(in))

	for i := 1; i < len(in); i++ {
		if in[i]-1 == in[i-1] {
			// условие добавления в ряд
			end = in[i]
			continue
		}
		out = append(out, ff(begin, end))
		begin = in[i]
		end = begin
	}
	out = append(out, ff(begin, end))

	return out
}
