package grpc_go_error_handling

import (
	"fmt"
	epb "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"os"
	"testing"
	"ultimate-go/design-pattern/go-design/grpc_go_error_handling/status"
)

func Test_error_handling(t *testing.T) {
	// Code to Status
	st := status.New(codes.ResourceExhausted, "Request limit exceeded.")
	fmt.Printf("status.New: %v| %v| %v| %v \n", st.Err(), st.Message(), st.Code(), st.Details())
	// Status add Details
	ds, _ := st.WithDetails(
		&epb.QuotaFailure{
			Violations: []*epb.QuotaFailure_Violation{{
				Description: "Limit one greeting per person",
			}},
		},
	)
	fmt.Printf("status.WithDetails: %v | %v | %v| %v \n", st.Err(), st.Message(), st.Code(), st.Details())
	// Status to Error
	err := ds.Err()

	if err != nil {
		// Error to Status
		s := status.Convert(err)
		fmt.Printf("status.Convert: %v | %v | %v | %v \n", st.Err(), st.Message(), st.Code(), st.Details())
		// Status Details
		for _, d := range s.Details() {
			switch info := d.(type) {
			case *epb.QuotaFailure:
				fmt.Printf("Quota failure: %s", info)
			default:
				fmt.Printf("Unexpected type: %s", info)
			}
		}
		os.Exit(1)
	}
}
