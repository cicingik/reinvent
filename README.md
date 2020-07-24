Image Processing with OpenCV in Go
- 
- Install GoCV (https://gocv.io/getting-started/linux/)

        go get -u -d gocv.io/x/gocv
        cd $GOPATH/src/gocv.io/x/gocv
        make install

- Troubleshooting

    If you get error: `Gtk-Message: Failed to load module "canberra-gtk-module"`

        sudo apt-get install libcanberra-gtk-module