---
sidebar_position: 1
---

# Overview

![Docusaurus logo](/img/digital-twin.png)

We have developed a digital twin of a water treatment facility's IoT devices as a means of practicing cyber red team operations. Our digital twin is a powerful simulation tool that enables us to test and optimize the security posture of the facility and its IoT devices.

Using Mininet, a software-defined network emulator, we have created a virtual network topology that consists of virtual hosts and switches. We store the topology of the network using Neo4j, which allows us to represent the topology as a graph and store additional information about each device, such as the type of device and the MQTT topic it publishes or subscribes to.

The purpose of this digital twin is to simulate realistic cyber attacks on the water treatment facility's IoT devices, and to identify vulnerabilities and weaknesses in the network's security posture. By practicing cyber red team operations in a simulated environment, we can improve our skills and better prepare for real-world scenarios.

To run multiple instances of the digital twin, we use Docker containers and Docker Compose. By isolating each instance in a container, we can ensure that each instance has its own resources and is not impacted by other instances. Using Docker Compose simplifies the process of creating and managing the containers, allowing us to start, stop, and manage the containers as a single unit.

Overall, our digital twin for cyber red team operations is a powerful tool for improving the security posture of a water treatment facility's IoT devices, while leveraging modern technologies like Mininet, Neo4j, and Docker to ensure scalability and efficiency.

