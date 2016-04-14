package baby

// Good returns a good name for a baby depending his/her gender.
func Good(s string) (string, error) {
	switch s {
	case "male":
		return "Joaquin", nil
	case "female":
		return "Matilda", nil
	}
	return "", errors.New("your baby is special")
}
