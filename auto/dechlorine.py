import paho.mqtt.client as mqtt
import time
import subprocess
import topics


broker_address = '192.168.1.179'
broker_port = 1883

# Previous topic
chlorine_pump_topic = 'chlorine/pump'

# Dechlorine topics
dechlorine_ph_topic = 'dechlorine/ph'
dechlorine_lvl_topic = 'dechlorine/level'
dechlorine_pump_topic = 'dechlorine/pump'
acid_pump_topic = 'acid/pump'
agitator_topic = 'quality/plug/agitator'

max_water_lvl = 450
max_ph = 9

# Initialize the dechlorine level and pH variables
dechlorine_lvl = 0
dechlorine_ph = 0

def on_connect(client, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(dechlorine_ph_topic)
    client.subscribe(dechlorine_lvl_topic)

def on_message(client, userdata, msg):
    global dechlorine_lvl, dechlorine_ph
    if msg.topic == dechlorine_lvl_topic:
        dechlorine_lvl = int(msg.payload.decode())
    elif msg.topic == dechlorine_ph_topic:
        dechlorine_ph = int(msg.payload.decode())
    print("Water level: " + str(dechlorine_lvl))
    print('pH: ' + str(dechlorine_ph))
        
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
        client.publish(chlorine_pump_topic, 'on')
        time.sleep(5)
    else:
        print("Disconnected from MQTT broker")
        break
    if dechlorine_lvl > max_water_lvl:
        client.publish(chlorine_pump_topic, 'off')
    if dechlorine_ph > max_ph:
        client.publish(acid_pump_topic, 'on')
        time.sleep(2)
    if dechlorine_ph <= max_ph:
        client.publish(acid_pump_topic, 'off')
        client.publish(agitator_topic, 'on')
        time.sleep(10)
        client.publish(agitator_topic, 'off')
        client.publish(dechlorine_pump_topic, 'on')
        subprocess.Popen(['python', 'quality.py'])

client.loop_stop()
        

