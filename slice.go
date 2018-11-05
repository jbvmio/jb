package jb

import (
	"bytes"
	"sort"
	"strings"
)

// SortSlice here
func SortSlice(sl interface{}) {
	switch sl := sl.(type) {
	case []string:
		sort.Slice(sl, func(i, j int) bool {
			return sl[i] < sl[j]
		})
	case []int:
		sort.Slice(sl, func(i, j int) bool {
			return sl[i] < sl[j]
		})
	}
}

// FilterUnique here
func FilterUnique(strSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func parseStdin(b []byte) (string, []string) {
	bits := bytes.TrimSpace(b)
	lines := string(bits)

	var args []string
	a := strings.Split(lines, "\n")
	kindIn := strings.Fields(strings.TrimSpace(a[0]))[0]

	for _, b := range a[1:] {
		b := strings.TrimSpace(b)
		c := cutField(b, 1)
		args = append(args, c)
	}

	return kindIn, args
}

func returnHeaders(b []byte) []string {
	bits := bytes.TrimSpace(b)
	lines := string(bits)

	a := strings.Split(lines, "\n")

	h := strings.TrimSpace(a[0])
	return strings.Fields(h)
}

func columnReturn(b []byte, n int) []string {
	bits := bytes.TrimSpace(b)
	lines := string(bits)

	var col []string
	a := strings.Split(lines, "\n")

	for _, b := range a {
		b := strings.TrimSpace(b)
		c := cutField(b, n)
		col = append(col, c)
	}
	return col
}

func cutField(s string, f int) string {
	d := f - 1
	fields := strings.Fields(s)
	if len(fields) < f {
		d = len(fields) - 1
	}
	return fields[d]
}

// struct sorting example:
/*
func sortSlice(sl interface{}) {
	switch sl := sl.(type) {
	case []string:
		sort.Slice(sl, func(i, j int) bool {
			return sl[i] < sl[j]
		})
	case []Pod:
		sort.Slice(sl, func(i, j int) bool {
			return sl[i].PodName < sl[j].PodName
		})
	case []Node:
		sort.Slice(sl, func(i, j int) bool {
			return sl[i].NodeName < sl[j].NodeName
		})
	case []ReplicaSet:
		sort.Slice(sl, func(i, j int) bool {
			return sl[i].ReplicaSet < sl[j].ReplicaSet
		})
	case []Deployment:
		sort.Slice(sl, func(i, j int) bool {
			return sl[i].Deployment < sl[j].Deployment
		})
	}
}
*/
