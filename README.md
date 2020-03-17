# wingbar-blur
 
## What does it do?

It fakes a blur effect under the wingbar panel on elementary OS by editing your wallpaper.

## Usage

```sh
> wingbar-blur --help

Usage:
  wingbar-blur [OPTIONS]

Application Options:
      --path=    Path to the wallpaper.
  -t, --theme=   Choose theme: none, light, dark. (default: none)
  -o, --opacity= Theme opacity (default: 0.15)
  -b, --blur=    Choose blur intensity (default: 20)
  -s, --size=    Choose size (default: 30)

Help Options:
  -h, --help     Show this help message

```

## Example usage

```sh
wingbar-blur --path ~/Pictures/cool_wallpaper.png --theme light --blur 50 --size 10
```