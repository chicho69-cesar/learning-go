# Aclaración

- ### **Es todo acerca de tipo**

    Es go un lenguaje orientado a objetos? Go tiene aspectos de OOP. Pero todo es
    acerca de TIPOS. Creamos TIPOS en Go; TIPOS definidos por el usuario.
    Entonces podemos tener VALORES de ese tipo. Podemos asignar VALORES
    de un TIPO definido por el usuario a VARIABLES.

- ### **Go es Orientado a Objetos**
1. Encapsulación
    - Estado ("campos")
    - Comportamiento ("métodos")
    - Exportado & No exportado; visible & no visible
2. Reusabilidad
    - herencia ("tipos embebidos")
1. Polimorfismo
    - interfaces
1. Overriding
    - "promoción"

- ### **OOP Tradicional**
1. Clases
    - Estructura de dato describiendo un tipo de objeto
    - Puedes crear “instancias”/ "objetos" de la clase / prototipo
    - Las clases almacenan ambos:
        * estado / datos / campos
        * comportamiento / métodos
    - público / privado
2. Herencia

- ### **En Go:**
1. No creas clases, creas un TIPO
2. No creas instancias, creas un VALOR de un TIPO

- ### **Tipos definidos por el usuario**
1. Podemos declarar un nuevo tipo
2. foo
    - El tipo subyacente de foo es int
    - Conversión a int
        * int(miEdad)
        * Convirtiendo tipo foo a tipo int
3. ES UNA MALA PRÁCTICA HACER ALIAS DE TIPOS
    - Excepción:
        * Si necesitas asignarle métodos a un tipo
        * Ver el paquete de tiempo para un ejemplo godoc.org/time
            - type Duration int64
            - Duration tiene métodos asignados

- ### **Tipos Nombrados vs Tipos Anónimos**
    * Tipos anónimos son indeterminados. Aún no han sido declarados como
    un tipo. El compilador tiene flexibilidad con tipos anónimos. Puedes
    asignar un tipo anónimo a una variable declarada de cierto tipo. Si la
    asignación puede ocurrir, el compilador hará el trabajo de determinar el
    tipo; el compilador hará una conversión implícita. No puedes asignar un
    tipo nombrado a un tipo de diferente nombre.

- ### **Alineamiento arquitectónico y más**
    * Convención: organiza tus campos lógicamente. Legibilidad y claridad
    ganan en rendimiento como punto crítico. Go será de buen rendimiento.
    Ve primero por legibilidad. Sin embargo, si estás en una situación donde
    necesitas darle prioridad al rendimiento: agrega los campos del más
    grande al de menor tamaño, por ejemplo: int 64, int32, float32, bool