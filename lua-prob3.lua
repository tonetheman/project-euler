-- prob3

local utils = require("utils")
local target = 600851475143
local debug = false

function is_prime(n)
	if n == 1 or n == 2 then 
		return true 
	end
	if n%2==0 then
		return false
	end
	local check_to = math.sqrt(n)
	if debug then
		print("\tchecking to this val: " .. check_to)
	end
	local _res = false
	for i=3,check_to,2 do
		local val = n % i
		if debug then
			print("\tneed to check this factor: " .. i .. " " .. val)
		end
		if val == 0 then
			return false
		end
	end
	return true
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
	utils.pt(nums)
end

function prime_test()
	for i=1,100 do
		if is_prime(i) then
			print(i .. " PRIME")
		end
	end
end

prime_test()
-- print(is_prime(9))
