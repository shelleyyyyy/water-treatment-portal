---
sidebar_position: 3
---

# Neo4j

![Docusaurus logo](/img/neo4j.png)

## What is Neo4j

Neo4j is a graph database management system that is designed to store, manage, and query large and complex datasets as graphs. It uses the property graph data model, which consists of nodes, relationships, and properties, to represent data and their connections.

Nodes represent entities, such as people, places, and things, and relationships represent the connections between them. Properties are key-value pairs that provide additional information about nodes and relationships. Together, these elements form a graph that allows for flexible and efficient querying of complex and interconnected data.

Neo4j is particularly useful for applications that require complex relationship queries, such as social networks, recommendation systems, and fraud detection. It is open-source software and is widely used in various industries, including finance, healthcare, and e-commerce.

## How are we using it?

We have built a Mininet simulator that uses Neo4j to store the topology of the network. Our simulator allows us to create a virtual network topology using Mininet, and test various network scenarios. To store the topology of the network, we have utilized Neo4j's graph database management system, which allows us to represent the network topology as a graph. In this graph, nodes represent devices in the network and relationships represent the connections between them. Additionally, we have used properties to store additional information about each device, such as the type of device and the MQTT topic it publishes or subscribes to.

Using Neo4j to store the topology of our network has several advantages. First, it provides us with flexible and efficient querying of the graph, which can be useful for various network operations such as finding the shortest path between two devices or identifying devices that are causing network congestion. Second, Neo4j offers a scalable solution that can handle large and complex network topologies. Finally, Neo4j's ability to represent and store graph data makes it a natural fit for storing network topologies, which can be complex and interconnected.

Overall, our Mininet simulator that uses Neo4j to store the topology is a powerful tool for testing and simulating network environments, and provides us with a flexible and efficient way to store and query network data.

## Official Site

https://neo4j.com/