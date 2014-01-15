-- prob 11

function oneway()
	io.input("prob11.data")
	local data = io.read("*all")
	return data
end

for line in io.lines("prob11.data") do
	print(line)
end
