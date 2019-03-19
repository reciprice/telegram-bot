package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

// DefineBasicCommands load basics Telegram functions like /start
func DefineBasicCommands() {
	StartCmd()
}

// DefineCommands load enabled commands
func DefineCommands() {
	ChefCmd()
	RecipeCmd()
}

// StartCmd is the implementation of the welcome message, when you use for the first time RecipriceBot (/start)
func StartCmd() {
	var response string
	response = "Bienvenue !\n" +
		"Je suis programmé pour vous aider dans vos choix de recettes et dans vos courses.\n\n" +
		"Requêtes disponibles:\n" +
		"/chef - En manque d'inspiration ? Demandez la recette du chef !\n" +
		"/recipe <recipe> - Recherche une recette\n" +
		""
	bot.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		bot.Send(m.Sender, response)
	})
}

// ChefCmd is the implementation of the /chef command. Gives you a random recipe. (We're lying to people saying its the choice of the chef.. shhhht.)
func ChefCmd() {
	bot.Handle("/chef", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		recipe := GetRandomRecipe()
		bot.Send(m.Sender, recipe.Video)
		//bot.Send(m.Sender, "J'ai glissé chef !")
	})
}

// RecipeCmd is the implementation of the /recipe command, to search for a recipe and get a shopping list with instructions (to do the meal, not the groceries	)
func RecipeCmd() {
	bot.Handle("/recipe", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		bot.Send(m.Sender, "Commande non implémentée. Revenez plus tard !")
	})
}
