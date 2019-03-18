package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

func DefineBasicCommands() {
	StartCmd()
}

func DefineCommands() {
	ChefCmd()
	RecipeCmd()
}

func StartCmd() {
	var response string
	response = "Bienvenue !\n" +
		"Je suis programmé pour vous aider dans vos choix de recettes et dans vos courses.\n\n" +
		"Commandes disponibles:\n" +
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

func RecipeCmd() {
	bot.Handle("/recipe", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		bot.Send(m.Sender, "Commande non implémentée. Revenez plus tard !")
	})
}
