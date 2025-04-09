package org

import org.server.startKtorServer
import org.DiscordBot.DiscordBot
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.requests.GatewayIntent
import io.github.cdimascio.dotenv.dotenv



fun main() {
  startKtorServer()

  val dotenv = dotenv()
  val token = dotenv["DISCORD_TOKEN"] ?: throw IllegalArgumentException("Token not found") 
  val jda = JDABuilder.createLight(token, GatewayIntent.GUILD_MESSAGES, GatewayIntent.DIRECT_MESSAGES, GatewayIntent.MESSAGE_CONTENT)
  .addEventListeners(DiscordBot())
  .build()
}
