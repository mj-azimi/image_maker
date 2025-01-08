# Image Generator in Go

This Go program generates a PNG image with a specified width and height. Each pixel's color is calculated based on its coordinates, creating a gradient effect.

## Features
- Generates a PNG image with customizable dimensions.
- Uses a gradient formula based on pixel coordinates.
- Saves the image to the local filesystem.

## Requirements
- Go (version 1.16 or higher).

## Installation
1. Clone this repository or copy the source code.
2. Ensure you have Go installed on your system.

## Usage
1. Open the `main.go` file.
2. Modify the `width` and `height` variables to set the desired dimensions of the image.
3. Run the program using the following command:

   ```bash
   go run main.go
   ```

4. The generated image will be saved as `image_50x50.png` in the same directory.

## Code Explanation
- **Image Creation**: An RGBA image is created with the specified dimensions using `image.NewRGBA`.
- **Pixel Coloring**: A nested loop iterates through each pixel, setting its color based on the formula:
  ```go
  color.RGBA{
      R: uint8(y/150),
      G: uint8(5),
      B: uint8(x/90),
      A: uint8(250),
  }
  ```
  - `R`: Red component changes with the `y` coordinate.
  - `G`: Green component is fixed at `5`.
  - `B`: Blue component changes with the `x` coordinate.
  - `A`: Alpha (opacity) is fixed at `250`.
- **Image Saving**: The image is saved to a PNG file using the `png.Encode` function.

## Example Output
The generated image will have a gradient effect where the red component varies vertically, and the blue component varies horizontally.

## License
This project is open-source and available under the [MIT License](LICENSE). Feel free to use and modify it as needed.
