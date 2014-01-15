-- prob 11

local utils = require("utils")

function oneway()
	io.input("prob11.data")
	local data = io.read("*all")
	return data
end

local data = {}

for line in io.lines("prob11.data") do
	local tmp_data = {}
	for num in string.gmatch(line, "%d+") do
		table.insert(tmp_data, tonumber(num))
	end
	table.insert(data, tmp_data)
end


