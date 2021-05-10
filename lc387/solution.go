package solution

func firstUniqChar(s string) int {
  rsu := -1
  hash := make([]int, 26)
  for i := 0; i < 26; i++ {
    hash[i] = -1
  }
  for i := 0; i < len(s); i++ {
    if hash[s[i]-97] > -1 {
      hash[s[i]-97] = -2
      continue
    }
    if hash[s[i]-97] == -1 {
      hash[s[i]-97] = i
    }
  }
  for i := 0; i < 26; i++ {
    if hash[i] > -1 {
      if rsu == -1 {
        rsu = hash[i]
        continue
      }
      if rsu > hash[i] {
        rsu = hash[i]
      }
    }
  }
  return rsu
}
