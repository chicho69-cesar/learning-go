# Entendiendo canales (channels)

**Introducción a Channels**
- Creando un canal
    - c := make(chan int)
- Colocando un valor en un canal
    - c <- 42
- Tomando un valor de un canal
    - <-c
- buffered channels
    - c := make(chan int, 4)
- Los channels bloquean
    - Son como corredores en una carrera de relevo
        - Están sincronizados
        - Tienen que pasar/recibir el valor al mismo tiempo como los
        corredores en la carrera de relevo tienen que pasar / recibir testigo
        uno al otro al mismo tiempo.
            - Un corredor no puede pasar el bastón en un
            momento dado
            - y luego, un tiempo después, haga que el otro
            corredor reciba el testigo
            - El testigo es pasado / recibido por los corredores al
            mismo tiempo
        - El valor es pasado / recibido sincronizadamente; al mismo
tiempo.
- Los canales nos permiten pasar valores entre gorutinas.