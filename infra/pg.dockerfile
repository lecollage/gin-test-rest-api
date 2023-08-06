FROM postgres:latest

# Custom initialization scripts
COPY ./create-user.sh   /docker-entrypoint-initdb.d/10-create_user.sh
COPY ./create-db.sh     /docker-entrypoint-initdb.d/20-create_db.sh