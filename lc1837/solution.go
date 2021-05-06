package solution

func sumBase(n int, k int) int {
  // I dont know the best way to convert the base
  // so Lets just convert it by divide
  // WTF, I write a wrong answer here
  // sum := 0
  // for n / k > 0 {
  //   sum += n / k
  //   n %= k
  // }
  // sum += n

  sum := 0
  for n > 0 {
    sum += n%k
    n /= k
  }
  return sum
}
