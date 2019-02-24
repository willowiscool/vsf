function sort(array)
	local j
	for j = 2, #array do
		local key = array[j]
		local i = j - 1
		while i > 0 and array[i] > key do
			array[i + 1] = array[i]
			show(array)
			i = i - 1
		end
		array[i + 1] = key
		show(array)
	end
	return array
end
