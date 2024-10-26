#include "stdio.h"

int add(int x, int y) {
  return x + y;
}

int main() {
  int a = 10;
  int b = 10;
  int sum = add(a, b);
  printf("%d\n", sum);
  return 0;
}

