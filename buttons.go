package vrcinput

import (
	"errors"
	"time"

	"github.com/hypebeast/go-osc/osc"
)

// Buttons expect an int of 1 for 'pressed' and 0 for 'released'.
// They will not function correctly without resetting to 0 first - sending /input/Jump 1 and then /input/Jump 1 will only result in a single jump.

// Arguments require a 32 bit types (f32, int32, bool)

type MoveDirection string

const (
	MoveForward  MoveDirection = "MoveForward"
	MoveBackward MoveDirection = "MoveBackward"
	MoveLeft     MoveDirection = "MoveLeft"
	MoveRight    MoveDirection = "MoveRight"
)

// Move sends a move direction and a boolean value to the OSC server.
// It takes a MoveDirection and a boolean value as parameters.
// The boolean value indicates whether to enable or disable the move direction.
// It returns an error if there was a problem sending the message.
func (c *localOscClient) Move(direction MoveDirection, b bool) error {
	message := osc.NewMessage("/input/"+string(direction), b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

type LookDirection string

const (
	LookLeft  LookDirection = "LookLeft"
	LookRight LookDirection = "LookRight"
)

// Look sends an OSC message to control the look direction.
// It takes a LookDirection and a boolean value as parameters.
// The boolean value indicates whether to enable or disable the look direction.
// It returns an error if there was a problem sending the message.
func (c *localOscClient) Look(direction LookDirection, b bool) error {
	message := osc.NewMessage("/input/"+string(direction), b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

// Jump if the world supports it.
func (c *localOscClient) Jump() error {
	var jumping bool
	for i := 0; i < 2; i++ {
		jumping = !jumping
		message := osc.NewMessage("/input/Jump", jumping)
		err := c.Send(message)
		if err != nil {
			return errors.New("failed to send message: " + err.Error())
		}

		// Sometimes it doesn't work if we don't sleep for a bit
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

// Run if the world supports it.
func (c *localOscClient) Run(b bool) error {
	message := osc.NewMessage("/input/Run", b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

type ComfortLookDirection string

const (
	ComfortLookLeft  ComfortLookDirection = "ComfortLeft"
	ComfortLookRight ComfortLookDirection = "ComfortRight"
)

// ComfortLook works the same as Look, but for comfort turning using snap turns. VR Only.
func (c *localOscClient) ComfortLook(direction ComfortLookDirection, b bool) error {
	message := osc.NewMessage("/input/"+string(direction), b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

type DropHand string

const (
	DropLeftHand  DropHand = "DropLeft"
	DropRightHand DropHand = "DropRight"
)

// DropHand drops the item in the hand of the specified side.
// It takes a DropHand and a boolean value as parameters.
// The boolean value indicates whether to enable or disable the drop hand action.
func (c *localOscClient) DropHand(hand DropHand, b bool) error {
	message := osc.NewMessage("/input/"+string(hand), b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

type UseHand string

const (
	UseLeftHand  UseHand = "UseLeft"
	UseRightHand UseHand = "UseRight"
)

// UseHand uses the item in the hand of the specified side.
// It takes a UseHand and a boolean value as parameters.
// The boolean value indicates whether to enable or disable the use hand action.
func (c *localOscClient) UseHand(hand UseHand, b bool) error {
	message := osc.NewMessage("/input/"+string(hand), b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

type GrabHand string

const (
	GrabLeftHand  GrabHand = "GrabLeft"
	GrabRightHand GrabHand = "GrabRight"
)

// GrabHand grabs the item in the hand of the specified side.
// It takes a GrabHand and a boolean value as parameters.
// The boolean value indicates whether to enable or disable the grab hand action.
func (c *localOscClient) GrabHand(hand GrabHand, b bool) error {
	message := osc.NewMessage("/input/"+string(hand), b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

// Turn off and on Safe Mode. 
func (c *localOscClient) PanicButton(b bool) error {
	message := osc.NewMessage("/input/PanicButton", b)
	err := c.Send(message)
	if err != nil {
		return errors.New("failed to send message: " + err.Error())
	}
	return nil
}

// TODO: Voice, Quick Menu
