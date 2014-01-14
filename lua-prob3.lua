-- prob3

local target = 600851475143

function pt(t)
	for k,v in pairs(t) do
		print(k .. " " .. v)
	end
end

function factor(t)
	local top = math.sqrt(target)
	print("top number to check: " .. top)
	local nums = {}
	local index  = 1
	for i=1,top do
		if t%i==0 then
			nums[index] = i
			index = index + 1
		end
	end
	pt(nums)
end

factor(target)
