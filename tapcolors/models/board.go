package models

import (
	"log"
	"math/rand"
	"strconv"
)

type Game struct {
	Level int
	Board [][]int
	Hash  string
}

const MAX_VALUES = 2

func NewGame(level int) *Game {
	if level < 4 || level > 15 {
		level = 4
	}
	g := &Game{
		Level: level,
	}
	g.Board = make([][]int, g.Level)

	for i := 0; i < len(g.Board); i++ {
		g.Board[i] = make([]int, g.Level)
	}

	return g
}

func NewRandom(level int) *Game {

	g := NewGame(level)

	for i := 0; i < 1000; i++ {
		r := rand.Intn(g.Level * g.Level)
		g.Tap(r/g.Level, r%g.Level)
	}

	return g
	/*
		g := NewGame(5)

		g.Board[2][2] = 1
		g.Board[2][3] = 1
		g.Board[4][2] = 1
		g.UpdateHash()

		return g
	*/
}

func (g *Game) UpdateHash() {
	str := ""
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[i]); j++ {
			str += strconv.Itoa(g.Board[i][j])
		}
	}
	g.Hash = str
}

func (g *Game) Clone() *Game {
	g2 := NewGame(g.Level)
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[i]); j++ {
			g2.set(i, j, g.get(i, j))
		}
	}
	g2.Hash = g.Hash
	return g2
}

func (g *Game) Tap(i, j int) {
	g.set(i, j, g.nextValue(g.get(i, j)))
	g.set(i-1, j, g.nextValue(g.get(i-1, j)))
	g.set(i+1, j, g.nextValue(g.get(i+1, j)))
	g.set(i, j-1, g.nextValue(g.get(i, j-1)))
	g.set(i, j+1, g.nextValue(g.get(i, j+1)))
	g.UpdateHash()
}

func (g *Game) IsValid() bool {
	for i := 0; i < len(g.Board); i++ {
		for j := 0; j < len(g.Board[i]); j++ {
			if g.get(i, j) != 0 {
				return false
			}
		}
	}

	return true
}

func (g *Game) get(i, j int) int {
	if i < 0 || i > (len(g.Board)-1) || j < 0 || j > (len(g.Board[i])-1) {
		return -1
	}

	return g.Board[i][j]
}

func (g *Game) set(i, j, value int) {
	if i < 0 || i > (len(g.Board)-1) || j < 0 || j > (len(g.Board[i])-1) {
		return
	}

	g.Board[i][j] = value
}

func (g *Game) nextValue(i int) int {
	return (i + 1) % MAX_VALUES
}

func (g *Game) FindSolution() []int {
	posibles := make([]*Game, 0)
	posibles = append(posibles, g)

	pasos := make([][]int, 0)
	pasos = append(pasos, make([]int, 0))

	conocidos := make(map[string]struct{}, 0)

	for len(posibles) > 0 && !posibles[0].IsValid() {

		// Pasar a visitado
		tableroActual := posibles[0]
		pasosActuales := pasos[0]

		posibles = posibles[1:]
		pasos = pasos[1:]

		conocidos[tableroActual.Hash] = struct{}{}

		// Agregar todos los hijos
		for i := 0; i < tableroActual.Level; i++ {
			for j := 0; j < tableroActual.Level; j++ {
				nuevoTablero := tableroActual.Clone()
				nuevoTablero.Tap(i, j)

				if _, ok := conocidos[nuevoTablero.Hash]; !ok {
					nuevosPasos := append(pasosActuales, i*g.Level+j+1)
					posibles = append(posibles, nuevoTablero)
					pasos = append(pasos, nuevosPasos)
					conocidos[nuevoTablero.Hash] = struct{}{}
					if len(conocidos)%1000 == 0 {
						log.Println(len(conocidos))
					}
				}
			}
		}
	}

	if len(posibles) > 0 {
		return pasos[0]
	}

	return nil
}

/*
public Vector<Integer> encontrarSolucion() {
	// Tableros a visitar
	Vector<Tablero> posibles = new Vector<Tablero>();
	posibles.add(t);

	// Pasos anteriores
	Vector<Vector<Integer>> pasos = new Vector<Vector<Integer>>();
	Vector<Integer> v = new Vector<Integer>();
	v.addElement(-1);
	pasos.addElement(v);

	// Visitados
	Vector<Tablero> visitados = new Vector<Tablero>();

	while (posibles.size() > 0 && !posibles.firstElement().isValid()) {
	    // Pasar a visitado
	    Tablero tableroActual = posibles.firstElement();
	    visitados.addElement(tableroActual);
	    posibles.removeElementAt(0);

	    Vector<Integer> pasosActuales = pasos.firstElement();
	    pasos.removeElementAt(0);

	    // Agregar todos los hijos
	    for (int i = 0; i < Tablero.MATRIX_SIZE; i++) {
		for (int j = 0; j < Tablero.MATRIX_SIZE; j++) {
		    Tablero nuevoTablero = tableroActual.click(i, j);

		    if (!visitados.contains(nuevoTablero) && !posibles.contains(nuevoTablero)) {
			Vector<Integer> nuevosPasos = (Vector<Integer>) pasosActuales.clone();
			nuevosPasos.add(i * Tablero.MATRIX_SIZE + j + 1);

			posibles.addElement(nuevoTablero);
			pasos.addElement(nuevosPasos);
		    }
		}
	    }
	}

	if (posibles.size() > 0) {
	    Vector<Integer> secuencia = pasos.firstElement();
	    secuencia.removeElementAt(0);
	    return secuencia;
	}

	return null;
		}

*/
