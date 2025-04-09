package org.DiscordBot

import net.dv8tion.jda.api.*
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.hooks.ListenerAdapter
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import kotlinx.coroutines.*

class DiscordBot : ListenerAdapter() {
    private val baseUrl = "http://localhost:8080" // Adres serwera Ktor

    private val client = HttpClient(CIO) // Tworzymy klienta Ktor

    // Musimy użyć coroutines do wykonania zapytania
    override fun onMessageReceived(event: MessageReceivedEvent) {
        val message = event.message
        val content = message.contentRaw

        if (event.author.isBot) return

        if (content.equals("!categories", ignoreCase = true)) {
            // Uruchamiamy funkcję asynchronicznie
            runBlocking {
                val categories = fetchCategories()
                event.channel.sendMessage("Dostępne kategorie:\n$categories").queue()
            }
        }

        if (content.startsWith("!products", ignoreCase = true)) {
            val args = content.split(" ")
            if (args.size == 2) {
                val category = args[1]
                
                // Uruchamiamy funkcję asynchronicznie
                runBlocking {
                    val products = fetchProducts(category)
                    event.channel.sendMessage("Produkty w kategorii $category:\n$products").queue()
                }
            } else {
                event.channel.sendMessage("Poprawne użycie: !products <kategoria>").queue()
            }
        }
    }

    // Funkcja suspend, która wykonuje zapytanie do serwera Ktor
    private suspend fun fetchCategories(): String {
        val response: HttpResponse = client.get("$baseUrl/categories")
        return response.bodyAsText() // Zwracamy odpowiedź jako tekst
    }

    // Funkcja suspend, która wykonuje zapytanie do serwera Ktor
    private suspend fun fetchProducts(category: String): String {
        val response: HttpResponse = client.get("$baseUrl/products/$category")
        return if (response.status.value == 200) {
            response.bodyAsText()
        } else {
            "Kategoria '$category' nie istnieje."
        }
    }
}
