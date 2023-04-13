package domain

// Notification contains necessary data to send notifications, for example, Summary notifications via mail
type Notification struct {
	Content    string
	Subject    string
	Recipients []string
	Branding   string
}
