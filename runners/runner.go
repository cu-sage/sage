package runners

import (
	"fmt"
	"log"
	"strings"

	"bitbucket.org/sage/models"
	"bitbucket.org/sage/utils"
)

type TestHandler func(test *models.Block, app *models.App) (*models.TestResult, error)

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

	result, err := testHandler(test.Value, app)
	if err != nil {
		log.Printf("Error executing test: %q", err.Error())
		return nil, err
	}

	if result.Pass {
		result.Actions = processTriggerPass(test.Next)
	} else {
		result.Actions = processTriggerFail(test.Next)
	}

	return result, nil
}

func processTriggerPass(trigger *models.Block) []string {
	return processTrigger(trigger, "trigger_pass")
}

func processTriggerFail(trigger *models.Block) []string {
	return processTrigger(trigger, "trigger_fail")
}

func processTrigger(trigger *models.Block, triggerType string) []string {
	var actions []string
	for ; trigger != nil; trigger = trigger.Next {
		if trigger.Type == triggerType {
			action := trigger.Value

			if action.Type == "action_say" {
				actions = append(actions, action.Field[0])
			}
		}
	}

	return actions
}

func sprite(test *models.Block, app *models.App) (*models.TestResult, error) {
	spriteName := test.Field[0]

	assertion := test.Value
	should := assertion.Type == "assert_should"

	matcher := assertion.Value

	if matcher.Type == "matcher_be_present" {
		log.Printf("Looking for sprite %q", spriteName)

		sprite := getSprite(spriteName, app)

		testResult := &models.TestResult{}
		if should {
			testResult.Description = fmt.Sprintf("Sprite %q should be present", spriteName)
			testResult.Pass = sprite != nil
		} else {
			testResult.Description = fmt.Sprintf("Sprite %q should not be present", spriteName)
			testResult.Pass = sprite == nil
		}

		return testResult, nil
	}

	return nil, fmt.Errorf("Matcher type %q not found", matcher.Type)
}

func spriteTouchColor(test *models.Block, app *models.App) (*models.TestResult, error) {
	spriteName := test.Field[0]
	color := test.Field[1]

	assertion := test.Value
	should := assertion.Type == "assert_should"

	matcher := assertion.Value

	if matcher.Type == "matcher_move_steps" {
		stepCount := matcher.Field[0]

		colorCode := 0
		if color == "red" {
			colorCode = -15399425
		}

		testResult := &models.TestResult{}

		if should {
			testResult.Description = fmt.Sprintf("%q should touch color %q and move %q steps", spriteName, color, stepCount)
		} else {
			testResult.Description = fmt.Sprintf("%q should not touch color %q and move %q steps", spriteName, color, stepCount)
		}

		command := fmt.Sprintf("doIf touchingColor: %d forward: %s", colorCode, stepCount)

		testResult.Pass = findCommand(command, spriteName, should, app)
		return testResult, nil
	}

	return nil, fmt.Errorf("Matcher type %q not found", matcher.Type)
}

func spriteTouchSprite(test *models.Block, app *models.App) (*models.TestResult, error) {
	sprite1 := test.Field[0]
	sprite2 := test.Field[1]

	assertion := test.Value
	should := assertion.Type == "assert_should"

	matcher := assertion.Value

	if matcher.Type == "matcher_say" {
		testResult := &models.TestResult{}

		say := matcher.Field[0]

		if should {
			testResult.Description = fmt.Sprintf("Sprite %q should touch sprite %q and say %q", sprite1, sprite2, say)
		} else {
			testResult.Description = fmt.Sprintf("Sprite %q should not touch sprite %q and say %q", sprite1, sprite2, say)
		}

		command := fmt.Sprintf("touching: %s say:duration:elapsed:from: %s", sprite2, say)
		testResult.Pass = findCommand(command, sprite1, should, app)

		if testResult.Pass {
			return testResult, nil
		}

		command = fmt.Sprintf("touching: %s say:duration:elapsed:from: %s", sprite1, say)
		testResult.Pass = findCommand(command, sprite2, should, app)

		return testResult, nil
	}

	return nil, fmt.Errorf("Matcher type %q not found", matcher.Type)
}

func whenKeyPressed(test *models.Block, app *models.App) (*models.TestResult, error) {
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
			return nil, utils.LogAndReturnError("Direction %q not found", direction)
		}

		testResult := &models.TestResult{}

		if should {
			testResult.Description = fmt.Sprintf("When key %q is pressed, the sprite %q should point in direction %d", keyPressed, sprite, heading)
		} else {
			testResult.Description = fmt.Sprintf("When key %q is pressed, the sprite %q should not point in direction %d", keyPressed, sprite, heading)
		}

		command := fmt.Sprintf("whenKeyPressed %s heading: %d", keyPressed, heading)

		testResult.Pass = findCommand(command, sprite, should, app)
		return testResult, nil
	}

	return nil, fmt.Errorf("Matcher type %q not found", matcher.Type)
}

func findCommand(command string, spriteName string, should bool, app *models.App) bool {
	log.Printf("Looking for command %q in sprite %q with should %v", command, spriteName, should)

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
