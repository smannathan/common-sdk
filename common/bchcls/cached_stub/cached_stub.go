//go:binary-only-package

package cached_stub

import (
	"common/bchcls/crypto"
	"common/bchcls/custom_errors"
	"common/bchcls/internal/common/global"
	"common/bchcls/utils"

	"sort"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)