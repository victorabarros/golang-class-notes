docker rm -f golang-sqlx
docker rmi -f golang-sqlx-image:v0.1
docker build -t golang-sqlx-image:v0.1 .
docker run -d --name golang-sqlx golang-sqlx-image:v0.1
docker logs -f golang-sqlx
