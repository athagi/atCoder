a,b = map(int,input().split())

res = "Odd"
if a* b %2 == 0:
  res = "Even"

print(res)