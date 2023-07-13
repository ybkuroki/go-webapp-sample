# Imagen base
FROM golang:1.20

# Establecer directorio de trabajo
WORKDIR /app

# Copiar el código fuente al contenedor
COPY . .

# Compilar la aplicación Go
RUN go build -o app

# Puerto expuesto por la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./app"]