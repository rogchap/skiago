package skia

import "testing"

func TestFontMgr(t *testing.T) {
	fm := NewSystemFontMgr()
	c := fm.CountFamiles()

	// TODO: This is not a test! make it one.
	for i := 0; i < c; i++ {
		println(fm.FamilyName(i))
	}
}
