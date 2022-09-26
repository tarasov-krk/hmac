package hmac

import "testing"

func TestSha256(t *testing.T) {
	type args struct {
		data string
		key  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				data: "message",
				key:  "private_key",
			},
			want: "22253dc44300b5e70067b77cd533d333b38871a0d015a19ee92eddf5641e46ec",
		},
		{
			name: "2",
			args: args{
				data: "",
				key:  "private_key",
			},
			want: "82d5208ddf2b2096c17f879c739a5f440e87155f3308b9cf44a91d9104b10b8c",
		},
		{
			name: "3",
			args: args{
				data: "message",
				key:  "",
			},
			want: "eb08c1f56d5ddee07f7bdf80468083da06b64cf4fac64fe3a90883df5feacae4",
		},
		{
			name: "4",
			args: args{
				data: "",
				key:  "",
			},
			want: "b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha256(tt.args.data, tt.args.key); got != tt.want {
				t.Errorf("Sha256() = %v, want %v", got, tt.want)
			}
		})
	}
}
