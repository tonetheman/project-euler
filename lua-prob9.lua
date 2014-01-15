-- prob 9
for i=1,1000 do
	for j=1,1000 do
		for k=1,1000 do
			if i+j+k == 1000 then
				if (i*i)+(j*j) == (k*k) then
					print(i .. " " .. j .. " " .. k)
					print(i*j*k)
					break
				end
			end
		end
	end
end