import paho.mqtt.client as mqtt
import subprocess
import topics as tp


broker_address = '192.168.1.179'
broker_port = 1883



max_water_lvl = 450

def on_connect(client, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(tp.grit_lvl_topic)
        
def on_message(client, userdata, msg):
    global grit_lvl
    grit_lvl = int(msg.payload.decode())
    print("Water level: " + str(grit_lvl))
   
        
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
    if grit_lvl > max_water_lvl:
        print("Max water level reached")
        client.publish(tp.outtake_topic, 'off')
        client.publish(tp.grit_pump_topic, 'on')
        subprocess.Popen(['python', 'chlorine.py'])
        
client.loop_stop()
