//go:binary-only-package

package cached_stub

import (
	"common/bchcls/internal/common/global"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	"github.com/pkg/errors"
)
