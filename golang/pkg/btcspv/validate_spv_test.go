package btcspv

import (
	// "bytes"
	// "encoding/hex"
	// "encoding/json"
	// "io/ioutil"
	// "log"
	// "os"
	// "testing"

	// "github.com/stretchr/testify/suite"

	// sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *UtilsSuite) TestProve() {
	fixture := suite.Fixtures["prove"]

	for i := range fixture {
		testCase := fixture[i]
		expected := testCase.Output.([]byte)
		inputs := testCase.Input.(map[string]interface{})
		txIdLE := inputs["txIdLE"].([]byte)
		merkleRootLE := inputs["merkleRootLE"].([]byte)
		proof := inputs["proof"].([]byte)
		index := uint(inputs["index"].(int))
		actual := Prove(txIdLE, merkleRootLE, proof, index)
		suite.Equal(expected, actual)
	}
}

func (suite *UtilsSuite) TestCalculateTxId() {
	fixture := suite.Fixtures["calculateTxId"]

	for i := range fixture {
		testCase := fixture[i]
		expected := testCase.Output.([]byte)
		inputs := testCase.Input.(map[string]interface{})
		version := inputs["version"].([]byte)
		vin := inputs["vin"].([]byte)
		vout := inputs["vout"].([]byte)
		locktime := inputs["locktime"].([]byte)
		actual := CalculateTxId(version, vin, vout, locktime)
		suite.Equal(expected, actual)
	}
}

// func (suite *UtilsSuite)