package solution

func numSpecial(mat [][]int) int {
  rsu := 0
  // define rowSpecial means the element is the only one in a row.
  rowSpecialHash := map[string]bool{}

  // check row
  var specialIdx int
  for i := 0; i < len(mat); i++ {
    specialIdx = -1
    for j := 0; j < len(mat[i]); j++ {
      if mat[i][j] == 1 {
        if specialIdx == -1 {
          specialIdx = j
        } else {
          specialIdx = -1
          break
        }
      }
    }
    if specialIdx > -1 {
      rowSpecialHash[serialize(i, specialIdx)] = false
    }
  }

  // check column
  for i := 0; i < len(mat[0]); i++ {
    specialIdx = -1
    for j := 0; j < len(mat); j++ {
      if mat[j][i] == 1 {
        if specialIdx == -1 {
          specialIdx = j
        } else {
          specialIdx = -1
          break
        }
      }
    }
    if specialIdx > -1 {
      if _, ok := rowSpecialHash[serialize(specialIdx, i)]; ok {
        rsu++
      }
    }
  }
  return rsu
}

func serialize(i, j int) string {
  return strconv.Itoa(i) + "-" + strconv.Itoa(j)
}
