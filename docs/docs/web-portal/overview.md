---
sidebar_position: 1
---

# Water Treatment Portal General Background

This application utilizes several components to control a water treatment test bed. Using MQTT to control IOT devices and Javascript to control them with a user interface, as well as basic front-end and backend development, this application controls several pieces of hardware that mimic a water treatment plant.

### What is MQTT

MQTT (Message Queuing Telemetry Transport) is a lightweight messaging protocol designed for the Internet of Things (IoT) and other low-bandwidth, high-latency or unreliable networks. It was first developed by Andy Stanford-Clark of IBM and Arlen Nipper of Arcom (now Eurotech) in 1999.

MQTT is a publish-subscribe messaging protocol that enables devices to send and receive small messages (payloads) with minimal overhead. It uses a broker-based architecture, where a central message broker (server) is responsible for receiving messages from publishers and delivering them to subscribers.

MQTT's simplicity and small footprint make it ideal for use in embedded systems and other resource-constrained environments. It has become a popular protocol in the IoT community due to its ability to efficiently handle large numbers of connected devices and its support for reliable, asynchronous communication.

MQTT is used in a variety of applications, including home automation, remote monitoring, industrial automation, and healthcare. It has been adopted by a number of organizations, including IBM, Amazon, Microsoft, and Google, and is an open standard maintained by the OASIS standards body.
Learn More at https://mqtt.org/

### What is Go lang

Go, also known as Golang, is a programming language developed by Google in 2007. It is an open-source language that is designed to be fast, efficient, and easy to learn. Go is a statically typed language, meaning that variable types are explicitly defined at compile time.

Go is designed for concurrency and scalability, making it a popular choice for developing web applications, network services, and other distributed systems. Its built-in concurrency support allows developers to write highly concurrent and parallel programs with ease.

Some of the key features of Go include its garbage collector, its simple and readable syntax, and its cross-platform support. It also has a strong focus on safety and security, with features like strict typing and memory safety.

Overall, Go is a popular language for building efficient and scalable systems, especially in cloud computing and networking environments.
Learn more at https://go.dev/

### What is SVELTE?

Svelte is a free and open-source front-end web application framework that allows developers to build reactive and dynamic user interfaces (UIs) for web applications. It was first introduced in 2016 by Rich Harris, a software developer at The New York Times, and has gained popularity in recent years due to its simplicity and performance.

Unlike other front-end frameworks, which typically rely on a virtual DOM (Document Object Model) to update the UI, Svelte compiles the application code into highly optimized JavaScript code at build time. This means that Svelte applications are generally smaller in size, faster to load, and require less memory than applications built using other front-end frameworks.

Svelte provides a range of features and tools for building modern web applications, including reactive data binding, component-based architecture, CSS encapsulation, and built-in animations. It also integrates seamlessly with other web technologies such as React, Vue.js, and TypeScript, making it a versatile choice for building complex web applications.
Learn more at https://svelte.dev/

### What is Docker?

Docker is a popular open-source platform that enables developers to create, deploy, and run applications in containers. Containers are lightweight, standalone executable packages that contain all the necessary software, libraries, and configuration files needed to run an application.

With Docker, developers can create containers that are portable and can run on any machine, regardless of the underlying hardware, operating system, or hosting environment. This makes it easier to develop and deploy applications across different environments and reduces the risk of compatibility issues.

Docker uses a client-server architecture, where the Docker client communicates with the Docker daemon, which is responsible for managing and running the containers. Developers can use the
Learn more at https://www.docker.com/#

### What is PocketBase?

Pocketbase is an open-source backend that utilizes a real-time embedded SQLite database. It offers file storage, authentication, and an admin dashboard where you can add drop tables and data. Pocketbase is lightweight and useful for creating Rest API’s.

## Running the Application through Docker

To launch the 3 docker container open a terminal and run ( docker compose up )
Click on http://localhost:5173/app

Navigate between different modules and the dashboard

