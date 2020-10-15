# PostMarketOS Hardware

This document is for accessing different parts of the hardware for development.

## Battery Sensor

**Pathname:** `/sys/class/power_supply/axp20x-battery`

* `/sys/class/power_supply/axp20x-battery/present` - Battery Present (1 or 0)
* `/sys/class/power_supply/axp20x-battery/capacity` - Battery capacity
* `/sys/class/power_supply/axp20x-battery/health` - Battery health
* `/sys/class/power_supply/axp20x-battery/status` - Battery Status 
* `/sys/class/power_supply/axp20x-battery/voltage_now` - Voltage entering battery
* `/sys/class/power_supply/axp20x-battery/voltage_max_design` - Voltage maximum
* `/sys/class/power_supply/axp20x-battery/voltage_min_design` - Voltage minimum
* `/sys/class/power_supply/axp20x-battery/utype` - All types pertaining to battery

## USB charge sensor

**Pathname** `/sys/devices/platform/soc/1f03400.rsb/sunxi-rsb-3a3/axp20x-usb-power-supply/power_supply/axp20x-usb`

* `/sys/devices/platform/soc/1f03400.rsb/sunxi-rsb-3a3/axp20x-usb-power-supply/power_supply/axp20x-usb/uevent` ALl types pertaining to usb 
* `/sys/devices/platform/soc/1f03400.rsb/sunxi-rsb-3a3/axp20x-usb-power-supply/power_supply/axp20x-usb/present` Is plugged in or not.
* `/sys/devices/platform/soc/1f03400.rsb/sunxi-rsb-3a3/axp20x-usb-power-supply/power_supply/axp20x-usb/online` Is plugged in or not?
* `/sys/devices/platform/soc/1f03400.rsb/sunxi-rsb-3a3/axp20x-usb-power-supply/power_supply/axp20x-usb/health` Powqer supply health 


## GPS

### AT Commands

* `AT+QGPS+1` - Activate GPS
* `AT+QGPSEND` - Deactivate GPS 
* `AT+CGPSINF=0` - Get location information (?). Need to try this with a clear sky.

### Command invocation

Pipe the command to `atinout - /dev/EG25.AT -` (This requires access to the `dialout` group)

## Magnetometer

### Heading

```
# Factory (August 2020 build)
cat /sys/bus/iio/devices/iio:device2/in_magn_y_raw

# September 2020 Build 
 cat /sys/bus/iio/devices/iio\:device3/in_magn_y_raw 
```

## Accelerometer

Currently as of Sep-24-2020 build

`/sys/devices/platform/soc/1c2b00.i2c/i2c-2/2-0068/iio:device2`
or
`/sys/bus/iio/devices/iio:device2`

### Sensors

* `in_accel_x_raw`
* `in_accel_y_raw`
* `in_accel_z_raw`

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

