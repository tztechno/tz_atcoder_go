//abc091_a.rb
################################
################################
################################
################################
################################
################################
a, b, c = gets.split.map(&:to_i)
puts a + b >= c ? 'Yes' : 'No'
################################
i=gets.split(" ").map{|x|x.to_i}
if i[0]+i[1] < i[2] then
  puts("No")
else
  puts("Yes")
end
################################
arr = gets.split
A = arr[0].to_i
B = arr[1].to_i
C = arr[2].to_i

if A + B >= C then
  print 'Yes'
else 
  print 'No'
end
################################
A, B, C = gets.chomp.split.map(&:to_i)
if A+B>=C
puts "Yes"
else
puts "No"
end
################################
