# Maven Builder Image
FROM maven:3.5-jdk-8 as builder
ENV MAVEN_OPTS -XX:+TieredCompilation -Xmx768m
WORKDIR /usr/src/app
COPY pom.xml /usr/src/app
RUN mvn -B -C -T 1C org.apache.maven.plugins:maven-dependency-plugin:3.0.2:go-offline
COPY src /usr/src/app/src
RUN mvn -B -o -T 1C verify

FROM openjdk:jre-alpine as target
COPY --from=builder /usr/src/app/target/app.jar /app.jar
ENTRYPOINT ["java", "-jar", "/app.jar"]
