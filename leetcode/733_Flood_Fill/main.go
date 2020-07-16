package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]

	if oldColor == newColor {
		return image
	}

	hasSameColor(image, sr, sc, oldColor, newColor)

	return image
}

func hasSameColor(image [][]int, pixelRow, pixelColumn, oldColor, newColor int) {
	if pixelRow >= 0 && pixelRow < len(image) && pixelColumn >= 0 && pixelColumn < len(image[pixelRow]) && image[pixelRow][pixelColumn] == oldColor {
		image[pixelRow][pixelColumn] = newColor

		hasSameColor(image, pixelRow-1, pixelColumn, oldColor, newColor)

		hasSameColor(image, pixelRow, pixelColumn-1, oldColor, newColor)

		hasSameColor(image, pixelRow+1, pixelColumn, oldColor, newColor)

		hasSameColor(image, pixelRow, pixelColumn+1, oldColor, newColor)
	}
}
