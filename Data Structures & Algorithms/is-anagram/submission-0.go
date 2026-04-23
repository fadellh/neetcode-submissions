func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    hmap := make(map[rune]int)

    for _, v := range s {
        hmap[v]++
    }

    for _, v := range t {
        val, exist := hmap[v]
        if exist && val > 0 {
            hmap[v]--
            continue
        }

        return false
    }

    return true
}

