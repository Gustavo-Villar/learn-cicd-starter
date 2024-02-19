FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app
COPY .env /app/.env
ADD notely /app/notely
CMD ["./notely"]
