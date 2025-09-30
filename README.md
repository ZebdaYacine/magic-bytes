# Magic Bytes

Magic Bytes is a Go utility that decodes a base64 string and saves it to a file with the correct file extension. It identifies the file type by inspecting the "magic bytes" of the decoded data.

## How it Works

The program takes a base64 encoded string, decodes it into binary data, and then compares the initial bytes of the data against a list of known file signatures (magic bytes) to determine the file's true type. Once the file type is identified, the program appends the appropriate extension to the output filename and saves the file.

The `main.go` file contains a hardcoded base64 string of a JPG image and saves it as `output.jpg`.

## Usage

To run the program, execute the following command in your terminal:

```bash
go run main.go
```

This will create a file named `output.jpg` in the project's root directory.

