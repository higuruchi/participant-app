package config

import (
	"testing"
)

func newConfig(
	serverPort int,
	dbUser string,
	dbPassword string,
	dbIP string,
	dbPort int,
	dbName string,
) *Config {
	return &Config {
		ServerConfig{
			Port: serverPort,
		},
		DBConfig{
			User: dbUser,
			Password: dbPassword,
			IP: dbIP,
			Port: dbPort,
			Name: dbName,
		},
	}
}

func TestConfigValidation(t *testing.T) {
	type args struct {
		serverPort int
		dbUser string
		dbPassword string
		dbIP string
		dbPort int
		dbName string
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "check minimu of port number",
			args: args{
				serverPort: 1,
				dbUser: "user",
				dbPassword: "password",
				dbIP: "127.0.0.1",
				dbPort: 1,
				dbName: "app-db",
			},
			want: nil,
		},
		{
			name: "check max of port number",
			args: args{
				serverPort: 65535,
				dbUser: "user",
				dbPassword: "password",
				dbIP: "127.0.0.1",
				dbPort: 65535,
				dbName: "app-db",
			},
			want: nil,
		},
		{
			name: "check out of range of server port number",
			args: args{
				serverPort: -1,
				dbUser: "user",
				dbPassword: "password",
				dbIP: "127.0.0.1",
				dbPort: 65535,
				dbName: "app-db",
			},
			want: ErrInvalidServerPort,
		},
		{
			name: "check out of range of db port number",
			args: args{
				serverPort: 65535,
				dbUser: "user",
				dbPassword: "password",
				dbIP: "127.0.0.1",
				dbPort: -1,
				dbName: "app-db",
			},
			want: ErrInvalidDBPort,
		},
		{
			name: "check database user name error",
			args: args{
				serverPort: 1323,
				dbUser: "",
				dbPassword: "password",
				dbIP: "127.0.0.1",
				dbPort: 3306,
				dbName: "app-db",
			},
			want: ErrInvalidDBUser,
		},
		{
			name: "check database password error",
			args: args{
				serverPort: 1323,
				dbUser: "user",
				dbPassword: "",
				dbIP: "127.0.0.1",
				dbPort: 3306,
				dbName: "app-db",
			},
			want: ErrInvalidDBPassword,
		},
		{
			name: "check database ip error",
			args: args{
				serverPort: 1323,
				dbUser: "user",
				dbPassword: "password",
				dbIP: "127.0.1",
				dbPort: 3306,
				dbName: "app-db",
			},
			want: ErrInvalidDBIP,
		},
		{
			name: "check database name error",
			args: args{
				serverPort: 1323,
				dbUser: "user",
				dbPassword: "password",
				dbIP: "127.0.0.1",
				dbPort: 3306,
				dbName: "",
			},
			want: ErrInvalidDBName,
		},
	}

	for _, tt := range tests {
		config := newConfig(
			tt.args.serverPort,
			tt.args.dbUser,
			tt.args.dbPassword,
			tt.args.dbIP,
			tt.args.dbPort,
			tt.args.dbName,
		)

		t.Run(tt.name, func(t *testing.T) {
			if got := config.configValidation(); got != tt.want {
				t.Errorf("config.configValidation() = %v, want=%v", got, tt.want)
			}
		})
	}
}