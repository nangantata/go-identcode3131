package identcode3131_test

import (
	"testing"

	identcode3131 "github.com/nangantata/go-identcode3131"
)

func runPackUnpackIdentCodeTest(t *testing.T, virSn int32) {
	randKey, _ := identcode3131.GenerateNonZeroRandomKey(8)
	identCode := identcode3131.PackIdentCode(virSn, randKey)
	if identCode == 0 {
		t.Errorf("identCode should not be 0: virSn=%d, randKey=%d", virSn, randKey)
	}
	restVirSn, restRandKey := identcode3131.UnpackIdentCode(identCode)
	if restVirSn != virSn {
		t.Errorf("failed to restore serialValue from identCode: virSn=%d, randKey=%d, identCode=0x%016X, restVirSn=%d, restRandKey=%d",
			virSn, randKey, identCode, restVirSn, restRandKey)
	}
	if restRandKey != randKey {
		t.Errorf("failed to restore randomKey from identCode: virSn=%d, randKey=%d, identCode=0x%016X, restVirSn=%d, restRandKey=%d",
			virSn, randKey, identCode, restVirSn, restRandKey)
	}

}

func TestPackUnpackIdentCode01(t *testing.T) {
	for virSn := int32(1); virSn < 0x0000FFFF; virSn++ {
		runPackUnpackIdentCodeTest(t, virSn)
	}
	for virSn := int32(0x7FFF0000); virSn < 0x7FFFFFFF; virSn++ {
		runPackUnpackIdentCodeTest(t, virSn)
	}
	runPackUnpackIdentCodeTest(t, 0x7FFFFFFF)
}

func runPackUnpackIdentTokenTest(t *testing.T, virSn int32) {
	randKey, _ := identcode3131.GenerateNonZeroRandomKey(8)
	identToken := identcode3131.PackIdentToken(virSn, randKey)
	if identToken == "" {
		t.Errorf("identToken should not be empty: virSn=%d, randKey=%d", virSn, randKey)
	}
	restVirSn, restRandKey, err := identcode3131.UnpackIdentToken(identToken)
	if nil != err {
		t.Errorf("failed to unpack identToken: virSn=%d, randKey=%d, identToken=%s, err=%v",
			virSn, randKey, identToken, err)
	}
	if restVirSn != virSn {
		t.Errorf("failed to restore serialValue from identToken: virSn=%d, randKey=%d, identToken=%s, restVirSn=%d, restRandKey=%d",
			virSn, randKey, identToken, restVirSn, restRandKey)
	}
	if restRandKey != randKey {
		t.Errorf("failed to restore randomKey from identToken: virSn=%d, randKey=%d, identToken=%s, restVirSn=%d, restRandKey=%d",
			virSn, randKey, identToken, restVirSn, restRandKey)
	}
}

func TestPackUnpackIdentToken01(t *testing.T) {
	for virSn := int32(1); virSn < 0x0000FFFF; virSn++ {
		runPackUnpackIdentTokenTest(t, virSn)
	}
	for virSn := int32(0x7FFF0000); virSn < 0x7FFFFFFF; virSn++ {
		runPackUnpackIdentTokenTest(t, virSn)
	}
	runPackUnpackIdentTokenTest(t, 0x7FFFFFFF)
}

func TestUnpackIdentToken02(t *testing.T) {
	restVirSn, restRandKey, err := identcode3131.UnpackIdentToken("")
	if nil == err {
		t.Errorf("unexpect unpack result: restVirSn=%d, restRandKey=%d", restVirSn, restRandKey)
	} else {
		t.Logf("expected unpack error: %v", err)
	}
	restVirSn, restRandKey, err = identcode3131.UnpackIdentToken("123456")
	if nil == err {
		t.Errorf("unexpect unpack result: restVirSn=%d, restRandKey=%d", restVirSn, restRandKey)
	} else {
		t.Logf("expected unpack error: %v", err)
	}
	restVirSn, restRandKey, err = identcode3131.UnpackIdentToken("123456abcd+")
	if nil == err {
		t.Errorf("unexpect unpack result: restVirSn=%d, restRandKey=%d", restVirSn, restRandKey)
	} else {
		t.Logf("expected unpack error: %v", err)
	}
}
