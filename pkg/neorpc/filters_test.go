package neorpc

import (
	"testing"

	"github.com/nspcc-dev/neo-go/pkg/core/mempoolevent"
	"github.com/nspcc-dev/neo-go/pkg/smartcontract"
	"github.com/nspcc-dev/neo-go/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestBlockFilterCopy(t *testing.T) {
	var bf, tf *BlockFilter

	require.Nil(t, bf.Copy())

	bf = new(BlockFilter)
	tf = bf.Copy()
	require.Equal(t, bf, tf)

	bf.Primary = new(byte)
	*bf.Primary = 42

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Primary = 100
	require.NotEqual(t, bf, tf)

	bf.Since = new(uint32)
	*bf.Since = 42

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Since = 100500
	require.NotEqual(t, bf, tf)

	bf.Till = new(uint32)
	*bf.Till = 42

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Till = 100500
	require.NotEqual(t, bf, tf)
}

func TestTxFilterCopy(t *testing.T) {
	var bf, tf *TxFilter

	require.Nil(t, bf.Copy())

	bf = new(TxFilter)
	tf = bf.Copy()
	require.Equal(t, bf, tf)

	bf.Sender = &util.Uint160{1, 2, 3}

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Sender = util.Uint160{3, 2, 1}
	require.NotEqual(t, bf, tf)

	bf.Signer = &util.Uint160{1, 2, 3}

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Signer = util.Uint160{3, 2, 1}
	require.NotEqual(t, bf, tf)
}

func TestMempoolFilterCopy(t *testing.T) {
	var bf, tf *MempoolEventFilter

	require.Nil(t, bf.Copy())

	bf = new(MempoolEventFilter)
	tf = bf.Copy()
	require.Equal(t, bf, tf)

	bf.Sender = &util.Uint160{1, 2, 3}

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Sender = util.Uint160{3, 2, 1}
	require.NotEqual(t, bf, tf)

	bf.Signer = &util.Uint160{1, 2, 3}

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Signer = util.Uint160{3, 2, 1}
	require.NotEqual(t, bf, tf)

	added := mempoolevent.TransactionAdded
	bf.Type = &added

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Type = mempoolevent.TransactionRemoved
	require.NotEqual(t, bf, tf)
}

func TestNotificationFilterCopy(t *testing.T) {
	var bf, tf *NotificationFilter

	require.Nil(t, bf.Copy())

	bf = new(NotificationFilter)
	tf = bf.Copy()
	require.Equal(t, bf, tf)

	bf.Contract = &util.Uint160{1, 2, 3}

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Contract = util.Uint160{3, 2, 1}
	require.NotEqual(t, bf, tf)

	bf.Name = new(string)
	*bf.Name = "ololo"

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Name = "azaza"
	require.NotEqual(t, bf, tf)

	var err error
	bf.Parameters, err = smartcontract.NewParametersFromValues(1, "2", []byte{3})
	require.NoError(t, err)

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	bf.Parameters[0], bf.Parameters[1] = bf.Parameters[1], bf.Parameters[0]
	require.NotEqual(t, bf, tf)
}

func TestExecutionFilterCopy(t *testing.T) {
	var bf, tf *ExecutionFilter

	require.Nil(t, bf.Copy())

	bf = new(ExecutionFilter)
	tf = bf.Copy()
	require.Equal(t, bf, tf)

	bf.State = new(string)
	*bf.State = "ololo"

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.State = "azaza"
	require.NotEqual(t, bf, tf)

	bf.Container = &util.Uint256{1, 2, 3}

	tf = bf.Copy()
	require.Equal(t, bf, tf)
	*bf.Container = util.Uint256{3, 2, 1}
	require.NotEqual(t, bf, tf)
}
