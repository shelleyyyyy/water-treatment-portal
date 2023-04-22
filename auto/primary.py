import paho.mqtt.client as mqtt
import time
import subprocess


broker_address = '192.168.1.179'
broker_port = 1883

# primary topics
water_lvl_topic = 'primary/level'
scraper_topic = 'primary/scraper'
intake_topic = 'primary/pump'
outtake_topic = 'outtake/pump'

max_water_lvl = 450

# Define the on_connect callback function
def on_connect(client, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(water_lvl_topic)

# Define the on_message callback function
def on_message(client, userdata, msg):
    water_level = int(msg.payload.decode())
    print("Water level: " + str(water_level))
  

# Create an MQTT client instance
client = mqtt.Client()

# Assign the on_connect and on_message callback functions
client.on_connect = on_connect
client.on_message = on_message

# Connect to the MQTT broker
client.connect(broker_address, broker_port, 60)

# Start the MQTT client loop
client.loop_start()

# Start the intake pump
client.publish(intake_topic, "on")
print("Intake pump started")

# Wait until max water level is reached
while True:
    if client.is_connected():
        client.publish(intake_topic, "on")
        client.publish(scraper_topic, "left")
        time.sleep(10)
        client.publish(scraper_topic, "right")
        time.sleep(10)
        client.publish(outtake_topic, "off")
        time.sleep(1)
        
        # Update the water level variable
        water_level = int(client.subscribe(water_lvl_topic))
        
    else:
        print("Disconnected from MQTT broker")
        break

    if water_level > max_water_lvl:
        print("Max water level reached")
        client.publish(intake_topic, "off")
        client.publish(scraper_topic, "right")
        client.publish(outtake_topic, "on")
        subprocess.Popen(["python", "grit.py"])
        break


# Stop the MQTT client loop
client.loop_stop()
