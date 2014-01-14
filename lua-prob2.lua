-- prob 2

local fibs = {}

function pt(t)
	for k,v in pairs(t) do
		print(k .. " " .. v)
	end
end

function gen()
	fibs[1] = 1
	fibs[2] = 2
	local index = 3
	while true do
		local next = fibs[index-1] + fibs[index-2]
		if next > 4000000 then
			break
		end
		fibs[index] = next
		index = index + 1
	end
end

function sum_it(t)
	local sum = 0
	for i=1,#t do
		if t[i]%2 == 0 then
			sum = sum + t[i]
		end
	end
	print("sum: " .. sum)
end

gen()
-- pt(fibs)
sum_it(fibs)
