
# Installation

*epgo* is a command line program run from a shell like Bash. You can find compiled
version in the [releases](https://github.com/caltechlibrary/epgo/releases/latest) 
in the Github repository in a zip file named *epgo-binary-release.zip*. Inside
the zip file look for the directory that matches your computer and copy that someplace
defined in your path (e.g. $HOME/bin). 

Compiled versions are available for Mac OS X (amd64 processor), Linux (amd64), Windows
(amd64) and Rapsberry Pi (both ARM6 and ARM7)

## Mac OS X

1. Go to [github.com/caltechlibrary/epgo/releases/latest](https://github.com/caltechlibrary/epgo/releases/latest)
2. Click on the green "epgo-binary-release.zip" link and download
3. Open a finder window and find the downloaded file and unzip it (e.g. epgo-binary-release.zip)
4. Look in the unziped folder and find dist/macosx-amd64/epgo
5. Drag (or copy) the *epgo* to a "bin" directory in your path
6. Open and "Terminal" and run `epgo -h`

## Windows

1. Go to [github.com/caltechlibrary/epgo/releases/latest](https://github.com/caltechlibrary/epgo/releases/latest)
2. Click on the green "epgo-binary-release.zip" link and download
3. Open the file manager find the downloaded file and unzip it (e.g. epgo-binary-release.zip)
4. Look in the unziped folder and find dist/windows-amd64/epgo.exe
5. Drag (or copy) the *epgo.exe* to a "bin" directory in your path
6. Open Bash and and run `epgo -h`

## Linux

1. Go to [github.com/caltechlibrary/epgo/releases/latest](https://github.com/caltechlibrary/epgo/releases/latest)
2. Click on the green "epgo-binary-release.zip" link and download
3. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/epgo-binary-release.zip)
4. In the unziped directory and find for dist/linux-amd64/epgo
5. copy the *epgo* to a "bin" directory (e.g. cp ~/Downloads/epgo-binary-release/dist/linux-amd64/epgo ~/bin/)
6. From the shell prompt run `epgo -h`

## Raspberry Pi

If you are using a Raspberry Pi 2 or later use the ARM7 binary, ARM6 is only for the first generaiton Raspberry Pi.

1. Go to [github.com/caltechlibrary/epgo/releases/latest](https://github.com/caltechlibrary/epgo/releases/latest)
2. Click on the green "epgo-binary-release.zip" link and download
3. find the downloaded zip file and unzip it (e.g. unzip ~/Downloads/epgo-binary-release.zip)
4. In the unziped directory and find for dist/raspberrypi-arm7/epgo
5. copy the *epgo* to a "bin" directory (e.g. cp ~/Downloads/epgo-binary-release/dist/raspberrypi-arm7/epgo ~/bin/)
    + if you are using an original Raspberry Pi you should copy the ARM6 version instead
6. From the shell prompt run `epgo -h`

