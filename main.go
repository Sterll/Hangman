package main

import (
	"bufio"
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

var filename string
var playerName string
var word string

//var wordtest list

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

func setScore(scores *Scores, playerName string, newScore int) {
	for i, score := range scores.Scores {
		if score.Name == playerName {
			scores.Scores[i].Score = newScore
			return
		}
	}
}

func main() {
	temp := list.New()
	temp = temp
	file, err := os.Open("word.txt")
	if err != nil {
		fmt.Println("Erreur d'ouverture du fichier:", err)
		return
	}
	defer file.Close()
	var mots []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ligne := scanner.Text()
		mots = append(mots, ligne)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	index := rand.Intn(len(mots))
	word = mots[index]
	fmt.Println("Mot sélectionné pour le pendu:", word)

	filename = "score.json"
	scores, err := readScores(filename)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}
	fmt.Println("Bienvenu sur le jeu du pendu")
	fmt.Println("Veuillez préciser votre pseudo")
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
	welcome()
}

func welcome() {
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
		welcome()
	case 3:
		scores, err := readScores(filename)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier:", err)
			welcome()
		}
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
			scores, err := readScores(filename)
			if err != nil {
				fmt.Println("Sauvegarde de votre score impossible !")
			}
			scoreOfPlayer, err2 := getScore(scores, playerName)
			setScore(&scores, playerName, scoreOfPlayer+1)
			writeScores(filename, scores)
			if err2 != nil {
				fmt.Println("Sauvegarde de votre score impossible !")
			}
			break
		}

		// Dev ici le pendu

	}
}
