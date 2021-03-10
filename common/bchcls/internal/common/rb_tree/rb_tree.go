//go:binary-only-package

package rb_tree

import (
	"bytes"
	"encoding/json"
	"os"

	"common/bchcls/cached_stub"
	"common/bchcls/custom_errors"
	"common/bchcls/internal/common/global"
	"common/bchcls/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	"github.com/pkg/errors"
)