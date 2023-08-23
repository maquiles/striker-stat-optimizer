package main

func Contains(x string, list []string) bool {
	for _, l := range list {
		if l == x {
			return true
		}
	}

	return false
}
