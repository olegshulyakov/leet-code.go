package main

import (
	"testing"
)

/*
There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has, and an integer extraCandies, denoting the number of extra candies that you have.

Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies, they will have the greatest number of candies among all the kids, or false otherwise.

Note that multiple kids can have the greatest number of candies.

Example 1:
Input: candies = [2,3,5,1,3], extraCandies = 3
Output: [true,true,true,false,true]
Explanation: If you give all extraCandies to:
- Kid 1, they will have 2 + 3 = 5 candies, which is the greatest among the kids.
- Kid 2, they will have 3 + 3 = 6 candies, which is the greatest among the kids.
- Kid 3, they will have 5 + 3 = 8 candies, which is the greatest among the kids.
- Kid 4, they will have 1 + 3 = 4 candies, which is not the greatest among the kids.
- Kid 5, they will have 3 + 3 = 6 candies, which is the greatest among the kids.

Example 2:
Input: candies = [4,2,1,1,2], extraCandies = 1
Output: [true,false,false,false,false]
Explanation: There is only 1 extra candy.
Kid 1 will always have the greatest number of candies, even if a different kid is given the extra candy.

Example 3:
Input: candies = [12,1,12], extraCandies = 10
Output: [true,false,true]


Constraints:
n == candies.length
2 <= n <= 100
1 <= candies[i] <= 100
1 <= extraCandies <= 50
*/

// kidsWithCandies determines which kids can have the greatest number of candies
// after being given extraCandies.
// For each kid, it checks if adding extraCandies to their current candies
// would make them have the maximum number of candies.
//
// Parameters:
//
//	candies []int - Array representing the number of candies each kid currently has
//	extraCandies int - Number of extra candies to distribute
//
// Returns:
//
//	[]bool - Boolean array where true indicates the kid can have the most candies
func kidsWithCandies(candies []int, extraCandies int) []bool {
	// Find the maximum number of candies any kid currently has
	maximum := 0
	for i := range candies {
		if candies[i] > maximum {
			maximum = candies[i]
		}
	}

	// Create result array by checking if each kid can have maximum candies
	// after receiving extraCandies
	res := make([]bool, len(candies))
	for i, c := range candies {
		res[i] = c+extraCandies >= maximum
	}

	return res
}

func TestKidsWithCandies(t *testing.T) {
	testCases := []struct {
		candies      []int
		extraCandies int
		want         []bool
	}{
		{candies: []int{2, 3, 5, 1, 3}, extraCandies: 3, want: []bool{true, true, true, false, true}},
		{candies: []int{4, 2, 1, 1, 2}, extraCandies: 1, want: []bool{true, false, false, false, false}},
		{candies: []int{12, 1, 12}, extraCandies: 10, want: []bool{true, false, true}},
	}
	for _, tc := range testCases {
		t.Run("1431. Kids With the Greatest Number of Candies", func(t *testing.T) {
			out := kidsWithCandies(tc.candies, tc.extraCandies)
			if len(out) != len(tc.want) {
				t.Errorf("len(kidsWithCandies() = %v, want: %v", len(out), len(tc.want))
			}
			for i := range out {
				if out[i] != tc.want[i] {
					t.Errorf("kidsWithCandies()[%v] = %v, want: %v", i+1, out[i], tc.want[i])
				}
			}
		})
	}
}
