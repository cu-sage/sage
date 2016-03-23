# Tests
Goal
whenGreenFlag doWaitUntil touching: Ball say:duration:elapsed:from: You win! 2

Ball
whenGreenFlag gotoX:y: -205 147 doForever doIf touchingColor: -15399425 forward: -10
whenKeyPressed up arrow heading: 0 forward: 10
whenKeyPressed down arrow heading: 180 forward: 10
whenKeyPressed right arrow heading: 90 forward: 10
whenKeyPressed left arrow heading: -90 forward: 10
