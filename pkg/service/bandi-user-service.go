package service

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"bandi.com/TestGo/pkg/data"
	"golang.org/x/net/context"

	// Dynamic Proto Registration

	"github.com/gogo/gateway"
	"github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	golang_jsonpb "github.com/golang/protobuf/jsonpb"

	//golang_proto "github.com/golang/protobuf/proto"

	// Import GORM-related packages.
	"github.com/jinzhu/gorm"

	// _ "github.com/jinzhu/gorm/dialects/postgres"

	// Necessary in order to check for transaction retry error codes.
	"github.com/gogo/protobuf/types"
	"github.com/lib/pq"

	// Dynamic creation of File Descriptors for Proto
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
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
	/*
		Proto we're trying to insert

		message Personal {
			Name name = 1;
			string phone_number = 2;
			int32 age = 3;
		}

		message Name {
			string first_name = 1;
			string last_name = 2;
			string middle_name = 3;
		}

	*/

	/*
		if se, ok := err.(interface{ GRPCStatus() *status.Status }); ok {
			return FromGRPCStatus(se.GRPCStatus()), true
		}
	*/

	p := protoparse.Parser{
		Accessor: protoparse.FileContentsFromMap(map[string]string{
			"foo/bar.proto": `
				syntax = "proto3";
				package foo;
				message Bar {
					string name = 1;
					int32 id = 2;
				}
				`,
			// imports above file as just "bar.proto", so we need an
			// import resolver to properly load and link
			"fu/baz.proto": `
				syntax = "proto3";
				package fu;
				import "foo/bar.proto";
				message Baz {
					repeated foo.Bar foobar = 1;
				}
				`,
		}),
		ImportPaths: []string{"foo"},
	}
	fds, err := p.ParseFilesButDoNotLink("foo/bar.proto", "fu/baz.proto")
	if err != nil {
		return nil, err
	}
	// sanity check: make sure linking fails without an import resolver
	linkedFiles, err := desc.CreateFileDescriptors(fds)
	if err != nil {
		return nil, err
	}
	/*
		// now try again with resolver
		var r desc.ImportResolver
		r.RegisterImportPath("foo/bar.proto", "bar.proto")
		linkedFiles, err := r.CreateFileDescriptors(fds)
		if err != nil {
			return nil, err
		}
	*/

	// quick check of the resulting files
	fd := linkedFiles["foo/bar.proto"]
	md := fd.FindMessage("foo.Bar")
	dm := dynamic.NewMessage(md)

	//proto.RegisterType((*dynamic.Message)(nil), "foo.Bar")
	//golang_proto.RegisterType((*dynamic.Message)(nil), "foo.Bar")

	dm = dynamic.NewMessage(md)
	dm.SetFieldByNumber(1, "kishore")
	dm.SetFieldByNumber(2, int32(123))

	a1, err := types.MarshalAny(dm)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Create Any Marshalled Data is : %v\n", a1)

	/*
		marshaler := util.NewErrorJSON(&gateway.JSONPb{
			OrigName:     true,
			EnumsAsInts:  false,
			EmitDefaults: true,
		})

			buf, err := marshaler.Marshal(dm)
			if err != nil {
				return nil, err
			}
			fmt.Printf("Buffer Created is : %v\n", buf)

				resp := &data.CreateAnyResponse{Object: &types.Any{
					TypeUrl: "bandi.com/dynamic/" + proto.MessageName(dm),
					Value:   buf,
				}}
	*/

	resp := &data.CreateAnyResponse{Object: a1}
	fmt.Printf("Final Resp is : %v\n", resp)

	resolver := dynamic.AnyResolver(nil, fd)
	golangMarshaller := golang_jsonpb.Marshaler{AnyResolver: resolver}
	js1, err := golangMarshaller.MarshalToString(resp)
	if err != nil {
		fmt.Printf("Error is is : %v\n", err)
		return nil, err
	}
	fmt.Printf("After Marshal Resp1 is : %v\n", js1)

	jsm := &gateway.JSONPb{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: true,
		AnyResolver:  AnyResolver(nil, fd),
	}
	js, err := jsm.Marshal(resp)
	if err != nil {
		fmt.Printf("Error is is : %v\n", err)
		return nil, err
	}
	fmt.Printf("After Marshal Resp is : %s\n", string(js))

	/*
		resp1 := a1
		dynamicResolver := dynamic.AnyResolver(nil, fd)
		anotherMarshaller := golang_jsonpb.Marshaler{AnyResolver: dynamicResolver}
		js, err := anotherMarshaller.MarshalToString(resp1)
		if err != nil {
			return nil, fmt.Errorf("could not serialize to string another marshaller")
		}
		fmt.Printf("string de-serializer is : %v", js)

		_, err = marshaler.Marshal(resp1)
		if err != nil {
			return nil, fmt.Errorf("could not serialize anything")
		}
		fmt.Printf("FInal Resp is : %v", resp1)
	*/

	return resp, nil
}

// AnyResolver returns a jsonpb.AnyResolver that uses the given file descriptors
// to resolve message names. It uses the given factory, which may be nil, to
// instantiate messages. The messages that it returns when resolving a type name
// may often be dynamic messages.
func AnyResolver(mf *dynamic.MessageFactory, files ...*desc.FileDescriptor) jsonpb.AnyResolver {
	return &anyResolver{mf: mf, files: files}
}

type anyResolver struct {
	mf      *dynamic.MessageFactory
	files   []*desc.FileDescriptor
	ignored map[*desc.FileDescriptor]struct{}
	other   jsonpb.AnyResolver
}

// Resolve ...
func (r *anyResolver) Resolve(typeURL string) (proto.Message, error) {
	mname := typeURL
	if slash := strings.LastIndex(mname, "/"); slash >= 0 {
		mname = mname[slash+1:]
	}

	// see if the user-specified resolver is able to do the job
	if r.other != nil {
		msg, err := r.other.Resolve(typeURL)
		if err == nil {
			return msg, nil
		}
	}
	// try to find the message in our known set of files
	checked := map[*desc.FileDescriptor]struct{}{}
	for _, f := range r.files {
		md := r.findMessage(f, mname, checked)
		if md != nil {
			return r.mf.NewMessage(md), nil
		}
	}
	// failing that, see if the message factory knows about this type
	var ktr *dynamic.KnownTypeRegistry
	if r.mf != nil {
		v := reflect.ValueOf(r.mf)
		y := v.FieldByName("ktr")
		ktr = y.Interface().(*dynamic.KnownTypeRegistry)
	} else {
		ktr = (*dynamic.KnownTypeRegistry)(nil)
	}
	m := ktr.CreateIfKnown(mname)
	if m != nil {
		return m, nil
	}

	// no other resolver to fallback to? mimic default behavior
	mt := proto.MessageType(mname)
	if mt == nil {
		return nil, fmt.Errorf("unknown message type %q", mname)
	}
	return reflect.New(mt.Elem()).Interface().(proto.Message), nil
}

func (r *anyResolver) findMessage(fd *desc.FileDescriptor, msgName string, checked map[*desc.FileDescriptor]struct{}) *desc.MessageDescriptor {
	// if this is an ignored descriptor, skip
	if _, ok := r.ignored[fd]; ok {
		return nil
	}

	// bail if we've already checked this file
	if _, ok := checked[fd]; ok {
		return nil
	}
	checked[fd] = struct{}{}

	// see if this file has the message
	md := fd.FindMessage(msgName)
	if md != nil {
		return md
	}

	// if not, recursively search the file's imports
	for _, dep := range fd.GetDependencies() {
		md = r.findMessage(dep, msgName, checked)
		if md != nil {
			return md
		}
	}
	return nil
}
