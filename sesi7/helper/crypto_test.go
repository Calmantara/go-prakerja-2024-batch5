package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// style yang digunakan di UBER
func TestHashPassword(t *testing.T) {
	type want struct {
		res string
		err error
	}

	testCases := []struct {
		desc  string
		input string
		want  want
	}{
		{
			desc:  "error has password too long",
			input: string(make([]byte, 100)),
			want: want{
				res: "",
				err: bcrypt.ErrPasswordTooLong,
			},
		},
		{
			desc:  "ok",
			input: "mysecretpassword",
			want: want{
				res: "",
				err: nil,
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			res, err := HashPassword(tC.input)
			if tC.want.err != nil {
				assert.Equal(t, tC.want.err, err)
				assert.Equal(t, tC.want.res, res)
				return
			}

			// kita inginkan response actual tidak sama dengan
			// string kosong
			assert.NotEqual(t, tC.want.res, res)
			// kita menginginkan err nil
			assert.Nil(t, err)
		})
	}
}

func TestCheckPasswordTest(t *testing.T) {
	type want struct {
		same bool
	}

	type input struct {
		password string
		hash     string
	}

	testCases := []struct {
		desc  string
		input input
		want  want
	}{
		{
			desc: "error not same",
			input: input{
				password: "mysecretpassword",
				hash:     "somerandomhash",
			},
			want: want{
				same: false,
			},
		},
		{
			desc: "ok",
			input: input{
				password: "mysecretpassword",
			},
			want: want{
				same: true,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if tC.input.hash == "" {
				tC.input.hash, _ = HashPassword(tC.input.password)
			}
			res := CheckPasswordHash(tC.input.password, tC.input.hash)
			assert.Equal(t, tC.want.same, res)
		})
	}
}
