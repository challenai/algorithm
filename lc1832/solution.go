package solution

func checkIfPangram(sentence string) bool {
  hash := make(map[byte]bool, 26)
  for i := 0; i < len(sentence); i++ {
    hash[sentence[i]] = false
  }
  return len(hash) == 26
}
