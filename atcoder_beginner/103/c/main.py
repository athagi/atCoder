N = int(input())
A = list(map(int, input().split()))

res = 0
for a in A:
  res += a -1 

print(res)