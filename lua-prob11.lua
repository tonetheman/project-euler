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

local gd = {}
function gd.set(self,a,b,c,d)
	self.a = a
	self.b = b
	self.c = c
	self.d = d
end
local guess = 0
function pg()
	print("guess: " .. guess)
	print(string.format("%d %d %d %d", gd.a, gd.b, gd.c, gd.d))
end

for i=1,20 do
	for j=1,17 do

		-- check the row from left to right
		local tmp = data[i][j] * data[i][j+1] * data[i][j+2] * data[i][j+3]
		if tmp > guess then
			guess = tmp
			gd:set(data[i][j],data[i][j+1],data[i][j+2],data[i][j+3])
		end

		tmp = data[j][i] * data[j+1][i] * data[j+2][i] * data[j+3][i]
		if tmp > guess then
			guess = tmp
			gd:set(data[j][i],data[j][i+1],data[j][i+2],data[j][i+3])
		end
	end
end

pg()