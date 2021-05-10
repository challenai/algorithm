package solution

func reorderSpaces(text string) string {
  rsu := ""
  // we can use split() func, but why dont use a raw travel and make it a bit more challenge
  spaceCounter := 0
  sz := 0
  wordStarts := false
  // add a sential
  text += " "
  for i := 0; i < len(text); i++ {
    if text[i] == ' ' {
      spaceCounter++
      if wordStarts {
        sz++
        wordStarts = false
      }
    } else {
      wordStarts = true
    }
  }
  spaceCounter--
  if sz == 0 {
    return getSpace(spaceCounter)
  }
  var margin, end string
  if sz == 1 {
    margin = ""
    end = getSpace(spaceCounter)
  } else {
    margin = getSpace(spaceCounter / (sz-1))
    end = getSpace(spaceCounter % (sz-1))
  }
  start := 0
  for i := 0; i < len(text); i++ {
    if text[i] == ' ' && text[start] != ' ' {
      if sz == 1 {
        rsu += text[start:i]
      } else {
        rsu += text[start:i] + margin
      }
      sz--
      start = i
      continue
    }
    if text[start] == ' ' {
      start = i
    }
  }
  rsu += end
  return rsu
}

func getSpace(n int) string {
  spaces := ""
  for i := 0; i < n; i++ {
    spaces += " "
  }
  return spaces
}
