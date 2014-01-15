-- prob 5

function mysolution()
	-- my solution is super slow
	-- 34 seconds at least
	local cc = 2

	function check(n)
		for i=1,20 do
			if (n % i) ~= 0 then
				return false
			end
		end
		return true
	end

	while true do
		local cv = check(cc)
		if cv then
			print("found a value that matched: " .. cc)
			break
		end
		cc = cc + 2
		if cc>500000000 then
			print("giving up")
			break
		end
	end

end

function onemore()
	local counter = 2120
	while true do
		if counter%11==0 and counter%12==0 and 
			counter%13==0 and counter%14==0 and 
			counter%15==0 and counter%16 == 0 and
			counter%17 == 0 and counter%18==0 and 
			counter%19 == 0 and counter%20 == 0 then
			print("counter: " .. counter)
			break
		end
		counter = counter + 2

		if counter > 500000000 then
			break
		end
	end
end


-- onemore()
mysolution()

