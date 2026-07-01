# Revisión de Method sets

<strong style="color: red">“El method set de un tipo determina las INTERFACES que ese tipo implementa.....”</strong>

| Receivers | Values    |
| --------- | ------    |
| (t T)     | T and *T  |
| (t *T)    | *T        | 

- Si tienes un *T, puedes llamar a los métodos que tienen un receptor tipo *T, así como a los métodos que tienen un receptor tipo T.
- Si tienes un T y es direccionable (addressable), puedes llamar a métodos que tienen un receptor de tipo *T así como a métodos que tienen un receptor de tipo T, porque la llamada de método t.Meth() será equivalente a (& t). Meth() (Llamadas).
- Si tiene un T y no es direccionable, solo puede llamar a métodos que tienen un tipo de receptor de T, no * T.
- Si tiene una interfaz I, y algunos o todos los métodos en el conjunto de métodos de I
son proporcionados por métodos con un receptor de *T (y el resto es proporcionado por métodos con un receptor de T), entonces *T satisface la interfaz I , pero T no. Esto se debe a que el conjunto de métodos de *T incluye T, pero no al revés (vuelve al primer punto nuevamente).

En resumen, puedes mezclar y emparejar métodos con receptores de valores de tipo T y métodos con receptores de punteros, y usarlos con variables que contengan valores y punteros, sin preocuparse por cuál es cuál. Ambos funcionarán, y la sintaxis es la misma. Sin embargo, si se necesitan métodos con receptores de tipo puntero para satisfacer una interfaz, sólo un puntero será asignable a la interfaz; un valor no será válido.
