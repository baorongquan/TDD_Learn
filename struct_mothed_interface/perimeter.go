package struct_mothed_interface

// Perimeter count the perimeter by getting width and heigth
func Perimeter(width, heigth float64) float64 {
	return 2 * (width + heigth)
}

// Area count the area by getting width and heigth
func Area(width, heigth float64) float64 {
	return width * heigth
}
