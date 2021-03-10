//go:binary-only-package

package test_utils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)