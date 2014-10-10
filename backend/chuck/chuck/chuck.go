package chuck
 
import (
	"fmt"
	"math/rand"
	"encoding/json"
	"net/http"
	"appengine"
	"appengine/datastore"
)

type Frase struct {
	LaFrase  string
	Usuario  string
	Ip string
}

func init() {
	http.HandleFunc("/", handleStart)
  http.HandleFunc("/frases/", frases)
	http.HandleFunc("/frases/add", frasesAdd)
	http.HandleFunc("/frases/all", frasesall)
}
 
func handleStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Chuck no esta aqui")
}

func initData(c appengine.Context, w http.ResponseWriter, frases []Frase) {

	if len(frases) == 0 {

		answers := []string{
			"Las lágrimas de Chuck Norris curan el cáncer. Es una pena que él no haya llorado nunca.",
			"Chuck Norris no duerme. Espera.",
			"Chuck Norris puede ganar el juego Conecta 4 en sólo 3 movimientos.",
			"Los Dinosaurios miraron mal a Chuck Norris una vez. UNA VEZ.",
			"Chuck Norris ha contado hasta el número infinito… dos veces.",
			"Chuck Norris no caza, porque la palabra caza implica la probabilidad de fracasar. Chuck Norris sale a matar.",
			"Para demostrar que vencer el cáncer no es tan difícil, Chuck Norris se fumó 15 cartones de tabaco al día durante dos años, y desarrolló 7 tipos diferentes de cáncer, sólo para librarse de ellos haciendo flexiones durante 30 minutos. ¡Chúpate ésa, Lance Armstrong!",
			"Chuck Norris es 1/8 Cherokee. No tiene nada que ver con sus antepasados, el tío se comió un puto indio.",
			"En la letra pequeña de la última página del libro de los records Guinness dice que todos los records registrados fueron realizados por Chuck Norris, aquellos que aparecen listados ahí son los que más cerca le llegaron.",
			"Chuck Norris ha demandado a la NBC, alegando que Ley y Orden son marcas registradas para sus piernas derecha e izquierda.",
			"Chuck Norris murió hace 10 años, solo que La Muerte no ha tenido el valor de decírselo.",
		}

		for _,answer := range answers {

			frase:= &Frase{
				LaFrase: answer,
				Usuario:  "Anonimo",
				Ip: "localhost",
			}

			keyFrase := datastore.NewIncompleteKey(c, "Frase", nil)

			if _, err := datastore.Put(c, keyFrase, frase); err != nil {
				c.Errorf("Error al inicializar los valores de google. %v", err)
			}

		}

	}
}

func frases(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	// This query it's not optimized, it should bing only the keys, not all the data
	q := datastore.NewQuery("Frase")

	var frases []Frase
	if _, err := q.GetAll(c, &frases); err != nil {
		c.Errorf("Error al obtener las frases desde el datastore. %v", err)
	}

	initData(c, w, frases)

	b, err := json.Marshal(frases[rand.Intn(len(frases))])
	if err != nil {
		c.Errorf("Error en conversión del Json. %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintln(w, string(b))

}

func frasesall(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	// This query it's not optimized, it should bing only the keys, not all the data
	q := datastore.NewQuery("Frase")

	var frases []Frase
	if _, err := q.GetAll(c, &frases); err != nil {
		c.Errorf("Error al obtener las frases desde el datastore. %v", err)
	}

	initData(c, w, frases)

	b, err := json.Marshal(frases)
	if err != nil {
		c.Errorf("Error en conversión del Json. %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Fprintln(w, string(b))

}

func frasesAdd(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	decoder := json.NewDecoder(r.Body)
	var f Frase
	err := decoder.Decode(&f)
	if err != nil {
		c.Errorf("Error en conversión de Json. %v", err)
		return
	}

	keyFrase := datastore.NewIncompleteKey(c, "Frase", nil)

	if _, err := datastore.Put(c, keyFrase, &f); err != nil {
		c.Errorf("Error al insertar. %v", err)
		fmt.Fprintln(w, "Error al insertar")
		return
	}

	fmt.Fprintln(w, "Se inserto correctamente")

}
