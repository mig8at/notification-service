package model

type Limit struct {
	Type     NotificationType `mapstructure:"type"`
	Duration string           `mapstructure:"duration"`
	Count    int              `mapstructure:"count"`
}
