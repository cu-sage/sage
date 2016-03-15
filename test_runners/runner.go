package test_runner

import (
	"fmt"
	"log"
	"strings"

	"bitbucket.org/sage/models"
	"bitbucket.org/sage/utils"
)

type TestHandler func(test *models.Block, app *models.App) (bool, error)

type TestRunner struct {
	handlers map[string]TestHandler
}

func NewTestRunner() *TestRunner {
	return &TestRunner{
		handlers: map[string]TestHandler{
			"actual_sprite":              sprite,
			"actual_sprite_touch_color":  spriteTouchColor,
			"actual_sprite_touch_sprite": spriteTouchSprite,
			"actual_key_pressed":         whenKeyPressed,
		},
	}
}

func (tr *TestRunner) RunTest(test *models.Block, app *models.App) (*models.TestResult, error) {
	blockType := test.Type
	if blockType != "expect" {
		return nil, fmt.Errorf("Unexpected test type: %s", blockType)
	}

	testType := test.Value.Type
	testHandler, prs := tr.handlers[test.Value.Type]
	if !prs {
		return nil, fmt.Errorf("No handler for test %q", testType)
	}

	pass, err := testHandler(test.Value, app)
	if err != nil {
		log.Printf("Error executing test: %q", err.Error())
		return nil, err
	}

	return &models.TestResult{Pass: pass}, nil
}

func sprite(test *models.Block, app *models.App) (bool, error) {
	return false, nil
}

func spriteTouchColor(test *models.Block, app *models.App) (bool, error) {
	return false, nil
}

func spriteTouchSprite(test *models.Block, app *models.App) (bool, error) {
	return false, nil
}

func whenKeyPressed(test *models.Block, app *models.App) (bool, error) {
	keyPressed := test.Field[0]
	sprite := test.Field[1]

	assertion := test.Value
	should := assertion.Type == "assert_should"

	matcher := assertion.Value

	if matcher.Type == "matcher_point_direction" {
		direction := matcher.Field[0]

		var heading int
		switch direction {
		case "direction_right":
			heading = 90
		case "direction_down":
			heading = 180
		case "direction_left":
			heading = -90
		case "direction_up":
			heading = 0
		default:
			return false, utils.LogAndReturnError("Direction %q not found", direction)
		}

		command := fmt.Sprintf("whenKeyPressed %s heading: %d", keyPressed, heading)
		log.Printf("Test command: %q, Sprite: %q, Should: %v", command, sprite, should)

		return findCommand(command, sprite, should, app), nil
	}

	return false, fmt.Errorf("Matcher type %q not found", matcher.Type)
}

func findCommand(command string, spriteName string, should bool, app *models.App) bool {
	sprite := getSprite(spriteName, app)
	if sprite == nil {
		return false
	}

	for _, script := range sprite.Scripts {
		if strings.Contains(script, command) {
			return true
		}
	}

	if should {
		return false
	}

	return true
}

func getSprite(name string, app *models.App) *models.Sprite {
	for _, sprite := range app.Sprites {
		if sprite.Name == name {
			return sprite
		}
	}

	return nil
}
