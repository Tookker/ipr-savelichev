package jwt

import (
	"fmt"
	"log"
	"testing"

	"go.uber.org/zap"

	"ipr-savelichev/internal/config"
	"ipr-savelichev/internal/store/tables/task/mocks"
)

// MiddlewareJWT(next http.Handler) http.Handler

func getZapLogger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("Err create logger.")
		return nil, err
	}

	return logger, nil
}

func loadConfig() *config.Config {
	config, err := config.LoadConfig("/home/adminlocal/go/src/ipr-savelichev/config/config.yaml")
	if err != nil {
		log.Fatalln(err.Error())
		return nil
	}

	return config
}

// GenerateJWT(login string) (string, error)
func TestGenerateJWT(t *testing.T) {
	type args struct {
		login string
	}

	type want struct {
		token string
		err   error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "get jwt token",
			args: args{
				login: "user",
			},
			want: want{
				token: "",
				err:   nil,
			},
			wantErr: false,
		},
	}

	config := loadConfig()
	if config == nil {
		t.Errorf("Err config")
		return
	}

	logger, err := getZapLogger()
	if err != nil {
		t.Errorf("Err logger")
		return
	}

	jwtRepo, err := NewJWTController(config, logger)
	if err != nil {
		t.Errorf("Err repo jwt")
		return
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("GenerateJWT", tt.args.login).Return(tt.want)

			_, err := jwtRepo.GenerateJWT(tt.args.login)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateJWT error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

// CheckJWT(token string) error
func TestCheckJWT(t *testing.T) {
	type args struct {
		token string
	}

	type want struct {
		err error
	}

	config := loadConfig()
	if config == nil {
		t.Errorf("Err config")
		return
	}

	logger, err := getZapLogger()
	if err != nil {
		t.Errorf("Err logger")
		return
	}

	jwtRepo, err := NewJWTController(config, logger)
	if err != nil {
		t.Errorf("Err repo jwt")
		return
	}

	token, err := jwtRepo.GenerateJWT("user")
	if err != nil {
		t.Errorf("Gen jwt token")
		return
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "check correct jwt token",
			args: args{
				token: token,
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "check wrong jwt token",
			args: args{
				token: token,
			},
			want: want{
				err: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &mocks.Task{}
			repo.On("CheckJWT", tt.args.token).Return(tt.want)

			_, err := jwtRepo.GenerateJWT(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckJWT error = %v, whantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
