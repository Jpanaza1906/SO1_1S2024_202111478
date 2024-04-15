import json
from random import randrange
from locust import HttpUser, between, task

class readFile():
    def __init__(self):
        self.data = []

    def getData(self): #Metodo donde se obtiene un elemento de la lista de registros
        size = len(self.data) #Tamaño de los datos
        if size > 0:
            index = randrange(0, size - 1) if size > 1 else 0
            return self.data.pop(index)
        else:
            print("size -> 0")
            return None
    
    def loadFile(self):
        print("Cargando ...")
        try:
            with open("traffic.json", 'r', encoding='utf-8') as file:
                self.data = json.loads(file.read())
        except Exception:
            print(f'Error : {Exception}')

class trafficData(HttpUser):
    wait_time = between(0.1, 0.9) #Tiempo de espera entre registros
    reader = readFile()
    reader.loadFile()

    def on_start(self):
        print("On Start")
    
    @task
    def sendMessage(self):
        data = self.reader.getData() #Registro obtenido de la lista
        if data is not None:
            res = self.client.post("/insert", json=data)
            response = res.json()
            print(response)
        else:
            print("Empty") #No hay mas datos por enviar
            self.stop(True)

# python3 -m venv venv
# source venv/bin/activate
# pip install locust
# locust -f tu_script.py -> locust -f traffic.py
# http://localhost:3300 
# http://0.0.0.0:8089




import json
from random import randrange
from locust import HttpUser, between, task

# Class to read data from file
class readFile():
    # Class constructor
    def __init__(self) -> None:
        self.data = []
    
    # Load data from file
    def loadFile(self):
        print(">>> Loading data from file ")
        try:
            with open("traffic.json", "r", encoding="utf-8") as file:
                self.data = json.loads(file.read())
        except Exception as e:
            print(f">>> Error loading file: {e}")
    
    # Get random data
    def getRandomData(self):
        size = len(self.data)
        if size > 0:
            index = randrange(0,size - 1) if size > 1 else 0
            return self.data.pop(index)
        else:
            return None

# Class to simulate traffic
class trafficData(HttpUser):
    wait_time = between(0.1,0.9)
    #Reader object
    reader = readFile()
    reader.loadFile()
    
    def onStart(self):
        print(">>> Starting traffic simulation")
    
    @task
    def sendTraffic(self):
        data = self.reader.getRandomData()
        if data is not None:
            res = self.client.post("/insert", json=data)
            response = res.json()
            print(f">>> Response: {response}")
        else:
            print(">>> No data available")
            self.stop(True)
            
# locust -f tu_script.py -> locust -f traffic.py