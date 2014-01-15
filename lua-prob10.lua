-- prob 10

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

local total = 0
for i=1,2000000 do
	if is_prime(i) then
		total = total + i
	end
end

print(total)