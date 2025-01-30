package style

import (
	"fmt"
	"strconv"
	"strings"
)

// Color is a struct that represents a color

type Color struct {
	R int // R is the red value of the color (0-255)
	G int // G is the green value of the color	(0-255)
	B int // B is the blue value of the color (0-255)
	A int // A is the alpha value of the color (0-255) 0 is transparent, 255 is opaque
}

// NewColor creates a new Color but panics if the color is invalid
func NewColor(r, g, b, a int) *Color {
	color := &Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
	err := ColorValidator(*color)
	if err != nil {
		panic(err)
	}
	return color
}

// NewColorWithErr creates a new Color but returns an error if the color is invalid
func NewColorWithErr(r, g, b, a int) (*Color, error) {
	color := &Color{
		R: r,
		G: g,
		B: b,
		A: a,
	}
	err := ColorValidator(*color)
	if err != nil {
		return nil, err
	}
	return color, nil
}

func ColorValidator(c Color) error {
	if c.R < 0 || c.R > 255 {
		return fmt.Errorf("invalid R value, must be between 0 and 255, got %d", c.R)
	}
	if c.G < 0 || c.G > 255 {
		return fmt.Errorf("invalid G value, must be between 0 and 255, got %d", c.G)
	}
	if c.B < 0 || c.B > 255 {
		return fmt.Errorf("invalid B value, must be between 0 and 255, got %d", c.B)
	}
	if c.A < 0 || c.A > 255 {
		return fmt.Errorf("invalid A value, must be between 0 and 255, got %d", c.A)
	}
	return nil
}

func ColorIsValid(c Color) bool {
	if c.R < 0 || c.R > 255 {
		return false
	}
	if c.G < 0 || c.G > 255 {
		return false
	}
	if c.B < 0 || c.B > 255 {
		return false
	}
	if c.A < 0 || c.A > 255 {
		return false
	}
	return true
}

// FromHex creates a Color from a hexadecimal string (e.g., "#RRGGBB" or "#RRGGBBAA")
func FromHex(hex string) (*Color, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 && len(hex) != 8 {
		return nil, fmt.Errorf("invalid hex format, must be 6 or 8 characters, got %d", len(hex))
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid red component in hex: %v", err)
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid green component in hex: %v", err)
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid blue component in hex: %v", err)
	}

	a := 255 // Default alpha value
	if len(hex) == 8 {
		a64, err := strconv.ParseInt(hex[6:8], 16, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid alpha component in hex: %v", err)
		}
		a = int(a64)
		if err != nil {
			return nil, fmt.Errorf("invalid alpha component in hex: %v", err)
		}
	}

	return NewColor(int(r), int(g), int(b), int(a)), nil
}

// ToHex converts the current Color to a hexadecimal string (e.g., "#RRGGBBAA")
func (c *Color) ToHex() string {
	if c.A == 255 {
		// Exclude alpha if it's fully opaque
		return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
	}
	return fmt.Sprintf("#%02X%02X%02X%02X", c.R, c.G, c.B, c.A)
}

// FromHSL creates a Color from HSL values (hue: 0-360, saturation: 0-1, lightness: 0-1)
func FromHSL(h float64, s float64, l float64) (*Color, error) {
	if h < 0 || h > 360 {
		return nil, fmt.Errorf("hue must be between 0 and 360, got %f", h)
	}
	if s < 0 || s > 1 {
		return nil, fmt.Errorf("saturation must be between 0 and 1, got %f", s)
	}
	if l < 0 || l > 1 {
		return nil, fmt.Errorf("lightness must be between 0 and 1, got %f", l)
	}

	c := func(p, q, t float64) float64 {
		if t < 0 {
			t += 1
		}
		if t > 1 {
			t -= 1
		}
		if t < 1.0/6.0 {
			return p + (q-p)*6*t
		}
		if t < 1.0/2.0 {
			return q
		}
		if t < 2.0/3.0 {
			return p + (q-p)*(2.0/3.0-t)*6
		}
		return p
	}

	q := l + s - l*s
	if l < 0.5 {
		q = l * (1 + s)
	}
	p := 2*l - q

	r := c(p, q, h/360.0+1.0/3.0) * 255
	g := c(p, q, h/360.0) * 255
	b := c(p, q, h/360.0-1.0/3.0) * 255

	return NewColor(int(r), int(g), int(b), 255), nil
}

// FromCSSString creates a Color from a CSS color string (e.g., "rgba(255, 0, 0, 0.5)" or "rgb(255, 0, 0)")
func FromCSSString(css string) (*Color, error) {
	css = strings.TrimSpace(css)
	if strings.HasPrefix(css, "rgba") {
		parts := strings.Split(strings.TrimPrefix(strings.TrimSuffix(css, ")"), "rgba("), ",")
		if len(parts) != 4 {
			return nil, fmt.Errorf("invalid rgba format")
		}

		r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, err
		}
		g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, err
		}
		b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			return nil, err
		}
		a, err := strconv.ParseFloat(strings.TrimSpace(parts[3]), 64)
		if err != nil {
			return nil, err
		}
		return NewColor(r, g, b, int(a*255)), nil
	}

	if strings.HasPrefix(css, "rgb") {
		parts := strings.Split(strings.TrimPrefix(strings.TrimSuffix(css, ")"), "rgb("), ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid rgb format")
		}

		r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, err
		}
		g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, err
		}
		b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			return nil, err
		}
		return NewColor(r, g, b, 255), nil
	}

	return nil, fmt.Errorf("unsupported CSS format")
}
