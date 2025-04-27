FROM eclipse-temurin:21-jdk-alpine

RUN adduser -D appuser

WORKDIR /app

COPY target/hobbie-backend-0.0.1-SNAPSHOT.jar app.jar

USER appuser

CMD ["java", "-jar", "app.jar"]