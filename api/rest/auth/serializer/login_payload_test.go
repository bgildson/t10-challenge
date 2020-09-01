package serializer

import "testing"

func TestLoginPayload(t *testing.T) {
	tt := []struct {
		description string
		in          LoginPayload
		out         error
	}{
		{
			description: "an empty input",
			in:          LoginPayload{},
			out:         ErrRequiredEmail,
		},
		{
			description: "only email empty",
			in:          LoginPayload{Password: "123456"},
			out:         ErrRequiredEmail,
		},
		{
			description: "only password empty",
			in:          LoginPayload{Email: "user@email.com"},
			out:         ErrRequiredPassword,
		},
		{
			description: "everything ok",
			in:          LoginPayload{Email: "user@email.com", Password: "123456"},
			out:         nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			err := tc.in.Validate()
			if err != tc.out {
				t.Errorf("was expecting\n%v\nbut returns\n%v", tc.out, err)
			}
		})
	}
}
