 p File.read('inputs/day1.txt').split("\n").map(&:to_i).combination(2).select {|x| x.sum==2020}.first.reduce(&:*)
 p File.read('inputs/day1.txt').split("\n").map(&:to_i).combination(3).select {|x| x.sum==2020}.first.reduce(&:*)
