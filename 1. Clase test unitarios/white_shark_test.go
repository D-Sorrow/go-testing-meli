package hunt_test

import (
	testify "github.com/stretchr/testify/assert"
	hunt "testdoubles"
	"testing"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 162.89)
		tuna := hunt.NewTuna("Gabriel", 142.89)

		errHunt := shark.Hunt(tuna)

		testify.Nil(t, errHunt)
		testify.Equal(t, shark.Hungry, false)
		testify.Equal(t, shark.Tired, true)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		shark := hunt.NewWhiteShark(false, true, 142.89)
		tuna := hunt.NewTuna("Gabriel", 142.89)

		errHunt := shark.Hunt(tuna)

		testify.Error(t, errHunt, "Error hunt")
		testify.Equal(t, hunt.ErrSharkIsNotHungry, errHunt, "Shark error hungry")
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, true, 142.89)
		tuna := hunt.NewTuna("Gabriel", 142.89)

		errHunt := shark.Hunt(tuna)

		testify.Error(t, errHunt, "Error hunt")
		testify.Equal(t, hunt.ErrSharkIsTired, errHunt, "Shark is tired")

	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 141.89)
		tuna := hunt.NewTuna("Gabriel", 142.89)

		errHunt := shark.Hunt(tuna)
		testify.Error(t, errHunt, "Error hunt")
		testify.Equal(t, hunt.ErrSharkIsSlower, errHunt, "Shark is slower than the tuna")
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 141.89)
		var tuna *hunt.Tuna

		errHunt := shark.Hunt(tuna)

		testify.Error(t, errHunt, hunt.ErrTunaNil, "Tuna is nil")
	})
}
