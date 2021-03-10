//go:binary-only-package

package cloudant_index

import (
	"github.com/smannathan/common-sdk/common/bchcls/cached_stub"
	"github.com/smannathan/common-sdk/common/bchcls/crypto"
	"github.com/smannathan/common-sdk/common/bchcls/datastore"
	"github.com/smannathan/common-sdk/common/bchcls/internal/datastore_i/datastore_c"
	"github.com/smannathan/common-sdk/common/bchcls/internal/datastore_i/datastore_c/cloudant"
	"github.com/smannathan/common-sdk/common/bchcls/utils"

	"encoding/json"
	"net/url"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/ledger/queryresult"
	"github.com/pkg/errors"
)