//go:binary-only-package

package index

import (
	"common/bchcls/cached_stub"
	"common/bchcls/index/table_interface"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/index_i"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	"strings"
)
