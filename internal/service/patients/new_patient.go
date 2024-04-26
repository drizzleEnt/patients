package patients

import (
	"context"
	"fmt"
)

func (s *srv) NewPatient(ctx context.Context) {
	fmt.Println("new")
}
