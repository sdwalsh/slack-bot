package main

import "github.com/slack-go/slack"

// This is a static modal, so we'll define it up here
var viewRequest = slack.ModalViewRequest{
	Type: "modal",
	Title: &slack.TextBlockObject{
		Type:  slack.PlainTextType,
		Text:  "Rootly Interview Bot",
		Emoji: true,
	},
	Blocks: slack.Blocks{
		BlockSet: []slack.Block{
			slack.HeaderBlock{
				Type: slack.MBTHeader,
				Text: &slack.TextBlockObject{
					Type:  slack.PlainTextType,
					Text:  "Create an Incident",
					Emoji: true,
				},
				BlockID: "",
			},
			slack.DividerBlock{Type: "divider"},
			slack.InputBlock{
				Type: slack.MBTInput,
				BlockID: "title-block",
				Label: &slack.TextBlockObject{
					Type:  slack.PlainTextType,
					Text:  "Title",
					Emoji: true,
				},
				Element: slack.PlainTextInputBlockElement{
					Type:     slack.METPlainTextInput,
					ActionID: "title",
				},
				Optional: false,
			},
			slack.InputBlock{
				Type: slack.MBTInput,
				BlockID: "desc-block",
				Label: &slack.TextBlockObject{
					Type:  slack.PlainTextType,
					Text:  "Description",
					Emoji: true,
				},
				Element: slack.PlainTextInputBlockElement{
					Type:     slack.METPlainTextInput,
					ActionID: "desc",
				},
				Optional: true,
			},
			slack.InputBlock{
				Type: slack.MBTInput,
				BlockID: "sev-block",
				Label: &slack.TextBlockObject{
					Type:  slack.PlainTextType,
					Text:  "Severity",
					Emoji: true,
				},
				Element: slack.SelectBlockElement{
					Type: slack.OptTypeStatic,
					Placeholder: &slack.TextBlockObject{
						Type:  slack.PlainTextType,
						Text:  "Select an item",
						Emoji: true,
					},
					ActionID: "sev",
					Options: []*slack.OptionBlockObject{
						{
							Text: &slack.TextBlockObject{
								Type:  slack.PlainTextType,
								Text:  "sev0",
								Emoji: true,
							},
							Value: "sev0",
						},
						{
							Text: &slack.TextBlockObject{
								Type:  slack.PlainTextType,
								Text:  "sev1",
								Emoji: true,
							},
							Value: "sev1",
						},
						{
							Text: &slack.TextBlockObject{
								Type:  slack.PlainTextType,
								Text:  "sev2",
								Emoji: true,
							},
							Value: "sev2",
						},
					},
				},
				Hint:           nil,
				Optional:       false,
				DispatchAction: false,
			},
		},
	},
	Close: &slack.TextBlockObject{
		Type:     slack.PlainTextType,
		Text:     "Close",
		Emoji:    true,
		Verbatim: false,
	},
	Submit: &slack.TextBlockObject{
		Type:     slack.PlainTextType,
		Text:     "Submit",
		Emoji:    true,
		Verbatim: false,
	},
	PrivateMetadata: "",
	CallbackID:      "",
	ClearOnClose:    true,
	NotifyOnClose:   false,
	ExternalID:      "",
}
