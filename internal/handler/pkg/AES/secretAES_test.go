package AES

import "testing"

func TestEncryptAES(t *testing.T) {
	type args struct {
		plaintext string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "测试加密",
			args:    args{plaintext: "odin123456"},
			wantErr: false,
			want:    "dac740b6b799c288401d4932c1d25f27",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := EncryptAES(tt.args.plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncryptAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("EncryptAES() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecryptAES(t *testing.T) {
	type args struct {
		ct             string
		placeholderLen int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "测试解密",
			args:    args{ct: "dac740b6b799c288401d4932c1d25f27", placeholderLen: 6},
			wantErr: false,
			want:    "odin123456",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecryptAES(tt.args.ct, tt.args.placeholderLen)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecryptAES() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecryptAES() got = %v, want %v", got, tt.want)
			}
		})
	}
}
