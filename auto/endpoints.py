from fastapi import FastAPI
import primary
from primary import subprocess_list, start_mqtt_subscriber
app = FastAPI()

# Get the start time of the script


@app.get('/scripts/starttime')
async def get_start_time():
    start_time = primary.starttime
    return {'start_time': start_time}

# Get the start automation


@app.get('/scripts/start')
async def get_start_script():
    start_mqtt_subscriber()
    return {'message': "Script started"}

# Get the stop automation


@app.get('/scripts/stop')
async def get_stop_script():
    global subprocess_list
    for p in subprocess_list:
        if p.poll() is None:
            p.terminate()
            subprocess_list.remove(p)
    return {'stop_script': "Script stopped"}

# Check if the script is running


@app.get('/scripts/isrunning')
async def is_running():
    global subprocess_list
    is_running = True if subprocess_list else False
    return {'is_running': is_running}
