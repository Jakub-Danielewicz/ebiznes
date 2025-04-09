package org.service

object CategoryService {
  private val categories = mapOf(
    "Guitars" to listOf("Gibson SG", "Gibson Les Paul", "Fender Stratocaster", "Fender Mustang"),
    "Amplifiers" to listOf("Peavey 5150", "Mesa Boogie Mark V", "Marshall JCM800", "Vox AC30"),
    "Guitar Effects" to listOf("Digitech Drop", "Digitech Whammy", "Dunlop CGB95", "Strymon Mobius")
  )
  
  fun getCategories(): List<String> {
   return categories.keys.toList()
  }
  fun getProductsByCategory(category: String): List<String>? {
    return categories[category]
  }
}

