package parsers

import (
	"bytes"
	"fmt"
	"strconv"

	"bitbucket.org/sage/models"
	"bitbucket.org/sage/utils"
)

func ParseSB2(project models.SB2Project) (*models.App, error) {
	app := models.NewApp()
	for _, child := range project.Children {
		sprite := models.NewSprite(child.ObjName)

		for i := 0; i < len(child.Scripts); i++ {
			for j := 2; j < len(child.Scripts[i]); j++ {
				buffer := bytes.NewBufferString("")

				if blocks, ok := child.Scripts[i][j].([]interface{}); ok {
					for k := 0; k < len(blocks); k++ {
						block, err := parseBlock(blocks[k])
						if err != nil {
							return nil, err
						}

						buffer.WriteString(block)
						buffer.WriteString(" ")
					}
				}

				sprite.AddScript(buffer.String())
			}
		}

		app.AddSprite(sprite)
	}

	return app, nil
}

func parseBlock(toParse interface{}) (string, error) {
	block, ok := toParse.([]interface{})
	if !ok {
		return "", utils.LogAndReturnError("Not a valid block")
	}

	if len(block) == 0 {
		return "", utils.LogAndReturnError("The block is empty")
	}

	// Addresses weird SB2 format where opcode is another block
	if innerBlock, ok := block[0].([]interface{}); ok {
		return parseBlock(innerBlock)
	}

	opcode, err := getOpcode(block[0])
	if err != nil {
		return "", utils.LogAndReturnError("Cannot retrieve opcode")
	}

	buffer := bytes.NewBufferString(opcode)
	for i := 1; i < len(block); i++ {
		arg, err := getArgument(block[i])
		if err != nil {
			return "", utils.LogAndReturnError("Cannot retrieve %d argument", i)
		}

		buffer.WriteString(" ")
		buffer.WriteString(arg)
	}
	return buffer.String(), nil
}

func getOpcode(opcode interface{}) (string, error) {
	s, ok := opcode.(string)
	if !ok {
		fmt.Println(opcode)
		return "", utils.LogAndReturnError("Opcode is not a string")
	}

	return s, nil
}

func getArgument(arg interface{}) (string, error) {
	s, ok := arg.(string)
	if ok {
		return s, nil
	}

	f, ok := arg.(float64)
	if ok {
		return strconv.FormatFloat(f, 'f', 0, 64), nil
	}

	s, err := parseBlock(arg)
	if err != nil {
		return "", utils.LogAndReturnError("Argument is not a string or block")
	}

	return s, nil
}
