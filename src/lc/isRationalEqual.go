package lc

import (
	"fmt"
	"strings"
)

func IsRationalEqual(S string, T string) bool {
	s_int := get_int(S)
	t_int := get_int(T)
	s_dot := get_dot(S)
	t_dot := get_dot(T)
	if is_9_dot(S) {
		s_dot += 1
		if s_dot > 99999999 {
			s_dot = 0
			s_int += 1
		}
	}
	if is_9_dot(T) {
		t_dot += 1
		if t_dot > 99999999 {
			t_dot = 0
			t_int += 1
		}
	}
	return s_int == t_int && s_dot == t_dot



}

func get_int(s string) int {
	n := len(s)
	ret := 0
	for i := 0; i < n && s[i] != '.'; i++ {
		ret = ret * 10 + int(s[i]) - '0'
	}
	return ret

}

func get_dot(s string) int {
	n := len(s)
	k := 0

	if !strings.ContainsAny(s, ".") {
		return 0
	}
	for s[k] !='.' && k < n {
		k++
	}
	if k == n || k == n - 1 {
		return 0
	}
	k++

	temp := [8]int{}
	for j := 0; j < 8; j++ {
		temp[j] = 48
	}


	j := 0
	fmt.Println(k)
	if strings.ContainsAny(s, "(") {
		for k < n && s[k] != '(' {
			temp[j] = int(s[k])
			j++
			k++
		}
		if s[k] == '(' {
			index := k + 1
			for j < 8 {
				temp[j] = int(s[index])
				j++
				index++
				if index == n - 1 {
					index = k + 1
				}
			}
		}
		fmt.Println(temp)
	} else {
		for k < n {
			temp[j] = int(s[k])
			j++
			k++
		}
	}

	m := len(temp)
	ret := 0
	for i := 0; i < m && temp[i] != '.'; i++ {
		ret = ret * 10 + (temp[i]) - '0'
	}
	//fmt.Println(ret)
	return ret

}

func is_9_dot(s string) bool {
	n := len(s)
	if s[n - 1] != ')' {
		return false
	}

	for i := n - 2; s[i] !=  '('; i-- {
		if s[i] != '9' {
			return false
		}
	}
	return true
}