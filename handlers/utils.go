package handlers

import (
	"fmt"
	"os"
	"regexp"
)

type Setting struct {
	Label    string `json:"label"`
	Type     string `json:"type"`
	Default  string `json:"default"`
	Required bool   `json:"required"`
}

type IntegrationStruct struct {
	Date struct {
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	} `json:"date"`
	Descriptions struct {
		AppName         string `json:"app_name"`
		AppDescription  string `json:"app_description"`
		AppLogo         string `json:"app_logo"`
		AppURL          string `json:"app_url"`
		BackgroundColor string `json:"background_color"`
	} `json:"descriptions"`
	IntegrationCategory string    `json:"integration_category"`
	IsActive            bool      `json:"is_active"`
	IntegrationType     string    `json:"integration_type"`
	KeyFeatures         []string  `json:"key_features"`
	Author              string    `json:"author"`
	Settings            []Setting `json:"settings"`
	TargetURL           string    `json:"target_url"`
}

type MsgRequest struct {
	ChannelID string    `json:"channel_id"`
	Settings  []Setting `json:"settings"`
	Message   string    `json:"message"`
}

func GenerateUniqueURL(settings []Setting) string {
	var (
		// form_name  string
		website    string
		channel_id string
	)
	var BASE_URL = os.Getenv("BASE_URL")
	for _, setting := range settings {
		switch setting.Label {
		// case "Form Name":
		// 	form_name = setting.Default

		case "Website":
			website = setting.Default
		case "ChannelID":
			channel_id = setting.Default
		}
	}
	return fmt.Sprintf("%s/formify/%s/%s", BASE_URL, website, channel_id)
}

// extractText extracts text from the first <p> tag
func ExtractText(input string) string {
	re := regexp.MustCompile(`<p>(.*?)</p>`) // Regex to match text inside <p> tags
	matches := re.FindStringSubmatch(input)

	if len(matches) > 1 {
		return matches[1] // Return only the extracted text
	}
	return ""
}
