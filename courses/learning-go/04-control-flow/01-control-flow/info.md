# Entendiendo el control de flujo

La computadora lee los programas de cierta forma, así como nosotros leemos libros de cierta
forma. Cuando, como humanos, leemos libros, en culturas occidentales, lo hacemos de la parte
frontal hacia la trasera del libro, de izquierda a derecha y de arriba hacia abajo. Cuando las
computadoras leen el código de un programa, lo hacen de arriba hacia abajo. Esto es conocido
como lectura en SECUENCIA y a su vez conocido como control de flujo secuencial. Hay otras
dos declaraciones el cuál pueden afectar cómo la computadora lee el código. Una computadora
puede encontrarse con un CICLO o LOOP. Si se encuentra con uno de esos entrará en un
bucle e iterará sobre él de una manera específica. Por eso es también conocido como control
de flujo ITERATIVO. Finalmente, hay también control de flujo CONDICIONAL. Cuando la
computadora se encuentra con una CONDICIÓN, como una "declaración if" o una "declaración
de switch" la computadora tomará acciones diferentes dependiendo de alguna condición.
Entonces las tres formas en la que la computadora lee el código son: SECUENCIAL, LOOP,
CONDICIONAL


- Secuencia
- loop / iterativo
  - for loop
    - init, cond, post
  - bool (con while)
    - infinite
  - Con do-while
    - break
  - contínuo
  - anidado
- condicionales
  - Declaraciones switch / case / default
    - Sin caída predeterminada
    - Creando caída (creating fall-through)
    - Casos múltiples
    - Casos pueden ser expresiones
      - Evaluar a true, ellos corren
    - tipo
  - if
    - El operador de negación
      - !
    - Declaración de inicialización
    - if / else
    - if / else if / else
    - if / else if / else if / … / else
