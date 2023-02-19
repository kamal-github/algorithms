package main

import "strings"

/*
/home/../../sub/./
*/
func simplifyPath(path string) string {
	dirs := strings.Split(path, "/")
	var canonicalPath []string

	for _, d := range dirs {
		if d == "" || d == "." {
			continue
		}

		if d == ".." {
			if len(canonicalPath) > 0 {
				canonicalPath = canonicalPath[:len(canonicalPath)-1]
			}
			continue
		}

		canonicalPath = append(canonicalPath, d)
	}

	return "/" + strings.Join(canonicalPath, "/")
}
