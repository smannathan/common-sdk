//go:binary-only-package

package datastore_c

import (
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/custom_errors"
	"common/bchcls/datastore"
	"common/bchcls/internal/common/global"
	"common/bchcls/internal/datastore_i/datastore_c/cloudant"
	"common/bchcls/internal/datastore_i/datastore_c/ledger"
	"common/bchcls/utils"

	"encoding/json"
	"net/url"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/pkg/errors"
)