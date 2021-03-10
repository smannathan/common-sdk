//go:binary-only-package

package cloudant_index

import (
	"common/bchcls/cached_stub"
	"common/bchcls/crypto"
	"common/bchcls/datastore"
	"common/bchcls/internal/datastore_i/datastore_c"
	"common/bchcls/internal/datastore_i/datastore_c/cloudant"
	"common/bchcls/utils"

	"encoding/json"
	"net/url"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	"github.com/pkg/errors"
)