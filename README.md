# HomeKit Computer

_Archived_

Tired of not having enough smart devices to automate with HomeKit?  Look no further because this project turns your (arguably) smartest device into a HomeKit accessory: your computer!

This project builds a daemon that runs on your personal computer to make it visible to HomeKit-enabled services.

The goal of this project is to enable basic control of one's personal computer via HomeKit-enabled devices.  For example, a user with a desktop computer in a room with a HomeKit-enabled light switch may wish to set an automation that, when controlling the switch to turn off the lights, the computer will also sleep or enabled its screensaver.

## Archival Notes

After nothing more than a brief (technical) soirée with this project, I have decided to publicly deprecate and archive it.  It would seem that reducing one's personal computer into device with no more functionality than that of a light bulb is not the best choice one can make.  Unfortunately, that is the case with using the HomeKit Accesory Protocol (HAP).

In the short attempt at development, I received lackluster and infrequent positive results, at best.  Most of the time, the computer did not even broadcast as a HAP accessory and, when it did, it either poorly communicated with my HomeKit controllers (e.g. Home app on iOS) or failed to connect entirely.  In fact, the one time—singular because I did not try this again, nor did I want to—that the computer was discovered and then failed during setup process, it simultaneously disconnected from WiFi and refused to use the "Connected" ethernet connection.  At that point, I needed to restart my computer to restore its networking capabilities.

In an active thread, other users of this project's main dependency, [`hc`](https://github.com/brutella/hc), acknowledge a recurring [issue](https://github.com/brutella/hc/issues/147) with network discovery and broadcast, which I believe also affected my use of the library.

### Hardware Tested

- Model Identifier: iMac18,2
- Processor Name: Quad-Core Intel Core i5
- Processor Speed: 3.4 GHz
- System Firmware Version: 429.60.3.0.0
- SATA Storage: PNY CS900 1TB SSD (_manually installed_)
