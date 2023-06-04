package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMsgCreateUser_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateUser
		err  error
	}{
		{
			name: "no validation required",
			msg: MsgCreateUser{
				Name:     "test",
				Password: "password",
			},
			err: nil, // No validation error expected
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
