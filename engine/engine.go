package engine

import (
	"fmt"
	
	"github.com/pedidosya/@project_name@/models"
)


type Spec interface {
	Hello(name string) string
}

type Engine struct {

}

func New(cfg *models.Configuration) Spec {
	return &Engine{}
}

func (e *Engine) Hello(name string) string {
	return fmt.Sprintf("Hi %s!! how are you?", name)
}