import paho.mqtt.client as mqtt
import time
import subprocess
import threading
import primary
import topics as tp

subprocess_list = primary.subprocess_list
broker_address = '192.168.1.179'
broker_port = 1883


max_water_lvl = 450
max_ph = 9

# Initialize the dechlorine level and pH variables
dechlorine_lvl = 0
dechlorine_ph = 0


def on_connect(client, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(tp.dechlorine_ph_topic)
    client.subscribe(tp.dechlorine_lvl_topic)


def on_message(client, userdata, msg):
    global dechlorine_lvl, dechlorine_ph
    if msg.topic == tp.dechlorine_lvl_topic:
        dechlorine_lvl = int(msg.payload.decode())
    elif msg.topic == tp.dechlorine_ph_topic:
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
        client.publish(tp.chlorine_pump_topic, 'on')

    else:
        print("Disconnected from MQTT broker")
        break
    if dechlorine_lvl > max_water_lvl:
        client.publish(tp.chlorine_pump_topic, 'off')
    if dechlorine_ph > max_ph:
        client.publish(tp.acid_pump_topic, 'on')

    if dechlorine_ph <= max_ph:
        client.publish(tp.acid_pump_topic, 'off')
        client.publish(tp.agitator_topic, 'on')
        # time.sleep(1)
        client.publish(tp.agitator_topic, 'off')
        client.publish(tp.dechlorine_pump_topic, 'on')

        def run_quality():
            p = subprocess.Popen(['python', 'quality.py'])
            subprocess_list.append(p)

        quality_thread = threading.Thread(target=run_quality)
        quality_thread.start()
        quality_thread.join()


client.loop_stop()
