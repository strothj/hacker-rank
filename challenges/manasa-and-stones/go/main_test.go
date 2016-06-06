package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	for i, test := range testCases {
		r := bytes.NewReader([]byte(test.input))
		w := &bytes.Buffer{}
		solve(r, w)
		actual := w.Bytes()
		expected := []byte(test.expected)
		if !reflect.DeepEqual(expected, actual) {
			t.Fatalf("test(%v): expected=\n|%v| \nactual=\n|%v|", i, expected, actual)
		}
	}
}

var testCases = []struct {
	input    string
	expected string
}{
	{
		input: `2
3
1
2
4
10
100`,
		expected: `2 3 4
30 120 210 300
`,
	},
	{
		input: `5
58
69
24
83
86
81
73
25
25
12
73
82
5
3
23`,
		expected: `1368 1413 1458 1503 1548 1593 1638 1683 1728 1773 1818 1863 1908 1953 1998 2043 2088 2133 2178 2223 2268 2313 2358 2403 2448 2493 2538 2583 2628 2673 2718 2763 2808 2853 2898 2943 2988 3033 3078 3123 3168 3213 3258 3303 3348 3393 3438 3483 3528 3573 3618 3663 3708 3753 3798 3843 3888 3933
6642 6647 6652 6657 6662 6667 6672 6677 6682 6687 6692 6697 6702 6707 6712 6717 6722 6727 6732 6737 6742 6747 6752 6757 6762 6767 6772 6777 6782 6787 6792 6797 6802 6807 6812 6817 6822 6827 6832 6837 6842 6847 6852 6857 6862 6867 6872 6877 6882 6887 6892 6897 6902 6907 6912 6917 6922 6927 6932 6937 6942 6947 6952 6957 6962 6967 6972 6977 6982 6987 6992 6997 7002 7007 7012 7017 7022 7027 7032 7037 7042 7047 7052
1800
803 812 821 830 839 848 857 866 875 884 893 902
12 32 52 72 92
`,
	},
}
