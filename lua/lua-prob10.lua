-- prob 10

local utils = require("utils")


local total = 0
for i=1,2000000 do
	if utils.is_prime(i) then
		total = total + i
	end
end

print(total)
