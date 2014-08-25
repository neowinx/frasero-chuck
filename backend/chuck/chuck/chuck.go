package nhemu
 
import (
	"fmt"
	"net/http"
	"math/rand"
	/*"encoding/json"
	"appengine"
	"appengine/datastore"
	"strconv"*/
)
 
func init() {
	http.HandleFunc("/", handleStart)
    http.HandleFunc("/frases/", frases)
}
 
func handleStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Chuck no esta aqui")
}

func frases(w http.ResponseWriter, r *http.Request) {
	//rand.Seed(13) // Try changing this number!
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

	fmt.Fprintln(w, answers[rand.Intn(len(answers))])
}
