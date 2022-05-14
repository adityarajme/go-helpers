package golang_helpers

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func StringToInt(s string) int {
	intval, _ := strconv.Atoi(s)
	return intval
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func UniqueStringSliceV1(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}

	return us
}

func UniqueStrSliceV2(elements []string) []string {

	pureslice := []string{}
	encountered := map[string]bool{}

	for v := range elements {
		if encountered[elements[v]] == false {
			encountered[elements[v]] = true
			pureslice = append(pureslice, elements[v])
		}
	}

	return pureslice
}

func GetUniqueSlug() string {
	db := DbConn

	slug_query := "SELECT AUTO_INCREMENT FROM information_schema.TABLES WHERE TABLE_NAME LIKE '%any_table_name%'"
	slug_dn, err := DbQueryGetRows(db, slug_query)
	CheckErr(err)

	var slugval string
	var joinedslugs string
	var slugvalconv []string
	for slug_dn.Next() {
		var auto_inc int
		slug_dn.Scan(&auto_inc)
		auto_inc++
		auto_inc_str := strconv.Itoa(auto_inc)

		hasher := md5.New()
		hasher.Write([]byte(auto_inc_str))
		slugval = hex.EncodeToString(hasher.Sum(nil))
		slugvalconv = ShuffleSlice(strings.Fields(slugval[0:10]))
		joinedslugs = strings.Join(slugvalconv, "-")
		joinedslugs = strings.Replace(joinedslugs, "-", "", 50)
	}
	slug_dn.Close()
	return strings.ToLower(joinedslugs)
}

func SlugIt(s string) string {
	var re = regexp.MustCompile("[^a-z0-9]+")
	return strings.Trim(re.ReplaceAllString(strings.ToLower(s), "-"), "-")
}

func GetDomainFromSubdomain(subdom string) string {
	domstart := strings.Index(subdom, ".") + 1
	return subdom[domstart:len(subdom)]
}

func FilterPhoneNo(phone string) string {
	phone = strings.Replace(phone, "(", "", 10)
	phone = strings.Replace(phone, ")", "", 10)
	phone = strings.Replace(phone, "-", "", 10)
	phone = strings.Replace(phone, " ", "", 10)
	phone = strings.Replace(phone, "#", "", 10)

	return phone
}

func FormatPhoneNo(phone string) string {
	phn, err := strconv.ParseUint(phone, 10, 64)
	CheckErr(err)

	no := phn % 1e4
	xc := phn / 1e4 % 1e3
	phoneNo := fmt.Sprintf("%03d-%04d", xc, no)

	pfx := phn / 1e7
	if pfx != 0 {
		phoneNo = fmt.Sprintf("%d-%s", pfx, phoneNo)
	}

	return phoneNo
}

func GetUniqueStrings(string_slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range string_slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func EncodeStr(str string) string {
	return string(base64.StdEncoding.EncodeToString([]byte(str)))
}

func DecodeStr(str string) string {
	retstr, err := base64.StdEncoding.DecodeString(str)
	CheckErr(err)
	return string(retstr)
}

func GetFileAsString(asbFilePath string) string {

	file, err := os.Open(asbFilePath)
	CheckErr(err)

	filebytem, err := ioutil.ReadAll(file)
	CheckErr(err)
	file.Close()
	return string(filebytem)
}

func TrimString(s string) string {
	return strings.TrimLeft(strings.TrimRight(s, " "), " ")
}

func GetCharVal(i int) string {
	var foo = "abcdefghijklmnopqrstuvwxyz"
	return string(foo[i-1])
}
