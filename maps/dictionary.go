package maps

type Dictionary map[string]string

// func Search(d Dictionary, word string) string {
// 	return d[word]
// }

func (d Dictionary) Search(word string) string {
	return d[word]
}
