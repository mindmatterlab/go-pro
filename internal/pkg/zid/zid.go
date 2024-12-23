package zid

import (
	"github.com/mindmatterlab/acex/pkg/id"
)

const defaultABC = "abcdefghijklmnopqrstuvwxyz1234567890"

type ZID string

const (
	// ID for the user resource in acex-usercenter.
	User ZID = "user"
	// ID for the order resource in acex-fakeserver.
	Order ZID = "order"
)

func (zid ZID) String() string {
	return string(zid)
}

func (zid ZID) New(i uint64) string {
	// use custom option
	str := id.NewCode(
		i,
		id.WithCodeChars([]rune(defaultABC)),
		id.WithCodeL(6),
		id.WithCodeSalt(Salt()),
	)
	return zid.String() + "-" + str
}
