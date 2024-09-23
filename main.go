package main

import "fmt"

func main() {
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
	fmt.Println("Mot sélectionné pour le pendu:", mots[index])
}