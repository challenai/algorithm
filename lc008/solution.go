package solution

func myAtoi(s string) int {
  s = strings.Trim(s, " ")
  if len(s) == 0 {
    return 0
  }
  s_ := ""
  var symbol byte
  if s[0] == '-' || s[0] == '+' {
    symbol = s[0]
    s = s[1:]
  }
  for _, char := range s {
    if char >= '0' && char <= '9' {
      if char == '0' && s_ == "" {
        continue
      }
      s_ += string(char)
    } else {
      break
    }
  }
  if s_ == "" {
    return 0
  }
  if symbol != '-' && isOverflow(s_, "2147483647") {
    return 2147483647
  }
  if symbol == '-' && isOverflow(s_, "2147483648") {
    return -2147483648
  }
  rsu, _ := strconv.Atoi(s_)
  if symbol == '-' {
    return -rsu
  }
  return rsu
}

// validate overflow, if overflow, return false
func isOverflow(s, border string) bool {
  if len(s) > len(border) {
    return true
  }
  if len(s) < len(border) {
    return false
  }
  return s > border
}
