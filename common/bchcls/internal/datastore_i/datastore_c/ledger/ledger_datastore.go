//go:binary-only-package

package ledger

import (
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/datastore"
	"common/bchcls/internal/common/global"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)