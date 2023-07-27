#  golang:alpine di sini adalah basis image yang dimaksud, yaitu image bernama golang dengan tag bernama
FROM golang:alpine

# Command apk update digunakan untuk meng-update index packages pada OS.
# Command apk add --no-cache git digunakan untuk meng-install Git. 
# Kebetulan pada basis image golang:alpine by default Git adalah tidak tersedia. 
# Jadi harus di-install terlebih dahulu.
RUN apk update && apk add --no-cache git

# Create working directory under /app
WORKDIR /app

# Digunakan untuk meng-copy file pada argument pertama yaitu . 
# yang merepresentasikan direktori yang aktif pada host atau 
# komputer kita (yang isinya file main.go, go.mod, dan Dockerfile), 
# untuk kemudian di-paste ke dalam Image ke working directory yaitu /app.
COPY . .


# Digunakan untuk validasi dependensi, dan meng-automatisasi proses download
# jika dependensi yang ditemukan belum ter-download. 
# Command ini akan mengeksekusi go get jika butuh untuk unduh dependensi,
# makanya kita perlu install Git.
RUN go mod tidy


# Command go build digunakan untuk build binary atau executable dari kode program Go.
# Dengan ini source code dalam working directory akan di-build ke executable dengan nama binary
RUN go build -o binary

ENTRYPOINT [ "/app/binary" ]






