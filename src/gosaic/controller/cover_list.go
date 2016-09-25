package controller

import "gosaic/environment"

func CoverList(env environment.Environment) {
	coverService := env.MustCoverService()

	covers, err := coverService.FindAll("covers.name ASC")
	if err != nil {
		env.Printf("Error finding covers: %s\n", err.Error())
		return
	}
	if len(covers) == 0 {
		// we are done
		return
	}

	for _, cover := range covers {
		env.Println(cover)
	}
}
