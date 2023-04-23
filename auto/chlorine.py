import paho.mqtt.client as mqtt
import time
import subprocess

broker_address = '192.168.1.179'
broker_port = 1883

# Previous topic
grit_pump = 'grit/pump'

# Chlorine topic
chlorine_ph_topic = 'chlorine/ph'
chlorine_lvl_topic = 'chlorine/level'
chlorine_pump_topic = 'chlorine/pump'
base_pump_topic = 'basic/punp'

max_water_lvl = 450
max_ph = 4

# Initialize the chlorine level and pH variables
chlorine_lvl = 0
chlorine_ph = 0

def on_connect(client, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(chlorine_lvl_topic)
    client.subscribe(chlorine_ph_topic)

def on_message(client, userdata, msg):
    global chlorine_lvl, chlorine_ph
    if msg.topic == chlorine_lvl_topic:
        chlorine_lvl = int(msg.payload.decode())
    elif msg.topic == chlorine_ph_topic:
        chlorine_ph = int(msg.payload.decode())
    print("Water level: " + str(chlorine_lvl))
    print('pH: ' + str(chlorine_ph))

# Create an MQTT client instance
client = mqtt.Client()

# Assign the on_connect and on_message callback functions
client.on_connect = on_connect
client.on_message = on_message

# Connect to the MQTT broker
client.connect(broker_address, broker_port, 60)

# Start the MQTT client loop
client.loop_start()

while True:
    if client.is_connected():
        client.publish(grit_pump, 'on')
        time.sleep(5)
    else:
        print("Disconnected from MQTT broker")
        break
    
    if chlorine_lvl >= max_water_lvl:
        client.publish(grit_pump, 'off')
    if chlorine_ph < max_ph:
        client.publish(base_pump_topic, 'on')
        time.sleep(2)
        client.publish(base_pump_topic, 'off')
    if chlorine_ph >= max_ph:
        client.publish(base_pump_topic, 'off')
        client.publish(chlorine_pump_topic, 'on')
        subprocess.run(['python', 'dechlorine.py'])
        

client.loop_stop()
