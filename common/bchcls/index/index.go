//go:binary-only-package

package index

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/index/table_interface"
	"github.com/smannathan/common-sdk/common/bchcls/internal/common/global"
	"github.com/smannathan/common-sdk/common/bchcls/internal/index_i"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	"strings"
)
