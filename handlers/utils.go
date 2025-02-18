package handlers

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

func ParseSettings(settings []Setting) (form_name, website string) {

	for _, setting := range settings {
		switch setting.Label {
		case "Form Name":
			form_name = setting.Default

		case "Website":
			website = setting.Default
		}
	}
	return
}
