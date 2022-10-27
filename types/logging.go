/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */

package types

import (
	"github.com/aergoio/aergo/internal/enc"
	"github.com/rs/zerolog"
)

type LogTxHash struct {
	*Tx
}

func (t LogTxHash) MarshalZerologObject(e *zerolog.Event) {
	e.Str("txID",enc.ToString(t.Hash))
}

type LogTx struct {
	*Tx
}

func (t LogTx) MarshalZerologObject(e *zerolog.Event) {
	e.Str("txID",enc.ToString(t.GetHash())).Str("account",enc.ToString(t.Body.Account)).Uint64("nonce",t.Body.Nonce)
}

type LogBase58 struct {
	Bytes *[]byte
}

func (t LogBase58) MarshalZerologObject(e *zerolog.Event) {
	e.Str("b58",enc.ToString(*t.Bytes))
}