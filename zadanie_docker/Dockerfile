FROM ubuntu:latest


RUN apt-get update && apt-get install -y \
    software-properties-common \
    && add-apt-repository -y ppa:deadsnakes/ppa \
    && apt-get update && apt-get install -y \
    python3.10  python3-pip \
    openjdk-8-jdk wget unzip \
    && rm /usr/bin/python3 \
    && ln -s /usr/bin/python3.10 /usr/bin/python3



RUN wget -O kotlin-compiler.zip https://github.com/JetBrains/kotlin/releases/download/v1.9.24/kotlin-compiler-1.9.24.zip \
    && mkdir /opt/kotlin \
    && unzip kotlin-compiler.zip -d /opt/kotlin \
    && ln -s /opt/kotlin/bin/kotlinc /usr/bin/kotlinc \
    && ln -s /opt/kotlin/bin/kotlin /usr/bin/kotlin \
    && rm kotlin-compiler.zip



RUN wget -O gradle.zip https://services.gradle.org/distributions/gradle-8.13-bin.zip \
    && mkdir /opt/gradle \
    && unzip gradle.zip -d /opt/gradle \
    && ln -s /opt/gradle/gradle-8.13/bin/gradle /usr/bin/gradle \
    && rm gradle.zip

RUN mkdir -p /app/libs \
    && wget -O /app/libs/sqlite-jdbc.jar https://repo1.maven.org/maven2/org/xerial/sqlite-jdbc/3.41.2.1/sqlite-jdbc-3.41.2.1.jar

WORKDIR /app
COPY app /app

RUN gradle build

CMD ["gradle", "run"]
