# Sistema de Notificaciones

## Descripción
Este proyecto gestiona y envía notificaciones a usuarios, asegurándose de respetar ciertos límites configurables.

## Servicios

- **NotificationsService**: Es el servicio principal que gestiona las notificaciones y verifica si se cumplen los límites.
- **StatusNotification, NewsNotification, ...**: Estos son diferentes implementaciones que permiten enviar notificaciones de acuerdo a su tipo.

## Utilidades (en `pkg`)

- **ConvertToSeconds**: Función que convierte distintas unidades de tiempo a segundos.
- **UserCache**: Un cache en memoria que guarda información sobre los usuarios.

Se proporcionan pruebas para diferentes componentes del proyecto, que se pueden encontrar en el directorio `tests`.
