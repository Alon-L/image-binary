# 🎨 Image Binary 👾
Encode and decode binary files into images 

Every PNG pixel is represented by 4 channels: red, green, blue, and the alpha channel, which are represented by an 8-bits unsigned value.  
The encode process, matches every binary byte, found in the binary file, into one of the core 3 channels (red, green and blue).
The decode process, does the opposite: match every channel's value into a binary byte.

## Usage
Here is a list of the available functionalities.
### Encode 🏗️
Encodes a binary file (.WAV, .MOV, .JPEG, etc...) into an image, with the given width and height dimensions.  
Outputs a new PNG image, filled with pixels that correlate to the given binary data.

```shell
$ image-binary encode --width [width] --height [height] [in] [out]
```

### Decode ⛏️
Decodes a PNG file into a binary file (.WAV, .MOV, .JPEG, etc...)  
Outputs a new binary file. Its data correlates to the RGB values of the entered PNG file.

```shell
$ image-binary decode [in] [out]
```
