package css

func DefaultCSS() string {
	return (MaterialIcons() + Framework() + Overrides())
}
