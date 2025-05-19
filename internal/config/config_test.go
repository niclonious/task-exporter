package config

import "testing"

func Test_validateConfig(t *testing.T) {
	tests := []struct {
		name    string
		conf    Config
		wantErr bool
	}{
		{
			name: "Should_Succeed",
			conf: Config{
				Port: 8080,
				Env:  "dev",
			},
			wantErr: false,
		},
		{
			name: "Should_Fail_Invalid_Port",
			conf: Config{
				Port: 10,
				Env:  "dev",
			},
			wantErr: true,
		},
		{
			name: "Should_Fail_Invalid_Env",
			conf: Config{
				Port: 8080,
				Env:  "ci",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateConfig(tt.conf); (err != nil) != tt.wantErr {
				t.Errorf("validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
