# PostMarketOS Hardware

## Camera - Front

`/dev/v4l-subdev1`

### Logging

```
$ v4l2-ctl -d /dev/v4l-subdev0 --list-ctrls

User Controls

error 5 getting ctrl Exposure
                horizontal_flip 0x00980914 (bool)   : default=0 value=0
                  vertical_flip 0x00980915 (bool)   : default=0 value=0

Camera Controls

error 5 getting ext_ctrl Auto Exposure
             auto_exposure_bias 0x009a0913 (intmenu): min=0 max=8 default=4 value=4

Image Source Controls

error 5 getting ext_ctrl Analogue Gain

Image Processing Controls

                   test_pattern 0x009f0903 (menu)   : min=0 max=14 default=0 value=0
error 5 getting ext_ctrl Digital Gain

```

## Camera - Rear

`/dev/v4l-subdev1`

### Logging

```
$ v4l2-ctl -d /dev/v4l-subdev1 --list-ctrls

User Controls

                       contrast 0x00980901 (int)    : min=0 max=255 step=1 default=0 value=0 flags=slider
                     saturation 0x00980902 (int)    : min=0 max=255 step=1 default=64 value=64 flags=slider
                            hue 0x00980903 (int)    : min=0 max=359 step=1 default=0 value=0 flags=slider
        white_balance_automatic 0x0098090c (bool)   : default=1 value=1 flags=update
                    red_balance 0x0098090e (int)    : min=0 max=4095 step=1 default=0 value=0 flags=inactive, slider
                   blue_balance 0x0098090f (int)    : min=0 max=4095 step=1 default=0 value=0 flags=inactive, slider
error 6 getting ctrl Exposure
                 gain_automatic 0x00980912 (bool)   : default=1 value=1 flags=update
error 6 getting ctrl Gain
                horizontal_flip 0x00980914 (bool)   : default=0 value=0
                  vertical_flip 0x00980915 (bool)   : default=0 value=0
           power_line_frequency 0x00980918 (menu)   : min=0 max=3 default=1 value=1

Camera Controls

error 6 getting ext_ctrl Auto Exposure

Image Processing Controls

                     pixel_rate 0x009f0902 (int64)  : min=0 max=2147483647 step=1 default=61430400 value=84000000 flags=read-only
                   test_pattern 0x009f0903 (menu)   : min=0 max=4 default=0 value=0

```

