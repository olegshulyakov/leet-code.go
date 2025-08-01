package main

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

/*

Given an encoded string, return its decoded string.

The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being repeated exactly k times. Note that k is guaranteed to be a positive integer.

You may assume that the input string is always valid; there are no extra white spaces, square brackets are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].

The test cases are generated so that the length of the output will never exceed 10^5.


Example 1:

Input: s = "3[a]2[bc]"
Output: "aaabcbc"


Example 2:

Input: s = "3[a2[c]]"
Output: "accaccacc"


Example 3:

Input: s = "2[abc]3[cd]ef"
Output: "abcabccdcdcdef"


Constraints:

1 <= s.length <= 30
s consists of lowercase English letters, digits, and square brackets '[]'.
s is guaranteed to be a valid input.
All the integers in s are in the range [1, 300].

*/

// decodeString decodes an encoded string according to the pattern k[encoded_string].
func decodeString(s string) string {
	regex := regexp.MustCompile(`(\d+)\[(\w+)\]`)

	for regex.MatchString(s) {
		matches := regex.FindStringSubmatch(s)
		original := matches[0]
		count, _ := strconv.Atoi(matches[1])
		str := matches[2]

		replace := strings.Repeat(str, count)
		s = strings.Replace(s, original, replace, 1)
	}

	return s
}

func TestDecodeString(t *testing.T) {
	testCases := []struct {
		s    string
		want string
	}{
		{"3[a]2[bc]", "aaabcbc"},
		{"3[a2[c]]", "accaccacc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"ef", "ef"},
	}
	for _, tc := range testCases {
		t.Run("394. Decode String", func(t *testing.T) {
			out := decodeString(tc.s)
			if out != tc.want {
				t.Errorf("decodeString(%v) = %v, want: %v", tc.s, out, tc.want)
			}
		})
	}
}
