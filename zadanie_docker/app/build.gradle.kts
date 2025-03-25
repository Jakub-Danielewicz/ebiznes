plugins {
  kotlin("jvm") version "1.9.24"
  application
}

repositories {
  mavenCentral()
}

dependencies {
  implementation(kotlin("stdlib"))
  implementation(files("libs/sqlite-jdbc.jar"))
}

application {
  mainClass.set("AppKt")
}
