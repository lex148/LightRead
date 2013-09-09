Light Read
========

####A utility that will text to speech all selected text when Ctrl-z is pressed


## Usage
1. Select some text
2. Press Control-z

or press control-q to read the text in your clipboard 



## Installation:

```bash
pacaur -S light-read-git
```

## Setup:
Light Read should be started when you login. This can be done in gnome with gnome-session-properties
or adding the following line to the top of your .xinitrc file before
```bash
LightRead &
```

## Dependencies 
```bash
xsel
xclip
festival
at least one festival voice
aplay (already on most systems)
```

