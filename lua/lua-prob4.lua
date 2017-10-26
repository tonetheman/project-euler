-- prob 4

local largest = -1

function is_pal(n)
	local s = tostring(n)
	if s == s:reverse() then
		return true
	end
	return false
end

for i = 100,999 do
	for j = 100, 999 do
		local n = i * j
		if is_pal(n) then
			print(i .. " " .. j .. " " .. n)
			if n > largest then
				largest = n
			end
		end
	end
end

print("largest value is " .. largest)