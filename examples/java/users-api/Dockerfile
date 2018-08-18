# Maven Builder Image
FROM maven:3.5-jdk-8 as builder
COPY src /usr/src/app/src
COPY pom.xml /usr/src/app
ENV MAVEN_OPTS -XX:+TieredCompilation -Xmx768m
RUN mvn -f /usr/src/app/pom.xml -T 1C clean package

FROM openjdk:jre-alpine as target
COPY --from=builder /usr/src/app/target/app.jar /app.jar
ENTRYPOINT ["java", "-jar", "/app.jar"]
