package main

func main() {
	boardingProcessor := BuildBoardingProcessorChain()
	passenger := &Passenger{
		name:                  "李四",
		hasBoardingPass:       false,
		hasLuggage:            true,
		isPassIdentityCheck:   false,
		isPassSecurityCheck:   false,
		isCompleteForBoarding: false,
	}
	boardingProcessor.ProcessFor(passenger)
}

func BuildBoardingProcessorChain() BoardingProcessor {
	completeBoardingNode := &completeBoardingProcessor{}

	securityCheckNode := &securityCheckProcessor{}
	securityCheckNode.SetNextProcessor(completeBoardingNode)

	identityCheckNode := &identityCheckProcessor{}
	identityCheckNode.SetNextProcessor(securityCheckNode)

	luggageCheckInNode := &luggageCheckInProcessor{}
	luggageCheckInNode.SetNextProcessor(identityCheckNode)

	boardingPassNode := &boardingPassProcessor{}
	boardingPassNode.SetNextProcessor(luggageCheckInNode)

	return boardingPassNode
}
