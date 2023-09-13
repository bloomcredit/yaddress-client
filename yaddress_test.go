package yaddress_test

import (
	"testing"
	"yaddress"
)

func TestYaddressWithEmptyUserKey(t *testing.T) {
	_, err := yaddress.NewClient("")
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}
}

func TestYaddressWithEmptyFields(t *testing.T) {
	yd, err := yaddress.NewClient("")
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = yd.ProcessAddress("", "")

	if err == nil {
		t.Error("Should give an error if both address lines are empty")
	}
}

func TestYaddressWithEmptyAddress1ButPresentAddress2(t *testing.T) {
	yd, err := yaddress.NewClient("")
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = yd.ProcessAddress("", "Chicago, IL")

	if err != nil {
		t.Error("Should not give an error, it's possible to provide only addressLine2")
	}
}

func TestYaddressWithPresentAddress1ButAbsentAddress2(t *testing.T) {
	yd, err := yaddress.NewClient("")
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = yd.ProcessAddress("1009 S Oakley", "")

	if err == nil {
		t.Error("Should give an error, address2 is required")
	}
}

func TestYaddressWithWrongAddress(t *testing.T) {
	yd, err := yaddress.NewClient("")
	if err != nil {
		t.Error("Failed with API KEY set.", err)
	}

	_, err = yd.ProcessAddress("506 Fourth Avenue Unit 1", "Chicago, IL")

	if err == nil {
		t.Error("Should give an error, passed address is wrong")
	}
}
