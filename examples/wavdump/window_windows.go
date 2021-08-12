package main

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func GetWinsize() (*Winsize, error) {
	return &Winsize{24, 80, 0, 0}, nil
}
