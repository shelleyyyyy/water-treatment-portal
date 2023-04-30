import paho.mqtt.client as mqtt
import subprocess
import threading
import primary
import topics as tp

subprocess_list = primary.subprocess_list
broker_address = '192.168.1.179'
broker_port = 1883
max_grit_lvl = 450


def on_connect(client, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(tp.grit_lvl_topic)


def on_message(client, userdata, msg):
    global grit_lvl
    grit_lvl = int(msg.payload.decode())
    print("Grit level: " + str(grit_lvl))


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
        client.publish(tp.outtake_topic, "on")
    else:
        print("Disconnected from MQTT broker")
        break
    if grit_lvl > max_grit_lvl:
        print("Max grit level reached")
        client.publish(tp.outtake_topic, 'off')
        client.publish(tp.grit_pump_topic, 'on')

        def run_chlorine():
            p = subprocess.Popen(['python', 'chlorine.py'])
            subprocess_list.append(p)

        chlorine_thread = threading.Thread(target=run_chlorine)
        chlorine_thread.start()
        chlorine_thread.join()

        break

client.loop_stop()
