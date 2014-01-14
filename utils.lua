

local me = {}

function pt(t)
	for k,v in pairs(t) do
		print(k .. " " .. v)
	end
end


me.pt = pt


return me
