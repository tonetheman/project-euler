-- prob3

local utils = require("utils")
local target = 600851475143

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
	utils.pt(nums)
end

factor(target)
