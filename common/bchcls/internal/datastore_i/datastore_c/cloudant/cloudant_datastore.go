//go:binary-only-package

package cloudant

import (
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/datastore"
	"common/bchcls/internal/common/global"
	"common/bchcls/utils"

	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)