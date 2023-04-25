---
sidebar_position: 2
---

# Mininet

Mininet is a powerful software-defined network emulator that enables users to create virtual networks using virtual hosts, switches, and links. It allows users to run real applications and network services on the virtual network, providing a realistic environment for testing and optimization.

To create a virtual network topology for our digital twin, we used Mininet. We created a virtual network that closely mimics the water treatment facility's IoT devices and network infrastructure, allowing us to simulate realistic cyber attacks on the devices.

We used Mininet to create a virtual network topology that consists of virtual hosts and switches. Each virtual device in the topology is represented by a separate Mininet process, enabling us to test and optimize the security posture of each device individually. By isolating each device in its own process, we can ensure that each device has its own resources and is not impacted by other devices.

To store the topology of the network, we used Neo4j, a graph database. Neo4j allows us to represent the network topology as a graph and store additional information about each device, such as the type of device and the MQTT topic it publishes or subscribes to. We used a Python script to pull the topology from Neo4j and populate it with the device types.

The purpose of our digital twin is to simulate realistic cyber attacks on the water treatment facility's IoT devices, and to identify vulnerabilities and weaknesses in the network's security posture. By practicing cyber red team operations in a simulated environment, we can improve our skills and better prepare for real-world scenarios. Using Mininet and Neo4j, we were able to create a powerful simulation tool that enables us to test and optimize the security posture of the facility and its IoT devices.