import paho.mqtt.client as mqtt

broker_address = ''
broker_port = 1883


water_lvl_topic = 'primary/level'
scraper_topic = 'primary/scraper'
intake_topic = 'primary/pump'
outtake_topic = 'outtake/pump'
max_water_lvl = 450

# Define the on_connect callback function


def on_connect(client, userdata, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(water_lvl_topic)

# Define the on_message callback function


def on_message(client, userdata, msg):
    water_level = int(msg.payload.decode())
    print("Water level: " + str(water_level))
    if water_level > max_water_lvl:
        print("Turning on outtake pump")
        client.publish(outtake_topic, "on")


# Create an MQTT client instance
client = mqtt.Client()

# Assign the on_connect and on_message callback functions
client.on_connect = on_connect
client.on_message = on_message

# Connect to the MQTT broker
client.connect(broker_address, broker_port, 60)

# Start the MQTT client loop
client.loop_forever()
