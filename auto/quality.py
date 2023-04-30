import paho.mqtt.client as mqtt
import time
import topics as tp

broker_address = '192.168.1.179'
broker_port = 1883

max_water_lvl = 450
min_temp = 60
max_temp = 90

quality_tmp = 0
quality_lvl = 0


def on_connect(client, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(tp.quality_tmp_topic)
    client.subscribe(tp.quality_lvl_topic)


def on_message(client, userdata, msg):
    global quality_tmp, quality_lvl
    if msg.topic == tp.quality_tmp_topic:
        quality_tmp = int(msg.payload.decode())
    elif msg.topic == tp.quality_lvl_topic:
        quality_lvl = int(msg.payload.decode())
    print("Water level: " + str(quality_lvl))
    print('Temp: ' + str(quality_tmp))


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
        client.publish(tp.dechlorine_pump_topic, 'on')

    else:
        print("Disconnected from MQTT broker")
        break
    if quality_lvl > max_water_lvl:
        client.publish(tp.dechlorine_pump_topic, 'off')
    if quality_tmp < min_temp:
        client.publish(tp.quality_tmp_topic, 'on')

    if quality_tmp >= min_temp:
        client.publish(tp.quality_tmp_topic, 'off')
        time.sleep(1)
        client.publish(tp.intake_topic, 'on')
        time.sleep(1)
        client.publish(tp.intake_topic, 'off')

client.loop_stop()
