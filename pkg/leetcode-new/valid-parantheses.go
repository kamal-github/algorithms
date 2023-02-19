package main

func isValid(s string) bool {
	m := map[string]string{
		"]": "[",
		"}": "{",
		")": "(",
	}

	stack := make([]string, 0)

	for _, ch := range s {
		mayBeOpenP, ok := m[string(ch)]
		if !ok {
			stack = append(stack, string(ch))
			continue
		}

		// there is no opening paranethese, so invalid.
		if len(stack) == 0 {
			return false
		}

		openP := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		if openP != mayBeOpenP {
			return false
		}
	}

	return len(stack) == 0
}
