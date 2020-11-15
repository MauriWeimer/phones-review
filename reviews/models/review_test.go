package models

import "testing"

// Convenciones: todo lo que arranque con New devuelve un puntero, y lo que
// arranque con Make devuelve una estructura

func NewReview(stars int, comment string) *CreateReviewCMD {
	return &CreateReviewCMD{
		Stars:   stars,
		Comment: comment,
	}
}

func Test_withCorrectParams(t *testing.T) {
	r := NewReview(4, "The Iphone X looks good")

	err := r.validate()

	if err != nil {
		t.Error("the validation not pass")
		t.Fail()
	}
}

func Test_shouldFailWithWrongNumbersOfStars(t *testing.T) {
	r := NewReview(8, "The Iphone X looks good")

	err := r.validate()

	if err == nil {
		t.Error("should fail with 5 stars")
		t.Fail()
	}
}
