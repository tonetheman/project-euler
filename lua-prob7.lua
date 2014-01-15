-- prob 7

function is_prime(n)
	if n == 1 then
		return false
	end
	if n == 2 then 
		return true 
	end
	if n%2==0 then
		return false
	end
	local check_to = math.sqrt(n)
	local _res = false
	for i=3,check_to,2 do
		local val = n % i
		if val == 0 then
			return false
		end
	end
	return true
end

local p = {}
p[1] = 2
local counter = 2
for i=3,200000,2 do
	if is_prime(i) then
		p[counter] = i
		if counter == 10001 then
			break
		end
		counter = counter + 1
	end
end

local utils = require("utils")
utils.pt(p)