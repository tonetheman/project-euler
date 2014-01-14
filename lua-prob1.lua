-- problem 1

local i
local sum = 0
for i=1,999 do
	if (i % 5 == 0) or (i%3 ==0) then
		print("matches: " .. i)
		sum = sum + i
	else
		print("NOPE: " .. i)
	end
end

print("sum: " .. sum)


