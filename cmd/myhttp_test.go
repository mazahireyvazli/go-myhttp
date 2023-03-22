package cmd

import (
	"testing"
)

func TestMyHTTPClient_SendHTTPRequest(t *testing.T) {
	type args struct {
		requestURL string
	}
	tests := []struct {
		name    string
		args    args
		want    *MyHTTPResponse
		wantErr bool
	}{
		{
			name: "successful request",
			args: args{requestURL: "https://httpbin.org/user-agent"},
			want: &MyHTTPResponse{
				URL:              "https://httpbin.org/user-agent",
				ResponseBodyHash: "6357927102a7c1025dc4e8bac96c00d2",
			},
			wantErr: false,
		},
		{
			name:    "invalid URL",
			args:    args{requestURL: "this is not a valid url"},
			want:    nil,
			wantErr: true,
		},
	}

	client := NewMyHTTPClient()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.SendHTTPRequest(tt.args.requestURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("MyHTTPClient.SendHTTPRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				if got.URL != tt.want.URL {
					t.Errorf("MyHTTPClient.SendHTTPRequest() URL = %v, want %v", got.URL, tt.want.URL)
				}
				if got.ResponseBodyHash != tt.want.ResponseBodyHash {
					t.Errorf("MyHTTPClient.SendHTTPRequest() ResponseBodyHash = %v, want %v", got.ResponseBodyHash, tt.want.ResponseBodyHash)
				}
			}
		})
	}
}

func TestMyHTTPClient_parseRequestURL(t *testing.T) {
	c := NewMyHTTPClient()

	tests := []struct {
		name    string
		url     string
		want    string
		wantErr bool
	}{
		{
			name: "valid http url",
			url:  "http://example.com",
			want: "http://example.com",
		},
		{
			name: "url wihtout scheme",
			url:  "example.com",
			want: "http://example.com",
		},
		{
			name: "valid https url",
			url:  "https://example.com",
			want: "https://example.com",
		},
		{
			name:    "invalid url",
			url:     "://example.com",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parsedURL, err := c.parseRequestURL(tt.url)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if parsedURL.String() != tt.want {
					t.Errorf("MyHTTPClient.SendHTTPRequest() url = %v, want %v", parsedURL.String(), tt.want)
				}
			}
		})
	}
}
