package types_test

import (
	"testing"

	"github.com/Sifchain/sifnode/x/dispensation/types"
	"github.com/stretchr/testify/assert"
)


func TestQueryRecordsByDistributionName(t *testing.T) {
	distributionName := types.QueryRecordsByDistrName
	status := types.DistributionStatus_DISTRIBUTION_STATUS_PENDING
	result := types.NewQueryRecordsByDistributionName(distributionName, status)
	assert.Equal(t, distributionName, result.DistributionName)
    assert.Equal(t, status, result.Status)
}

func TestQueryRecordsByRecipientAddr_ValidBasic(t *testing.T) {
	address := sdk.AccAddress("address___")
	t.Log(address.String())

	result := types.NewQueryRecordsByRecipientAddr(address.String())
	t.Log(result)
	assert.Equal(t, address.String(), result.Address)
}

func TestQueryUserClaims_ValidBasic(t *testing.T) {
	distributiontype := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	t.Log(distributiontype)
	result := types.NewQueryUserClaims(distributiontype)
	t.Log(result)
	assert.Equal(t, distributiontype, result.UserClaimType)
}
