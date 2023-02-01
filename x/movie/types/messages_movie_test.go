package types

import (
	"testing"

	"movie/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateMovie_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateMovie
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateMovie{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address and empty title",
			msg: MsgCreateMovie{
				Creator: sample.AccAddress(),
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid address and valid title",
			msg: MsgCreateMovie{
				Creator: sample.AccAddress(),
				Title: "title1",
			},
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

func TestMsgUpdateMovie_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateMovie
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateMovie{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateMovie{
				Creator: sample.AccAddress(),
			},
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

func TestMsgDeleteMovie_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteMovie
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteMovie{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteMovie{
				Creator: sample.AccAddress(),
			},
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
