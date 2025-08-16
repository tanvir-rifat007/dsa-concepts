#include <stdio.h>
#include <stdlib.h>

int climbStairs(int n) {
  if (n <= 1)
    return 1;

  int a = 1, b = 1, temp;
  for (int i = 2; i <= n; i++) {
    temp = a + b;
    a = b;
    b = temp;
  }
  return b;
}

int main() {
  int res = climbStairs(3);

  printf("%d\n", res);

  return 0;
}
