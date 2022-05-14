package golang_helpers

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

func GetUniqueID() string {
	hasher := md5.New()
	hasher.Write([]byte(GetTimeStamp())) //using timestamp as seed - check dates.go for this function
	slugval := hex.EncodeToString(hasher.Sum(nil))
	slugvalconv := ShuffleSlice(strings.Fields(slugval[0:10])) //function below
	joinedslugs := strings.Join(slugvalconv, "-")
	joinedslugs = strings.Replace(joinedslugs, "-", "", 50)

	return joinedslugs
}

func GetUniqueInts(int_slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range int_slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func StringSliceDiff(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func ShuffleSlice(src []string) []string {
	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}
	return final
}

func SliceContainsString(str string, strarr []string) bool {
	for _, v := range strarr {
		if v == str {
			return true
		}
	}
	return false
}
