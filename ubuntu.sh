#!/bin/sh

sudo apt-get install -y festival xsel golang
GOPATH=/usr/share/go
sudo go get -v -x github.com/lex148/LightRead
sudo go install github.com/lex148/LightRead
mkdir ~/.config/autostart -p

DESKTOP=$( cat <<EOF
[Desktop Entry]\n
Type=Application\n
Exec=/usr/lib/go/bin/LightRead\n
Hidden=false\n
NoDisplay=false\n
X-GNOME-Autostart-enabled=true\n
Name[en_US]=LightRead\n
Name=LightRead\n
Comment[en_US]=\n
Comment=\n
EOF
)

echo $DESKTOP > ~/.config/autostart/LightRead.desktop
