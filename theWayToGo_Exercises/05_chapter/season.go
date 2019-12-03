package main

func Season(m int) string {
	switch {
	case m >= 0 && m <= 4:
		return "season1"
	case m >= 5 && m <= 8:
		return "season2"
	case m >= 9 && m <= 11:
		return "season3"
	default: 
		return "invalid month"
	}
}