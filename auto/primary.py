import paho.mqtt.client as mqtt
import time
import subprocess
import topics as tp
import time
import threading

subprocess_list = []
starttime = time.time()
broker_address = '192.168.1.179'
broker_port = 1883

max_water_lvl = 450


def on_connect(client, flags, rc):
    print("Connected with result code " + str(rc))
    client.subscribe(tp.water_lvl_topic)


def on_message(client, userdata, msg):
    water_level = int(msg.payload.decode())
    print("Water level: " + str(water_level))



# Define the function to start the MQTT subscriber

def start_mqtt_subscriber():
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
    client.publish(tp.intake_topic, "on")
    print("Intake pump started")

    # Define a function to run the scraper in a separate thread
    def run_scraper():
        while True:
            if client.is_connected():
                client.publish(tp.scraper_topic, "left")
              
                #client.publish(tp.scraper_topic, "right")
            else:
                print("Disconnected from MQTT broker")
                break

    # Start the scraper in a separate thread
    scraper_thread = threading.Thread(target=run_scraper)
    scraper_thread.start()

    # Wait until max water level is reached
    while True:
        if client.is_connected():
        
            # Update the water level variable
            water_level = int(client.subscribe(tp.water_lvl_topic))

        else:
            print("Disconnected from MQTT broker")
            break

        if water_level > max_water_lvl:
            print("Max water level reached")
            client.publish(tp.intake_topic, "off")
            client.publish(tp.scraper_topic, "right")
            client.publish(tp.outtake_topic, "on")
            
            def run_grit():
                p = subprocess.Popen(["python", "grit.py"])
                subprocess_list.append(p)
            
            grit_thread = threading.Thread(target=run_grit)
            grit_thread.start()
            grit_thread.join()

            break

    client.loop_stop()


