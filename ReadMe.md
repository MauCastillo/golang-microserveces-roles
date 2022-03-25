Para realizar la ejecución en local lo primero es instalar la dependencias del go mod
<br>

> go mod tidy

<b>Ejecución de parte 1:</b>
<br>

En este caso he realizado un AWS lambda en el cual la forma más simple para probar su ejecución sería ejecutar los test unitarios, además que actualmente se encuentra desplegado en AWS
<br>
![diagrama_arquitectura](https://github.com/MauCastillo/golang-microserveice-test/blob/main/img/diagrama.png)


Y se puede acceder
https://tqscirqmnf.execute-api.us-west-1.amazonaws.com/staging/?date=2019-12-01&day=1

Donde el parámetro date hace referencia a la fecha
Y el parámetro day al número de días

También puede hacerlo desde la colección de Postman ubicado: 
https://github.com/MauCastillo/golang-microserveice-test/blob/main/parte_1/postman_collection.json

<br>

![postman](https://github.com/MauCastillo/golang-microserveice-test/blob/main/img/postman.png)

<b>Ejecución de la parte 2:</b>
Para el ejecución de debe dirigirse al directorio parte_2 y 
Usar el comando
<br>

> go run main.go

<br>

![executions](https://github.com/MauCastillo/golang-microserveice-test/blob/main/img/ejecucion_script.png)

<br>

Genera un archivo de salida llamando output.json