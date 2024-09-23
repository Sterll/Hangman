package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type STATE int64

const (
	WAITING STATE = iota
	PLAYING
	END
)

var GAME_STATE STATE

type Score struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}
type Scores struct {
	Scores []Score `json:"scores"`
}

func readScores(filename string) (Scores, error) {
	file, err := os.Open(filename)
	if err != nil {
		return Scores{}, err
	}
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var scores Scores
	err = json.Unmarshal(byteValue, &scores)
	if err != nil {
		return Scores{}, err
	}

	return scores, nil
}

// Fonction pour écrire des scores dans le fichier JSON
func writeScores(filename string, scores Scores) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(scores, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	return err
}

func getScore(scores Scores, playerName string) (int, error) {
	for _, score := range scores.Scores {
		if score.Name == playerName {
			return score.Score, nil
		}
	}
	return 0, fmt.Errorf("Score non trouvé pour le joueur %s", playerName)
}

func main() {
	filename := "score.json"
	scores, err := readScores(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	fmt.Println("Veuillez préciser votre pseudo")
	var playerName string
	fmt.Scan(&playerName)
	scoreOfPlayer, err2 := getScore(scores, playerName)
	if err2 != nil {
		newScore := Score{
			Name:  playerName,
			Score: scoreOfPlayer,
		}
		scores.Scores = append(scores.Scores, newScore)
		writeScores(filename, scores)
	}

	fmt.Println("Bienvenu sur le jeu du pendu")
	fmt.Println("Veuillez choisir une option ci-dessous \n\n 1. Lancer une partie \n 2. Voir les règles du jeu \n 3. Voir les scores \n 4. Quitter le jeu \n")

	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
	case 2:
		fmt.Println("Voici les règles du jeu : \n")
		fmt.Println("Un mot sera choisis aléatoirement parmis une liste de mots prédéfinis,")
		fmt.Println("Vous aurez 10 vies maximum pour trouver celui-ci,")
		fmt.Println("A chaque début de partie, trois lettres du mots vous seront données. \n")
		fmt.Println("Bonne partie !")
		main()
	case 3:
		fmt.Println("Voici les scores : \n")
		for _, sc := range scores.Scores {
			fmt.Println(sc.Name, ":", sc.Score)
		}
		main()
	case 4:
		fmt.Println("Vous pourrez retenter votre chance plus tard !")
		os.Exit(0)
	}
}

func play() {
	for {
		if GAME_STATE == END {
			break
		}
	}
}
