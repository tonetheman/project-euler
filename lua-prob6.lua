-- prob 6
local sum1 = 0
local sum2 = 0
for i=1,100 do
	sum1 = sum1 + (i*i)
	sum2 = sum2 + i
end
print(sum1)
sum2 = sum2 * sum2
print(sum2)
print(sum2-sum1)