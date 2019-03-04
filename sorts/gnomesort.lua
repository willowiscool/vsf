function sort(a)
	local i, j = 2, 3
 
	while i <= #a do
		if a[i-1] <= a[i] then
			i = j
			j = j + 1
		else
			a[i-1], a[i] = a[i], a[i-1] -- swap
			show(a)
			i = i - 1
			if i == 1 then -- 1 instead of 0
				i = j
				j = j + 1
			end
		end
	end
end
