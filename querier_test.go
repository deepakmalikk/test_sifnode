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

func TestQueryRecordsByRecipientAddr(t *testing.T) {
	address := types.QueryRecordsByRecipient
	result := types.NewQueryRecordsByRecipientAddr(address)
	assert.Equal(t, address, result.Address)
}

func TestQueryUserClaims_ValidBasic(t *testing.T) {
	distributiontype := types.DistributionType_DISTRIBUTION_TYPE_LIQUIDITY_MINING
	t.Log(distributiontype)
	result := types.NewQueryUserClaims(distributiontype)
	t.Log(result)
	assert.Equal(t, distributiontype, result.UserClaimType)
}
