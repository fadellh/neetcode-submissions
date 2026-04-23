func groupAnagrams(strs []string) [][]string {
    alphaMap := map[string][]string{}

    for _,v := range strs {
        sortAlpha := sortedAlpha(v)
        alphaMap[sortAlpha] = append(alphaMap[sortAlpha], v)
    }

    resultArr := [][]string{}

    for _, v := range alphaMap {
        resultArr = append(resultArr, v)
    }

    return resultArr

}


func sortedAlpha(alpha string) string {
    arr := [26]int{}
    for _, v := range alpha {
        arr[v-'a']++
    }

    words := ""
    for i, count := range arr {
        for j := 0; j < count; j++ {
            words += string(rune(i + 'a'))
        }
    }

    return words

}