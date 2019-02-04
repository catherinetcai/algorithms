# Bubble Sort
# Take an array of numbers and continually
# swap two adjacent members until the array is fully sorted
# Ex: [3,2,9,5]
# 1: [2,3,9,5] - We compared 3 and 2 and swapped them
# 2: [2,3,9,5] - We compared 3 and 9 and didn't swap
# 3: [2,3,5,9] - We compared 9 and 5 and swapped them
# 4: We actually have to loop through the entire array one more time to make sure we didn't swap anything

def sort(input)
  # If the length of the array is < or = 1, then we already
  # know it's sorted, so exit
  return input if input.length <= 1

  # Now we want to run this algorithm in a loop for forever
  # until we know that the entire array is sorted. That
  # will be our "break" condition - the condition to
  # break out of the loop.

  swapped = true
  while swapped do
    # We set swapped equal to false here, because we
    # basically want to set it only to true when we've
    # actually swapped something. If we never swap anything
    # in this iteration of the loop, it means that the
    # array is sorted and we're done.
    swapped = false

    # We use upto, which is a function from Enumerators
    # https://ruby-doc.org/core-2.2.0/Integer.html#method-i-upto
    # We basically are saying we are going to access through the
    # array we're given, up until the length of it -2. Why?
    # We're accessing each of the array elements by index.
    # Example: [1,2,3,4]
    # If we want to get 1 from the array, we access element 0.
    # In this example, we want to access the current index and the
    # index ahead of it. So, in this loop, we only want to try
    # and access the length of the array minus 2, because otherwise
    # we'll run out of the bounds of the array.
    0.upto(input.length - 2) do |index|
      # If the number at our current index is greater than the one
      # in front of it... swap!
      if input[index] > input[index+1]
        # Here's a neat trick you can do so you don't have to assign
        # any variables to hold onto a value.
        input[index], input[index+1] = input[index+1], input[index]
        swapped = true
      end
    end
  end

  input
end

sort([1,4,3,5])
