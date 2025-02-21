package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testdoubles/internal/hunter"
	"testdoubles/internal/positioner"
	"testdoubles/internal/prey"
	"testdoubles/internal/simulator"
	"testdoubles/platform/web/response"
)

// NewHunter returns a new Hunter handler.
func NewHunter(ht hunter.Hunter, pr prey.Prey) *Hunter {
	return &Hunter{ht: ht, pr: pr}
}

// Hunter returns handlers to manage hunting.
type Hunter struct {
	// ht is the Hunter interface that this handler will use
	ht hunter.Hunter
	// pr is the Prey interface that the hunter will hunt
	pr prey.Prey
}

// RequestBodyConfigPrey is an struct to configure the prey for the hunter in JSON format.
type RequestBodyConfigPrey struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigurePrey configures the prey for the hunter.
func (h *Hunter) ConfigurePrey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var tuna RequestBodyConfigPrey
		err := json.NewDecoder(r.Body).Decode(&tuna)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}
		// process
		h.pr.Configure(tuna.Speed, tuna.Position)
		// response
		response.JSON(w, http.StatusOK, "prey configured")
		return
	}
}

// RequestBodyConfigHunter is an struct to configure the hunter in JSON format.
type RequestBodyConfigHunter struct {
	Speed    float64              `json:"speed"`
	Position *positioner.Position `json:"position"`
}

// ConfigureHunter configures the hunter.
func (h *Hunter) ConfigureHunter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		var hunter RequestBodyConfigHunter
		err := json.NewDecoder(r.Body).Decode(&hunter)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}
		// process
		h.ht.Configure(hunter.Speed, hunter.Position)
		err = errors.New("body request invalid")
		// response
		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}
		response.JSON(w, http.StatusOK, "hunter configured")
		return
	}
}

// Hunt hunts the prey.
func (h *Hunter) Hunt() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		simul := simulator.NewCatchSimulatorDefault(&simulator.ConfigCatchSimulatorDefault{
			Positioner:     positioner.NewPositionerDefault(),
			MaxTimeToCatch: 100,
		})
		subHunt := &simulator.Subject{
			Speed: 100.0,
			Position: &positioner.Position{
				X: 0.0,
				Y: 0.0,
				Z: 0.0,
			},
		}
		subPrey := &simulator.Subject{
			Speed:    h.pr.GetSpeed(),
			Position: h.pr.GetPosition(),
		}
		// process
		duration, isCatch := simul.CanCatch(subHunt, subPrey)
		res := fmt.Sprintf("hunt done. Time: %1f", duration)
		if !isCatch {
			response.JSON(w, http.StatusOK, res)
			return
		}

		// response
		response.JSON(w, http.StatusOK, res)
		return
	}
}
