---
sidebar_position: 1
---

# Hardware

The hardware part of our project involves a variety of devices, including water pumps, solenoid valves, level sensors, flow sensors, pressure sensors, and pH sensors. These devices are connected to a control panel that interfaces with a Raspberry Pi running our custom Python automation scripts. The control panel includes physical switches and indicators that allow for manual control of the devices in case of automation failure. Sensor data and device commands are sent through an MQTT broker, allowing for real-time monitoring and control of the water treatment process. The devices are also connected to an industrial-grade Ethernet switch, which provides reliable communication between the devices and the Raspberry Pi. The entire setup is powered by a dedicated power supply, providing a stable voltage and current for the devices.