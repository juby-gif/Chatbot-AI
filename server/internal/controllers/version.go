package controllers

import (
	"net/http"
)

func (c *Controller) getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Chatbot-AI_v1.0"))
}
