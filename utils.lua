

local me = {}

function me.pt(t)
	for k,v in pairs(t) do
		if type(v) == "table" then
				for kk,vv in pairs(v) do
					print(string.format("%d %d", kk,vv))
				end
				print()
		else
			print(k .. " " .. v)
		end
	end
end

function me.is_prime(n)
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


return me
