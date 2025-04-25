/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   binGO.go                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: josfelip <josfelip@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2025/04/25 10:50:26 by josfelip          #+#    #+#             */
/*   Updated: 2025/04/25 11:19:32 by josfelip         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: ./binGO arquivo.csv numero_de_ganhadores")
		os.Exit(1)
	}
	
	arquivoCSV := os.Args[1]
	numGanhadores, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Erro: o segundo argumento deve ser um numero")
		os.Exit(1)
	}

	arquivo, err := os.Open(arquivoCSV)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo %s: %v\n", arquivoCSV, err)
		os.Exit(1)
	}
	defer arquivo.Close()

	leitor := csv.NewReader(arquivo)
	registros, err := leitor.ReadAll()

	participantes := []string{}
	for _, registro := range registros {
		if len(registro) > 0 {
			participantes = append(participantes, registro[0])
		}
	}

	if len(participantes) < numGanhadores {
		fmt.Printf("Erro: numero de ganhadores (%d) maior que o numero de \
		participantes (%d)\n", numGanhadores, len(participantes))
		os.Exit(1)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(participantes), func(i, j int) {
		participantes[i], participantes[j] = participantes[j], participantes[i]
	}

	ganhadores := participantes[:numGanhadores]

	for i, ganhador := range ganhadores {
		fmt.Printf("%d - %s\n", i+1, ganhador)
	}
}
