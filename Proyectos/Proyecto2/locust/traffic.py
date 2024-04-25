import json
from random import randrange
import random
from locust import HttpUser, between, task

debug = False

def printDebug(msg):
    if debug:
        print(msg)
        
#Class to read data from file
class Reader():
    #Class constructor
    def __init__(self) -> None:
        self.array = []

    #Load data from file
    def loadFile(self):
        print(">>> Loading data from file ")
        try:
            with open("traffic.json", "r", encoding="utf-8") as data_file:
                self.array = json.loads(data_file.read())
        except Exception as e:
            print(f">>> Error loading file: {e}")
            
    #Pick a random element from the array
    def pickRandom(self):
        length = len(self.array)
        if ( length > 0 ):
            random_index = randrange(0, length - 1) if length > 1 else 0
            return self.array.pop(random_index)
        else:
            return None

# Class to send traffic to the server
class trafficData(HttpUser):
    wait_time = between(0.1, 0.9)
    
    #Reader object
    reader = Reader()
    reader.loadFile()

    def on_start(self):        
        print(">>> Starting traffic simulation")

    @task
    def PostMessage(self):
        random_data = self.reader.pickRandom()

        if ( random_data is not None ):
            data_to_send = json.dumps(random_data)
            printDebug(data_to_send)
            #self.client.post("/grpc/insert", json=random_data)
            #self.client.post("/rust/send_data", json=random_data)
            
            #Elegir aleatoriamente entre las 2 rutas
            route = random.choice(["/grpc/insert", "/rust/send_data"])
            self.client.post(route, json=random_data)
            
        else:
            print(">>> Finished sending data")
            self.stop(True)
            
#locust -f traffic.py
#Ingress = 34.29.26.106