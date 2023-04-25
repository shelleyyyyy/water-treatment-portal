---
sidebar_position: 2
---

# Application Strucute

## Overview

![Docusaurus logo](/img/web-structure.png)


The web management component of our project is built using a Svelte front-end and a Go web server to handle MQTT connections. We're using PocketBase, a self-hosted SQLite-based database, for device and data storage as well as user management. PocketBase provides a full REST API and JavaScript SDK, as well as a web management platform, making it a great fit for our project's needs.

## Svelte Front End

The Svelte front end of our web management portal is a modern web framework that allowed us to easily create responsive and dynamic user interfaces. It was designed to work well with modern web technologies, such as JavaScript modules and web components, making it easy to integrate with our Go back end. We used Svelte to create a user-friendly dashboard that displays real-time data from our water treatment test bed, including sensor readings and system status information.

## Go Lang Web Server

Our Go web server was responsible for handling MQTT connections between the various components of our water treatment test bed. We chose Go for its fast performance, built-in concurrency features, and strong support for networking. We used the Eclipse Paho MQTT library to communicate with our MQTT broker, and the Go Gin framework to create a robust and scalable web server. The Go server also provided a REST API that allowed us to easily manage users, devices, and data in our PocketBase database.

## PocketBase

PocketBase is a self-hosted database that we used to store device data, user information, and other system data for our water treatment test bed. It is built on top of SQLite, which makes it easy to set up and use, and it provides a full REST API and JavaScript SDK for accessing data. PocketBase also includes a web management platform that we used to view and manage our data, including setting up new devices and users, and managing data storage and retention policies. We found PocketBase to be a flexible and powerful tool for managing the data needs of our project.

Overall, these components formed the backbone of our web management portal, providing a user-friendly and scalable way to manage our water treatment test bed and the various components that make it up.
