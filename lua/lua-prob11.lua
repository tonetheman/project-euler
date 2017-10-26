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

for i=1,20 do
	for j=1,20 do
		io.write(string.format("%2d ", data[i][j]))
	end
	print()
end

local gd = {}
function gd.set(self,a,b,c,d)
	self.a = a
	self.b = b
	self.c = c
	self.d = d
end
local guess = 0
local guess_i = 0
local guess_j = 0
local guess_direction = "NONE"
function pg()
	print("guess: " .. guess)
	print("i,j : " .. guess_i .. " " .. guess_j)
	print("dir : " .. guess_direction)
	print(string.format("%d %d %d %d", gd.a, gd.b, gd.c, gd.d))
end

for i=1,20 do
	for j=1,17 do

		-- check the row from left to right
		local tmp = data[i][j] * data[i][j+1] * data[i][j+2] * data[i][j+3]
		if tmp > guess then
			guess = tmp
			guess_i = i
			guess_j = j
			guess_direction = "L_TO_R"
			gd:set(data[i][j],data[i][j+1],data[i][j+2],data[i][j+3])
		end

		tmp = data[j][i] * data[j+1][i] * data[j+2][i] * data[j+3][i]
		if tmp > guess then
			guess = tmp
			guess_i = i
			guess_j = j
			guess_direction = "U_TO_D"
			gd:set(data[j][i],data[j][i+1],data[j][i+2],data[j][i+3])
		end
	end
end

-- diagonal down
for i=1,17 do
	for j=1,17 do
		local tmp = data[i][j]*data[i+1][j+1]*data[i+2][j+2]*data[i+3][j+3]
		if tmp>guess then
			guess = tmp
			guess_i = i
			guess_j = j
			guess_direction = "DD"
			gd:set(data[i][j],data[i+1][j+1],data[i+2][j+2],data[i+3][j+3])
		end
	end
end

pg()



