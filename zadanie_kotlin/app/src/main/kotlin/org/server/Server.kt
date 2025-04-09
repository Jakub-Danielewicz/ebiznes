package org.server

import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty
import io.ktor.server.application.*
import io.ktor.server.response.respondText
import io.ktor.server.routing.*
import io.ktor.http.HttpStatusCode
import org.service.CategoryService

fun startKtorServer() {
  embeddedServer(Netty, port = 8080) {
    routing {
      get("/categories") {
	call.respondText(CategoryService.getCategories().joinToString("\n"))
      }

      get("/products/{category}") {
	val category = call.parameters["category"]
	if (category != null && CategoryService.getProductsByCategory(category) != null) {
	  call.respondText(CategoryService.getProductsByCategory(category)!!.joinToString("\n"))
	} else {
	  call.respondText("Kategoria nie istnieje!", status = HttpStatusCode.NotFound)
	}
      }
    }
  }.start(wait = false)
}
