package main

import "fmt"

type Key [26]byte

func strKey(str string) (key Key) {
	for i := range str {
		key[str[i]-'a']++
	}
	return
}
func groupAnagram(strs []string) {
	groups := make(map[Key][]string)

	for _, v := range strs {
		key := strKey(v)
		fmt.Println("key :: after inserting v: ",key)
		groups[key] = append(groups[key], v)
		fmt.Println("group after insetion::;",groups)
		break
	}

	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

func main() {
	strs := []string{"eat", "tea", "tan", "nat"}
	var key Key
	fmt.Println("Key:::::::",key)
	groupAnagram(strs)
}