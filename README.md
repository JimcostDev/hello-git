# Hello Git 🚀

¡Bienvenido a mi laboratorio de aprendizaje de Git y GitHub! Este repositorio es mi espacio de práctica donde experimento con diferentes comandos, flujos de trabajo y conceptos de control de versiones.

## 📋 Descripción

Este proyecto es un **laboratorio de práctica** donde aprendo y experimento con:
- Comandos básicos y avanzados de Git
- Flujos de trabajo con GitHub
- Manejo de ramas (branches)
- Resolución de conflictos
- Colaboración y pull requests
- Integración continua

## 🎯 Objetivos de Aprendizaje

- [x] Configuración inicial de Git
- [x] Commits y mensajes descriptivos
- [x] Manejo de archivos y staging area
- [x] Trabajo con ramas (branching)
- [x] Merge
- [x] Resolución de conflictos
- [x] GitHub Actions (CI/CD)
- [x] Colaboración con pull requests

## 🏃‍♂️ Cómo ejecutar

### FastAPI
```bash
# Ejecutar el programa principal
fastapi dev main.py
```

## Cómo correr las pruebas

```bash
pytest test_main.py -v
```

---
## 🛠️ Docker 

### 1. Construir la imagen
Ejecuta en la terminal, estando en la carpeta donde está el Dockerfile:
```bash
docker build -t miapp:local .
```

📌 `-t miapp-go:local` le da un nombre y etiqueta `(local)` a tu imagen.
El `.` al final indica que el contexto es el directorio actual. 

### 2. Verificar que la imagen existe
```bash
docker images
```

### 3. Probar el contenedor
Corre tu imagen y mapea el puerto del contenedor al host:
```bash
docker run --rm -p 8080:8080 miapp:local
```
* `--rm` → elimina el contenedor cuando termine.

* `-p` 8080:8080 → mapea el puerto 8080 del contenedor al 8080 de tu máquina.

Si tu app escucha en 8080, abre en el navegador:
👉 http://localhost:8080

### 4. Eliminar una imagen por su ID o nombre:tag
```bash
docker rmi abc12345defg
```
o
```bash
docker rmi miapp:local
```

### 4. Listar todos los contenedores
```bash
docker ps -a
```

* Eliminar contenedor:
```bash
docker rm <container_id_o_nombre>
```

* Muy útil si quieres borrarlos en bloque:
```bash
docker rm $(docker ps -aq)
```

---
Autor: JimcostDev

