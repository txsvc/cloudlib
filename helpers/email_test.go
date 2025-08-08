package helpers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/txsvc/stdlib/v2"
)

func TestMailgunSimpleEmail(t *testing.T) {
	domain := stdlib.GetString(MailgunEmailDomainENV, "")
	apiKey := stdlib.GetString(MailgunApiKeyENV, "")
	if domain == "" || apiKey == "" {
		return // skip the test
	}

	err := MailgunSimpleEmail("test@podops.dev", "ops@txs.vc", "TestMailgunSimpleEmail", "Testing TestMailgunSimpleEmail")
	assert.NoError(t, err)
}

func TestMailgunSimpleEmailMissingDomain(t *testing.T) {
	// Store original env values to restore later
	originalDomain := stdlib.GetString(MailgunEmailDomainENV, "")

	// Temporarily unset domain env var
	_ = os.Unsetenv(MailgunEmailDomainENV)
	defer func() {
		if originalDomain != "" {
			_ = os.Setenv(MailgunEmailDomainENV, originalDomain)
		}
	}()

	err := MailgunSimpleEmail("test@example.com", "recipient@example.com", "Test", "Body")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid email configuration")
}

func TestMailgunSimpleEmailMissingApiKey(t *testing.T) {
	// Store original env values to restore later
	originalApiKey := stdlib.GetString(MailgunApiKeyENV, "")

	// Set domain but unset API key
	_ = os.Setenv(MailgunEmailDomainENV, "test.example.com")
	_ = os.Unsetenv(MailgunApiKeyENV)
	defer func() {
		_ = os.Unsetenv(MailgunEmailDomainENV)
		if originalApiKey != "" {
			_ = os.Setenv(MailgunApiKeyENV, originalApiKey)
		}
	}()

	err := MailgunSimpleEmail("test@example.com", "recipient@example.com", "Test", "Body")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid email configuration")
}

func TestMailgunSimpleEmailEmptyDomain(t *testing.T) {
	// Set empty domain
	_ = os.Setenv(MailgunEmailDomainENV, "")
	_ = os.Setenv(MailgunApiKeyENV, "test-key")
	defer func() {
		_ = os.Unsetenv(MailgunEmailDomainENV)
		_ = os.Unsetenv(MailgunApiKeyENV)
	}()

	err := MailgunSimpleEmail("test@example.com", "recipient@example.com", "Test", "Body")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid email configuration")
}

func TestMailgunSimpleEmailEmptyApiKey(t *testing.T) {
	// Set empty API key
	_ = os.Setenv(MailgunEmailDomainENV, "test.example.com")
	_ = os.Setenv(MailgunApiKeyENV, "")
	defer func() {
		_ = os.Unsetenv(MailgunEmailDomainENV)
		_ = os.Unsetenv(MailgunApiKeyENV)
	}()

	err := MailgunSimpleEmail("test@example.com", "recipient@example.com", "Test", "Body")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid email configuration")
}

func TestMailgunConstants(t *testing.T) {
	// Test that constants are properly defined
	assert.Equal(t, "MAILGUN_EMAIL_DOMAIN", MailgunEmailDomainENV)
	assert.Equal(t, "MAILGUN_API_KEY", MailgunApiKeyENV)
}
