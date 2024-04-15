<b>
Universidad de San Carlos de Guatemala<br>
Facultad de ingenieria<br>
Sistemas Operativos 1<br>
Tutor: Jhonathan Daniel Tocay <br>
Nombre: Jose David Panaza Batres<br>
Carné: 202111478<br>
</b>
<br>
<center>

# TIPOS DE SERVICIOS Y LA INTEGRACIÓN DE KAFKA CON STRIMZI

</center> 

<div style="text-align: justify;">

<center>

### ENSAYO

</center>

La comunicación es un proceso fundamental para cualquier tipo de vida o dispositivo con inteligencia. Los seres humanos hacen uso de esta desde la época prehistórica, donde era indispensable para poder vivir en sociedad y sobrevivir. Con el paso del tiempo los humanos evolucionaron y cambiaron sus lenguajes y maneras de comunicarse. La comunicación cuenta con elementos básicos como el emisor, receptor, mensaje, canal y código. A pesar de su evolución, la comunicación siempre tendrá dichos elementos, ya que sin ellos no tendría sentido comunicarse. Cuando hablamos de aplicaciones de software avanzado, la comunicación entre diferentes servicios o sistemas puede llegar a ser un problema, es por ello que es necesario utilizar contenedores y a su vez un orquestador como kubernetes para facilitar el proceso de despliegue y la comunicación entre servicios.

Kubernetes es una plataforma de código abierto que tiene como función orquestar contenedores. Este ayuda a automatizar la implementación, el escalado y la gestión de contenedores en clústeres de infraestructura, lo que facilita la ejecución de aplicaciones. Los principales beneficios que ofrece Kubernetes son la escalabilidad, alta disponibilidad, gestión de recursos, facilidad de implementación y portabilidad. 

Para trabajar con kubernetes es necesario entender ciertos conceptos que sirven para comprender su funcionamiento. Los pods son los objetos más atómicos que se pueden implementar en kubernetes, uno de estos representa una instancia única de un proceso en ejecución en tu clúster. Un deployment es un objeto que se puede representar como una aplicación de tu cluster. La manera de acceder a los pods se realiza a través de servicios. Estos pueden describir puertos y balanceadores de carga. Otro componente importante es un Ingress, este objeto permite controlar muchos aspectos de la red en el cluster de kubernetes. 

Cuando los contenedores están bien orquestados y controlados, es necesario asegurarse que los datos e información que se mandan, persistan y lleguen a su destino. Para que se cumpla esta cualidad es recomendable utilizar Apache Kafka, esta es una plataforma distribuida para la transmisión de datos que además de permitir las publicaciones, almacenar y procesar flujos de eventos de forma inmediata, esta se suscribe a ellos para no perder dicha información.

Los componentes de kafka son 3: los productores, los consumidores y los brokers. Los productores escriben los mensajes en kafka. Los brokers son nodos que almacenan y distribuyen los datos hacia los consumidores. Los consumidores son encargados de leer y procesar los mensajes.

Para hacer uso de esta herramienta e implementarla de manera sencilla, se puede utilizar Strimzi, siendo este un conjunto de operadores que sirven para ejecutar un cluster de kafka en Kubernetes, permitiendo de forma sencilla diferentes configuraciones de despliegue.

En conclusión el uso de Kubernetes ayuda a administrar una red de servicios de una manera potente. Para reducir la pérdida de información y que esta llegue a su consumidor, es necesario utilizar Apache Kafka, ya que replica la información en varios nodos, y garantiza el almacenamiento de los mensajes antes de que estos sean entregados. Siendo una herramienta muy potente, flexible y si su configuración es buena es confiable. Para contar con los beneficios de este servicio, se utiliza strimzi, ya que este es un operador de Kubernetes que simplifica su implementación y gestión.

<div>

<center>

### CAPTURAS

</center>

- Inicio de la conferencia

![inicio](<inicio.png>)

- Fin de la conferencia

![fin](<fin.png>)
