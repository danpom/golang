package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

//slice to track the names issued
var names []string

// Define the Robot type here.
type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = UniqueName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

//contains searches an slice of string for a specified string and returns true if it's found
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

//Anyname generates a random robot name
func AnyName() string {
	rand.Seed(time.Now().UnixNano())
	//Generate 2 random character between uppercase A and Z.
	rc1 := 'A' + rune(rand.Intn(26))
	rc2 := 'A' + rune(rand.Intn(26))

	//Generate random 3 digit number - this may result in sub 100 numbers so need to pad with leading 0's
	rn := rand.Intn(999)
	rns := strconv.Itoa(rn)
	//padding with leading 0's
	for i := len(rns); i < 3; i++ {
		rns = "0" + rns
	}

	//joining our letters numbers to create the robot name
	tn := string(rc1) + string(rc2) + rns
	return tn
}

func UniqueName() string {
	tempName := AnyName()

	//repeatedly generate a new name until we find one that doesn't exist
	for contains(names, tempName) {
		tempName = AnyName()
	}
	names = append(names, tempName)
	return tempName
}
