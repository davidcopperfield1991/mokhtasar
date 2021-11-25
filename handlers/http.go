package handler

import (
	"fmt"
	"net/http"

	"github.com/davidcopperfield1991/mokhtasar/pkg"
	"go.uber.org/zap"
)

type HTTPHandler struct {
	Mokhtasar *pkg.PostgresStore
	Logger    *zap.SugaredLogger
}

func (h *HTTPHandler) Long(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("need a url to short"))
		h.Logger.Errorf("url nadad")
		return
	}
	key, err := h.Mokhtasar.GetOrginalurl(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("nemidunam vaghan"))
		h.Logger.Errorf("daghighan nemidunim chi shod , ino bebin : %v", err)
		return
	}

	toClickUrl := "localhost:8011/long?key=" + key
	h.Logger.Info("inja click kon dadaaaash %v", toClickUrl)
	w.Write([]byte(toClickUrl))

}

func (h *HTTPHandler) Shorten(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("need a short url"))
		return
	}
	key, err := h.Mokhtasar.Shorten(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("nemidunam"))
		h.Logger.Errorf("daghighan nemidunim chi shod , ino bebin : %v", err)
		return
	}
	w.Write([]byte(key))
	fmt.Println("injiye")
}
