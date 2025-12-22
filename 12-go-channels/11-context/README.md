# Context

En servidores de Go, cada request entrante es manejado por su propia gorutina. Los request handlers usualmente inicializan gorutinas adicionales para acceder a backends tales como bases de datos o servicios RPC. El conjunto de gorutinas que trabajan en un request típicamente necesitan acceso a valores que son específicos del request, tales como identidad del usuario final, tokens de autorización y un tiempo límite del request. Cuando un request es cancelado o termina su tiempo, todas las gorutinas trabajando para ese request deberían finalizar rápidamente de manera que el sistema pueda reclamar cualquier recurso que estén usando. En Google, desarrollaron un paquete context que hace fácil pasar valores que pertenecen al scope del request, señales de cancelación y deadlines a través de APIs hacia las gorutinas que están involucradas en manejar un request. El paquete está públicamente disponible como context. Este artículo describe cómo usar el paquete, provee un ejemplo completo y funcional.

***Lecturas recomendadas:***
* https://blog.golang.org/context
* https://medium.com/@matryer/context-has-arrived-per-request-state-in-go-1-7-4d095be83bd8
* https://peter.bourgon.org/blog/2016/07/11/context.html

***código:***
* Explorando context
    - background
        - https://play.golang.org/p/S10aI3TXzxI
    - WithCancel
        - Desechando CancelFunc
            - https://play.golang.org/p/XOknf0aSpx
        - Usando CancelFunc
            - https://play.golang.org/p/4ESWl73der5
        - Example
            - https://play.golang.org/p/nzZgwgHurBv
* func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
    - https://play.golang.org/p/WabMOiQnw3a
* Cancelando gorutinas con deadline
* func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
    - https://play.golang.org/p/Q6mVdQqYTt
* Con timeout
* func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
    - https://play.golang.org/p/XwepAlHYf4C
* Con valor
* func WithValue(parent Context, key, val interface{}) Context
    - https://play.golang.org/p/kJAwIpiBnJM