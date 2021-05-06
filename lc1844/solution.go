package solution

func replaceDigits(s string) string {
  rsu := ""
  for i := 1; i < len(s); i += 2 {
    rsu += string(s[i-1]) + string(s[i-1]+s[i]-48)
  }
  if s[len(s)-1] <= 'z' && s[len(s)-1] >= 'a' {
    rsu += string(s[len(s)-1])
  }
  return rsu
}
