package main

var S = []string{"I", "fell", "in", "love", "with", "you"}

func splitStringSlice(s []string) ([]string, []string){
	mid := len(s)/2
	return s[:mid], s[mid:]
}

// make the function dynamic
func splitAnySlice[T any](s []T) ([]T, []T){
	mid := len(s)/2
	return s[:mid], s[mid:]
}