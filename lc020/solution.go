package solution

func isValid(s string) bool {
  st := []byte{}
  for i := 0; i < len(s); i++ {
    switch s[i] {
    case '(':
      st = append(st, ')')
    case '{':
      st = append(st, '}')
    case '[':
      st = append(st, ']')
    default:
      if len(st) == 0 || s[i] != st[len(st)-1] {
        return false
      }
      st = st[:len(st)-1]
    }
  }
  return len(st) == 0
}
