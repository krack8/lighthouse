package utils

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "Valid password",
			password: "password123",
			wantErr:  false,
		},
		{
			name:     "Empty password",
			password: "",
			wantErr:  true,
		},
		{
			name:     "Long password",
			password: strings.Repeat("a", 72), // bcrypt's maximum length
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword := HashPassword(tt.password)
			if (hashedPassword == "") != tt.wantErr {
				t.Errorf("HashPassword() error = %v, wantErr %v", hashedPassword == "", tt.wantErr)
				return
			}

			if !tt.wantErr && !strings.HasPrefix(hashedPassword, "$2a$") {
				t.Error("HashPassword() did not generate a valid bcrypt hash")
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		wrongPassword string
		wantMatch     bool
	}{
		{
			name:          "Matching passwords",
			password:      "password123",
			wrongPassword: "password123",
			wantMatch:     true,
		},
		{
			name:          "Non-matching passwords",
			password:      "password123",
			wrongPassword: "password124",
			wantMatch:     false,
		},
		{
			name:          "Case sensitivity",
			password:      "Password123",
			wrongPassword: "password123",
			wantMatch:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword := HashPassword(tt.password)
			if hashedPassword == "" {
				t.Fatal("Failed to hash password")
			}

			// Test matching password
			if got := CheckPassword(tt.wrongPassword, hashedPassword); got != tt.wantMatch {
				t.Errorf("CheckPassword() = %v, want %v", got, tt.wantMatch)
			}
		})
	}
}

func TestGenerateSecureToken(t *testing.T) {
	tests := []struct {
		name   string
		length int
		want   int // Expected length of hex string (twice the input length)
	}{
		{
			name:   "16 bytes token",
			length: 16,
			want:   32,
		},
		{
			name:   "32 bytes token",
			length: 32,
			want:   64,
		},
		{
			name:   "Zero length",
			length: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateSecureToken(tt.length)
			if len(got) != tt.want {
				t.Errorf("GenerateSecureToken() length = %v, want %v", len(got), tt.want)
			}

			// Verify it's a valid hex string
			if tt.length > 0 {
				if _, err := hex.DecodeString(got); err != nil {
					t.Errorf("GenerateSecureToken() generated invalid hex string: %v", err)
				}
			}
		})
	}
}

func TestGenerateResetToken(t *testing.T) {
	// Test multiple token generations
	tokens := make(map[string]bool)
	for i := 0; i < 100; i++ {
		token := GenerateResetToken()

		// Check token is not empty
		if token == "" {
			t.Error("GenerateResetToken() returned empty string")
		}

		// Check token length (32 bytes = 64 hex chars)
		if len(token) != 64 {
			t.Errorf("GenerateResetToken() returned token of length %d, want 64", len(token))
		}

		// Check uniqueness
		if tokens[token] {
			t.Error("GenerateResetToken() generated duplicate token")
		}
		tokens[token] = true

		// Verify it's a valid hex string
		if _, err := hex.DecodeString(token); err != nil {
			t.Errorf("GenerateResetToken() generated invalid hex string: %v", err)
		}
	}
}

func TestPasswordComplexity(t *testing.T) {
	tests := []struct {
		name              string
		password          string
		wantHash          bool
		testWrongPassword bool // Flag to control wrong password test
	}{
		{
			name:              "Complex password",
			password:          "Password123!@#",
			wantHash:          true,
			testWrongPassword: true,
		},
		{
			name:              "Simple password",
			password:          "password",
			wantHash:          true,
			testWrongPassword: true,
		},
		{
			name:              "Maximum length password",
			password:          strings.Repeat("a", 72),
			wantHash:          true,
			testWrongPassword: false, // Skip wrong password test for max length
		},
		{
			name:              "Too long password",
			password:          strings.Repeat("a", 73),
			wantHash:          false,
			testWrongPassword: false,
		},
		{
			name:              "Empty password",
			password:          "",
			wantHash:          false,
			testWrongPassword: false,
		},
		{
			name:              "Unicode password",
			password:          "パスワード123",
			wantHash:          true,
			testWrongPassword: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword := HashPassword(tt.password)
			hashExists := hashedPassword != ""

			if hashExists != tt.wantHash {
				t.Errorf("HashPassword() with %v returned hash = %v, want %v",
					tt.password, hashExists, tt.wantHash)
				return
			}

			// Only verify CheckPassword if we expect a valid hash
			if tt.wantHash {
				// Verify correct password works
				if !CheckPassword(tt.password, hashedPassword) {
					t.Error("CheckPassword() failed to verify valid password")
				}

				// Only test wrong password for non-maximum length passwords
				if tt.testWrongPassword {
					wrongPassword := tt.password + "wrong"
					if CheckPassword(wrongPassword, hashedPassword) {
						t.Error("CheckPassword() incorrectly verified wrong password")
					}
				}
			}
		})
	}
}

// Additional test specifically for maximum length password
func TestMaxLengthPassword(t *testing.T) {
	maxLengthPwd := strings.Repeat("a", 72)
	hashedPassword := HashPassword(maxLengthPwd)

	if hashedPassword == "" {
		t.Fatal("Failed to hash maximum length password")
	}

	// Test correct password
	if !CheckPassword(maxLengthPwd, hashedPassword) {
		t.Error("Failed to verify correct maximum length password")
	}

	// Test slightly different password
	slightlyDifferent := strings.Repeat("a", 71) + "b"
	if CheckPassword(slightlyDifferent, hashedPassword) {
		t.Error("Incorrectly verified different maximum length password")
	}
}
