package service

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"bandi.com/TestGo/pkg/data"
	"golang.org/x/net/context"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"

	// Necessary in order to check for transaction retry error codes.
	"github.com/lib/pq"
)

// BandiUserService helps manage users
type BandiUserService struct {
	userName string
}

// NewBandiUserService creates a new user service
func NewBandiUserService() *BandiUserService {
	return &BandiUserService{}
}

// CreateBandiUser creates a new user
func (us *BandiUserService) CreateBandiUser(ctx context.Context, req *data.CreateBandiUserRequest) (*data.CreateBandiUserResponse, error) {
	fmt.Printf("Request for CreateBandiUser is : %v", req)
	// Connect to the "bank" database as the "maxroach" user.
	const addr = "postgres://acmeuser:acmepassword@localhost:26257/acmecorp?sslmode=verify-full&sslrootcert=/Users/bandi/go/src/skyflow.com/privacydb/assets/certs/localhost/ca.crt"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	// Set to `true` and GORM will print out all DB queries.
	db.LogMode(false)

	// Automatically create the "Bandi User" table based on the proto
	// model.
	db.AutoMigrate(&data.BandiUser{}, &data.BandiCreditCard{})
	db.Table("user_details").CreateTable(&data.BandiUserDetails{})

	// Insert one row
	db.Create(req.User)
	userdetails := &data.BandiUserDetails{
		Firstname:          "Kishore",
		Lastname:           "Bandi Details",
		ForeignKeyUserName: req.User.Username,
	}
	db.Table("user_details").Create(userdetails)
	/*
		if err := runTransaction(db,
			func(*gorm.DB) error {
				return updateUser(db, req.User)
			},
		); err != nil {
			fmt.Printf("Error running transaction: %v", err)
			return nil, err
		}
	*/

	return &data.CreateBandiUserResponse{}, nil
}

func updateUser(db *gorm.DB, user *data.BandiUser) error {
	var loadUser data.BandiUser

	db.Preload("CreditCards").Where(&data.BandiUser{Username: user.Username}).First(&loadUser)

	loadUser.CreditCards = []*data.BandiCreditCard{
		&data.BandiCreditCard{
			Number: "1234566789",
			Cvv:    123,
			ID:     "abc",
		},
		&data.BandiCreditCard{
			Number: "9876543321",
			Cvv:    254,
			ID:     "def",
		},
	}

	if err := db.Save(&loadUser).Error; err != nil {
		return err
	}
	return nil
}

// Functions of type `txnFunc` are passed as arguments to our
// `runTransaction` wrapper that handles transaction retries for us
// (see implementation below).
type txnFunc func(*gorm.DB) error

// Wrapper for a transaction.  This automatically re-calls `fn` with
// the open transaction as an argument as long as the database server
// asks for the transaction to be retried.
func runTransaction(db *gorm.DB, fn txnFunc) error {
	var maxRetries = 3
	for retries := 0; retries <= maxRetries; retries++ {
		if retries == maxRetries {
			return fmt.Errorf("hit max of %d retries, aborting", retries)
		}
		txn := db.Begin()
		if err := fn(txn); err != nil {
			// We need to cast GORM's db.Error to *pq.Error so we can
			// detect the Postgres transaction retry error code and
			// handle retries appropriately.
			pqErr := err.(*pq.Error)
			if pqErr.Code == "40001" {
				// Since this is a transaction retry error, we
				// ROLLBACK the transaction and sleep a little before
				// trying again.  Each time through the loop we sleep
				// for a little longer than the last time
				// (A.K.A. exponential backoff).
				txn.Rollback()
				var sleepMs = math.Pow(2, float64(retries)) * 100 * (rand.Float64() + 0.5)
				fmt.Printf("Hit 40001 transaction retry error, sleeping %s milliseconds\n", sleepMs)
				time.Sleep(time.Millisecond * time.Duration(sleepMs))
			} else {
				// If it's not a retry error, it's some other sort of
				// DB interaction error that needs to be handled by
				// the caller.
				return err
			}
		} else {
			// All went well, so we try to commit and break out of the
			// retry loop if possible.
			if err := txn.Commit().Error; err != nil {
				pqErr := err.(*pq.Error)
				if pqErr.Code == "40001" {
					// However, our attempt to COMMIT could also
					// result in a retry error, in which case we
					// continue back through the loop and try again.
					continue
				} else {
					// If it's not a retry error, it's some other sort
					// of DB interaction error that needs to be
					// handled by the caller.
					return err
				}
			}
			break
		}
	}
	return nil
}

// GetBandiUser Gets existing user
func (us *BandiUserService) GetBandiUser(ctx context.Context, req *data.GetBandiUserRequest) (*data.GetBandiUserResponse, error) {
	fmt.Printf("Request for CreateBandiUser is : %v", req)
	// Connect to the "bank" database as the "maxroach" user.
	const addr = "postgres://acmeuser:acmepassword@localhost:26257/acmecorp?sslmode=verify-full&sslrootcert=/Users/bandi/go/src/skyflow.com/privacydb/assets/certs/localhost/ca.crt"
	db, err := gorm.Open("postgres", addr)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	// Set to `true` and GORM will print out all DB queries.
	db.LogMode(false)

	// Automatically create the "Bandi User" table based on the proto
	// model.
	db.AutoMigrate(&data.BandiUser{}, &data.BandiCreditCard{})
	db.Table("user_details").CreateTable(&data.BandiUserDetails{})

	var bandiUser data.BandiUser
	db.Preload("CreditCards").Where(&data.BandiUser{Username: req.Name}).First(&bandiUser)
	return &data.GetBandiUserResponse{User: &bandiUser}, nil
}

// CreateAny ...
func (us *BandiUserService) CreateAny(ctx context.Context, req *data.CreateAnyRequest) (*data.CreateAnyResponse, error) {
	return &data.CreateAnyResponse{}, nil
}
